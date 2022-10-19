#!/bin/bash

# Location variables
Location="world"
if [ "$1" == "china" ] ; then
  Location=$1
elif [ "$2" == "china" ] ; then
  Location=$1
fi

# 架构
S=""
if [ "$1" == "arm" ] ; then
  S="_$1"
elif [ "$2" == "arm" ] ; then
  S="_$1"
fi

if [ "$1" == "arm64" ] ; then
  S="_$1"
elif [ "$2" == "arm64" ] ; then
  S="_$1"
fi









# Install wget & curl
if command -v wget >/dev/null 2>&1; then
  echo 'wget already installed.Exit'
else
  yum install -y wget
  apt-get install -y wget
  apk add wget
fi

if command -v curl >/dev/null 2>&1; then
  echo 'curl already installed.Exit'
else
  yum install -y curl
  apt-get install -y curl
  apk add curl
fi

if command -v unzip >/dev/null 2>&1; then
  echo 'unzip already installed.Exit'
else
  yum install -y unzip
  apt-get install -y unzip
  apk add unzip
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
  if [ "$Location" == "china" ] ; then
    curl -L "https://get.daocloud.io/docker/compose/releases/download/v2.11.2/docker-compose-$(uname -s)-$(uname -m)" > /usr/local/bin/docker-compose
  else
    curl -L "https://github.com/docker/compose/releases/download/v2.11.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
  fi
 chmod +x /usr/local/bin/docker-compose
 ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
fi



# Create Floder
if [ "$(uname)" == "Darwin" ] ; then
  cd ~

  if [  -d "xii" ] ; then
    mv xii xii.bakup.$RANDOM
  fi

  mkdir xii
  cd xii

  url=" https://github.com/xiiapp/xii/raw/main/release/xii_mac.zip"
  if [  -n "$S" ]; then
    url="https://github.com/xiiapp/xii/raw/main/release/xii_mac_arm.zip"
  fi

  wget -c "$url"  -O xii.zip
  unzip xii.zip -d ./
  mv -f ~/xii/mac"$S"/* ./
  rm -rf ~/xii/release
  rm -f ~/xii/xii.zip
  if [ Location == "world" ] ; then
      sed -i ""  's/CONTAINER_PACKAGE_URL=mirrors.ustc.edu.cn/CONTAINER_PACKAGE_URL=/g' ~/xii/repo/base/env.sample
  fi
  chmod +x ~/xii/xii
  echo "创建软链接需要输入密码授权"
  sudo ln -s ~/xii/xii /usr/local/bin/xii
  sudo ln -s ~/xii/xii /usr/local/bin/xxi




else
  cd /home

  if [  -d "xii" ] ; then
    mv xii xii.bakup.$RANDOM
  fi
  mkdir xii
  cd xii

  url=" https://github.com/xiiapp/xii/raw/main/release/xii_linux.zip"
  if [  -n "$S" ]; then
    url="https://github.com/xiiapp/xii/raw/main/release/xii_linux$S.zip"
  fi

#  wget -c https://github.com/xiiapp/xii/raw/main/release/xii_linux.zip  -O xii.zip
  wget -c "$url"  -O xii.zip
  unzip xii.zip -d ./
  mv -f /home/xii/linux"$S"/* ./
  rm -rf /home/xii/release
  rm -f /home/xii/xii.zip
  if [ Location == "world" ] ; then
      sed -i "" 's/CONTAINER_PACKAGE_URL=mirrors.ustc.edu.cn/CONTAINER_PACKAGE_URL=/g' /home/xii/repo/base/env.sample
  fi
  chmod +x /home/xii/xii
  ln -s /home/xii/xii /usr/local/bin/xii
  ln -s /home/xii/xii /usr/local/bin/xxi

fi








echo "  "
echo "  "
echo "  "
echo -e "\033[33m ---Success---\033[0m"
echo -e "\033[33mXII is installed successfully.\033[0m"
if [ "$(uname)" == "Darwin" ] ; then
  echo -e "\033[33m Your app location is ~/xii \033[0m"
else
    echo -e "[33m Your app location is /home/xii [0m"

fi
echo -e "\033[33mPlease visit https://xii.app for detail information before use.\033[0m"
echo "  "
echo "  "
echo "  "
echo "  "

# export PATH=$PATH:/home/xii/assistant or ln -s /home/xii/assistant/xii  /usr/bin/xii