#!/bin/bash

# Install prerequisites
apt-get update
apt-get install -y wget make

# Download and install Hugo
wget https://github.com/gohugoio/hugo/releases/download/v0.55.0/hugo_0.55.0_Linux-64bit.deb
dpkg -i hugo_0.55.0_Linux-64bit.deb

#INSTALL golangcilint
#export PATH="$PATH:/app/bin"
#sudo curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s
