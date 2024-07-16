#!/bin/bash

apt update && apt upgrade -y
sudo apt-get install build-essential procps curl file git -y

wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz
rm go1.22.5.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" > ~/.bashrc
source ~/.bashrc

apt install docker.io
apt install docker-compose

mkdir /var/www
git clone https://github.com/kubeden/openssd.git /var/www