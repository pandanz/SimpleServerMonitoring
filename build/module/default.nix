{ stdenv, fetchurl, systemd, lib, unzip, autoPatchelfHook }:

stdenv.mkDerivation rec {
  pname = "server_monitoring";
  version = "v0.1";

 src = ../../build;

  nativeBuildInputs = [
    unzip
    autoPatchelfHook
 ];

  buildInputs = [ systemd ];

  unpackPhase = ''
    ls -lah $src
  '';

  installPhase = ''
    runHook preInstall
    mkdir -p $out/bin
    cp $src/server_monitoring $out/bin
    runHook postInstall
  '';

  meta = with lib; {
    description = "Server Monitoring Service";
    platforms = platforms.linux;
  };
}
