#!/bin/bash

apt update -y && apt upgrade -y

# install nginx
apt install -y nginx

# install docker
apt install -y docker.io

# install docker-compose
apt install -y docker-compose

# install git
apt install -y git

# clone the blog repository
git clone [repo-url]

# build and start docker containers
cd openssd

docker-compose up -d

echo "Installation complete! Your blog is now running!"