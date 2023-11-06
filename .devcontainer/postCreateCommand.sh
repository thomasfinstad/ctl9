#!/bin/bash -ex

###
# Apt packages
export DEBIAN_FRONTEND=noninteractive
sudo apt-get update

# Get some completion up in here
sudo apt-get install -y bash-completion

###
# Pre-commit / git tools

# Used for commit message formatting
pipx install argcomplete
pipx install Commitizen
register-python-argcomplete cz >> ~/.bashrc

# Used to improve commits before they are commited
pipx install pre-commit
pre-commit install --hook-type pre-commit --hook-type commit-msg

# Tooling
go install mvdan.cc/gofumpt@latest
go install golang.org/x/tools/cmd/goimports@latest

# Libraries
go mod download

# https://ebitengine.org/en/documents/install.html?os=linux
sudo apt-get install -y \
    gcc \
    libc6-dev \
    libgl1-mesa-dev \
    libxcursor-dev \
    libxi-dev \
    libxinerama-dev \
    libxrandr-dev \
    libxxf86vm-dev \
    libasound2-dev \
    pkg-config


# https://github.com/go-vgo/robotgo#ubuntu
sudo apt-get install -y \
    gcc \
    libc6-dev \
    libx11-dev \
    xorg-dev \
    libxtst-dev \
    xsel \
    xclip \
    libpng++-dev \
    xcb \
    libxcb-xkb-dev \
    x11-xkb-utils \
    libx11-xcb-dev \
    libxkbcommon-x11-dev \
    libxkbcommon-dev
