#!/bin/bash

# Install Homebrew if not already installed
if ! [ -x "$(command -v brew)" ]; then
  echo 'Installing Homebrew...'
  /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
fi

# Install Docker
if ! [ -x "$(command -v docker)" ]; then
  echo 'Installing Docker...'
  brew install --cask docker
  open /Applications/Docker.app
  echo "Please wait for Docker to start..."
  sleep 30
fi

# Install Docker Compose
if ! [ -x "$(command -v docker-compose)" ]; then
  echo 'Installing Docker Compose...'
  brew install docker-compose
fi

# Clone the blog repository
git clone https://github.com/yourusername/blog-project.git
cd blog-project

# Build and start Docker containers
docker-compose up -d

echo "Installation complete! Your blog is now running at http://localhost:8080"