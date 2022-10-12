#!/bin/bash

# Install Git
if command -v git >/dev/null 2>&1; then
  echo 'Git already installed.Exit'
else
  yum install -y git
  apt-get install -y git
  apk add git
fi


# Install docker
if command -v docker >/dev/null 2>&1; then
  echo 'Docker already installed.Exit'
else
  curl -fsSL get.docker.com -o get-docker.sh
  sh get-docker.sh --mirror Aliyun
  systemctl enable docker
  systemctl start docker
  groupadd docker
  if [ "root" !=  "$USER" ]; then
      usermod -aG docker $USER
  fi
fi


# Install the docker-compose
if command -v docker-compose >/dev/null 2>&1; then
  echo 'docker-compose already installed.'
else
 curl -L https://get.daocloud.io/docker/compose/releases/download/v2.11.2/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
 #curl -L "https://github.com/docker/compose/releases/download/v2.11.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
 chmod +x /usr/local/bin/docker-compose
 ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
fi



# Create Floder
cd /home
if [  -d "./xii" ] ; then
  mv xii xii.bakup.$RANDOM
fi
git clone https://github.com/xiiapp/xii.git
cd xii
echo "  "
echo "  "
echo "  "
echo -e "\033[41;33m ---Success---\033[0m"
echo -e "\033[41;33mXII is installed successfully.\033[0m"
echo -e "\033[41;33mYour App path is /home/xii .\033[0m"
echo -e "\033[41;33mPlease run 'xii init' to initialize the environment.\033[0m"
echo -e "\033[41;33mPlease visit https://xii.app for detail information before use.\033[0m"
echo "  "
echo "  "
echo "  "
echo "  "

# export PATH=$PATH:/home/xii/assistant or ln -s /home/xii/assistant/xii  /usr/bin/xii