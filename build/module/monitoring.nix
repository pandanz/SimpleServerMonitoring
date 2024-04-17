{ pkgs, ... }:

let
  # "let" allows you to define local variables
  server_monitoring = pkgs.callPackage ./default.nix { };
in
  # "in" allows you to used the variables defined in the "let" block
{
  environment.systemPackages = with pkgs; [
    server_monitoring
  ];

  systemd.user.services.server_monitoring = {
    description = "Simple Server Monitoring API";
    enable = false;
    serviceConfig = {
      ExecStart = "${server_monitoring}/bin/server_monitoring";
      Restart = "always";
      Environment = [
        "SERVER_MONITORING_PORT=8081"
        "SERVER_MONITORING_DEBUG=True"
      ];
    };
    restartTriggers = [ server_monitoring ];
    wantedBy = [ "multi-user.target" ];
  };
}