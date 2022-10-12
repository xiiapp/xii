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
if [ "$(uname)" == "Darwin" ] ; then
  cd ~

  if [  -d "xii" ] ; then
    mv xii xii.bakup.$RANDOM
  fi
  git clone -n  https://github.com/xiiapp/xii.git
  mkdir xii
  cd xii
  git checkout main -- release/xii_mac.zip
  mv  -f ~/xii/release/xii_mac.zip ~/xii/xii.zip
  unzip xii.zip -d ./
  mv -f ~/xii/mac/* ./
  rm -rf ~/xii/release
  rm -f ~/xii/xii.zip
  mv -f ~/xii/mac ~/xii/
  cp -f ~/xii/env.config ~/xii/.env
  chmod +x ~/xii/xii
  ln -s ~/xii/xii /usr/local/bin/xii
  ln -s ~/xii/xii /usr/local/bin/xxi

else
  cd /home

  if [  -d "xii" ] ; then
    mv xii xii.bakup.$RANDOM
  fi
  git clone -n  https://github.com/xiiapp/xii.git
  mkdir xii
  cd xii
  git checkout main -- release/xii_linux.zip
  mv  -f /home/xii/release/xii_linux.zip /home/xii/xii.zip
  unzip xii.zip -d ./
  mv -f /home/xii/linux/* ./
  rm -rf /home/xii/release
  rm -f /home/xii/xii.zip
  mv -f /home/xii/linux /home/xii/
  cp -f /home/xii/env.config /home/xii/.env
  chmod +x /home/xii/xii
  ln -s /home/xii/xii /usr/local/bin/xii
  ln -s /home/xii/xii /usr/local/bin/xxi

fi







mv
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