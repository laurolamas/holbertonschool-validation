#!/bin/bash

# Update apt-get package lists
apt-get update

# Install Hugo and Make
apt-get install -y hugo make

# Build the website
make build
