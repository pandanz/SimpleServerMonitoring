## Simple Resource Monitoring
This repository contains the Nix expression for building the Resource Monitoring Service.

## default.nix

The `default.nix` file is a Nix expression that describes how to build the Resource Monitoring Service.

Here's a breakdown of what each part of the file does:

- The file starts by importing several packages that are needed to build the service. These include `stdenv` (the standard Nix environment), `fetchurl` (a function for downloading files), `systemd` (a system and service manager), `lib` (a collection of useful functions), `unzip` (a utility for unpacking zip files), and `autoPatchelfHook` (a tool for automatically patching the ELF files in a package).

- `stdenv.mkDerivation rec` is used to create a new derivation (a build action). The `rec` keyword allows the attributes inside the derivation to refer to each other.

- `pname` and `version` define the name and version of the package.

- `src` uses `fetchurl` to download the source code of the service from a GitHub release. The `sha256` attribute is used to verify the integrity of the downloaded file.

- `nativeBuildInputs` and `buildInputs` list the packages that are needed to build the service. `unzip` and `autoPatchelfHook` are needed during the build process, while `systemd` is a runtime dependency.

- `unpackPhase` and `installPhase` define custom phases of the build process. `unpackPhase` unzips the source file, and `installPhase` copies the binary to the output directory.

- `meta` contains metadata about the package, such as its description and the platforms it supports.
