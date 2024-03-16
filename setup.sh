#!/usr/bin/env sh
chmod a+x install-docker.sh
sh install-docker.sh

# node and go managers
sudo curl -fsSL https://raw.githubusercontent.com/tj/n/master/bin/n | bash -s lts
curl -sSL https://git.io/g-install | sh -s

# install Caddy

sudo apt install -y debian-keyring debian-archive-keyring apt-transport-https curl
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
sudo apt update
sudo apt install caddy