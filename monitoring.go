package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// var port string

type Sensor struct {
	Name       string                 `json:"name"`
	UniqueId   string                 `json:"unique_id"`
	State      float64                `json:"state"`
	Attributes map[string]interface{} `json:"attributes"`
}

func getHostname() string {
	hostname, _ := os.Hostname()
	return hostname
}

func createSensor(state float64, name string, unit string, attributes map[string]interface{}) Sensor {
	attributes["unit_of_measurement"] = unit
	return Sensor{
		Name:       name,
		UniqueId:   name,
		State:      state,
		Attributes: attributes,
	}
}

func getCPUUsage(hostname string) (Sensor, error) {
	attributes := make(map[string]interface{})
	percent, err := cpu.Percent(time.Second, true)
	if err != nil {
		return Sensor{}, err
	}

	for idx, usage := range percent {
		roundedUsage := fmt.Sprintf("%.2f", usage)
		attributes[fmt.Sprintf("%s_cpu_core%d", hostname, idx+1)] = roundedUsage
	}

	roundedTotalStr := fmt.Sprintf("%.2f", percent[0])
	roundedTotal, err := strconv.ParseFloat(roundedTotalStr, 64)
	if err != nil {
		return Sensor{}, err
	}

	sensor := createSensor(roundedTotal, fmt.Sprintf("%s CPU Usage", hostname), "%", attributes)

	return sensor, nil
}

func getMemoryUsage(hostname string) (Sensor, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return Sensor{}, err
	}

	// Convert memory usage to a percentage and round to 2 decimal places
	roundedMemory := fmt.Sprintf("%.2f", v.UsedPercent)
	memoryPercent, err := strconv.ParseFloat(roundedMemory, 64)
	if err != nil {
		return Sensor{}, err
	}

	attributes := make(map[string]interface{})
	sensor := createSensor(memoryPercent, fmt.Sprintf("%s Memory Usage", hostname), "%", attributes)

	return sensor, nil
}

func main() {
	var r *gin.Engine

	// Set Gin mode based on environment variable
    if debug, ok := os.LookupEnv("SERVER_MONITORING_DEBUG"); ok && debug == "True" {
        gin.SetMode(gin.DebugMode)
		r = gin.Default()
    } else {
        gin.SetMode(gin.ReleaseMode)
		r = gin.New()
    }

	// Set a default port
    port := "8080"
    // Try to read the port from the environment
    if envPort, ok := os.LookupEnv("SERVER_MONITORING_PORT"); ok {
        port = envPort
    }

	r.GET("/cpu", func(c *gin.Context) {
		hostname := getHostname()
		cpuSensor, err := getCPUUsage(hostname)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, cpuSensor)
	})
	r.GET("/memory", func(c *gin.Context) {
		hostname := getHostname()
		memorySensor, err := getMemoryUsage(hostname)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, memorySensor)
	})

	r.Run(":" + port) // listen and serve on 0.0.0.0:8080
}
