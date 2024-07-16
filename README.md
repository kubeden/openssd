# openssd

Openssd stands for Open (source) Super Simple Diary. Openssd is a plug-n-run project to host a version of yourself on the webternet.

## Prerequisites

1. Debian or Ubuntu internet-facing virtual machine
2. This repository forked or your own public github repository

## Installation

To set up your environment, log into your server and run the install command.

```
ssh root@[ip_address]
curl -sSL https://raw.githubusercontent.com/kubeden/openssd/main/scripts/install.sh | bash
```

## Configuration & Deployment

To configure and deploy your openssd, go to the openssd directory:

```
cd /var/www/openssd
```

open *docker-compose.yml* with a text editor of your choice and change the following strings:

```
- GITHUB_USERNAME=kubeden
- GITHUB_REPO=kubeden
- README_FILE=README.md
- INFO_FILE=INFO.md
- X_USER_FULL_NAME=Kuberdenis
- X_USERNAME=kubeden
- TEMPLATE_CHOICE=ssi
```

Make sure the github username & repository you have added to *docker-compose.yml* exist and are browsable.

Then when you have everything set up, run the following command:

```
curl -sSL https://raw.githubusercontent.com/kubeden/openssd/main/scripts/start.sh | bash
```