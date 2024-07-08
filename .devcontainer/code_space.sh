#!/bin/bash
DEBIAN_FRONTEND=noninteractive
sudo apt-get update
sudo apt-get install -y --no-install-recommends apt-utils dialog dnsutils httpie wget unzip curl jq
GOVERSION=1.18.10
wget https://go.dev/dl/go${GOVERSION}.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go${GOVERSION}.linux-amd64.tar.gz
CURRENT_DIR=`pwd`
cd /workspaces/movie-rating/playground/
./build.sh
npm install
echo "Completed Setup"
exit 0