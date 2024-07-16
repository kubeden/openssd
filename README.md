# openssd
Openssd stands for Open (source) Super Simple Diary. Openssd is a plug-n-run project to host a version of yourself on the webternet.

## Features

- Markdown-based articles stored in a Git repository
- Responsive design using Tailwind CSS
- Dynamic content loading with HTMX
- Docker-based deployment for easy setup

## Prerequisites

- Git
- Docker and Docker Compose (installed automatically by the installation script)

## Installation

### Linux

1. Open a terminal
2. Run the following command:

```bash
curl -sSL https://raw.githubusercontent.com/yourusername/blog-project/main/scripts/install_linux.sh | bash
```

### macOS

1. Open a terminal
2. Run the following command:

```bash
curl -sSL https://raw.githubusercontent.com/yourusername/blog-project/main/scripts/install_macos.sh | bash
```

## Usage

deploy server
```
apt update && apt upgrade -y
sudo apt-get install build-essential procps curl file git

wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz
rm go1.22.5.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" > ~/.bashrc
source ~/.bashrc

apt install nginx
apt install docker.io
apt install docker-compose

cd /var/www
git clone https://github.com/kubeden/openssd.git
```


## Customization

### Changing the theme

To change the blog's appearance:

1. Edit the HTML templates in the `ui/templates` directory
2. Modify the Tailwind CSS classes to adjust the styling

## License

This project is licensed under the MIT License - see the LICENSE file for details.