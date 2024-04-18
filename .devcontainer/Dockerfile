# Reuse a base image made for devcontainers.
FROM mcr.microsoft.com/vscode/devcontainers/base:bookworm

LABEL maintainer="PandaNZ"
ARG user=vscode

# These dependencies are required by Nix.
RUN apt update -y && \
    apt -y install --no-install-recommends curl xz-utils python3 python3-pip && \
    pip install --break-system-packages pre-commit

# Swap to a non-root user
USER ${user}

# Install nix in single user mode
RUN curl -L https://nixos.org/nix/install | sh && \
    . /home/${user}/.nix-profile/etc/profile.d/nix.sh

# Add the nix profile to the path
COPY profile.sh /etc/profile.d/devcontainer.sh

# Set the user for the devcontainer
ENV USER=${user}

# Set the volume to /nix
VOLUME /nix
