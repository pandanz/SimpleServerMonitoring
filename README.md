# Simple Server Monitoring

[![GitHub Release][releases-shield]][releases]
[![GitHub Activity][commits-shield]][commits]
[![License][license-shield]](LICENSE)

![Project Maintenance][maintenance-shield]

[![Build](https://github.com/pandanz/SimpleServerMonitoring/actions/workflows/build.yml/badge.svg)](https://github.com/pandanz/SimpleServerMonitoring/actions/workflows/build.yml)

API endpoint to monitoring server resources without the need for node_exporter

**Configurable via environment variables.**

Variable | Description
-- | --
`RESOURCE_MONITORING_PORT` | Change the port the API is listening on eg. `8081`.
`RESOURCE_MONITORING_DEBUG` | API should run in debug mode eg. `True` or `False`

### CallOuts
Tom for devcontainer:
https://levelup.gitconnected.com/vs-code-remote-containers-with-nix-2a6f230d1e4e

### Useful
- autoPatchelfHook
- https://nixos.wiki/wiki/Packaging/Binaries
- https://blog.lenny.ninja/part-1-quickly-packaging-services-using-nix-flakes.html
- https://elatov.github.io/2022/01/building-a-nix-package/


[releases-shield]: https://img.shields.io/github/release/pandanz/SimpleServerMonitoring.svg?style=for-the-badge
[releases]: https://github.com/pandanz/SimpleServerMonitoring/releases
[commits-shield]: https://img.shields.io/github/commit-activity/y/pandanz/SimpleServerMonitoring.svg?style=for-the-badge
[commits]: https://github.com/pandanz/SimpleServerMonitoring/commits/main
[license-shield]: https://img.shields.io/github/license/pandanz/SimpleServerMonitoring.svg?style=for-the-badge
[maintenance-shield]: https://img.shields.io/badge/maintainer-pandanz-blue.svg?style=for-the-badge