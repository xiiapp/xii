#!/bin/bash




exit 0


# Location variables
Location="world"
if [ "$1" == "china" ] ; then
  Location=$1
elif [ "$2" == "china" ] ; then
  Location=$1
fi


echo "$Location"



# Install Git
if command -v wget >/dev/null 2>&1; then
  echo 'wget already installed.Exit'
else
  yum install -y wget
  apt-get install -y wget
  apk add wget
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
    curl -L https://get.daocloud.io/docker/compose/releases/download/v2.11.2/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
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
  wget -c https://github.com/xiiapp/xii/raw/main/release/xii_mac.zip  -O xii.zip
  unzip xii.zip -d ./
  mv -f ~/xii/mac/* ./
  rm -rf ~/xii/release
  rm -f ~/xii/xii.zip
  cp -f ~/xii/env.config ~/xii/.env
  if [ Location == "world" ] ; then
      sed -i ""  's/CONTAINER_PACKAGE_URL=mirrors.ustc.edu.cn/CONTAINER_PACKAGE_URL=/g' ~/xii/.env
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
  wget -c https://github.com/xiiapp/xii/raw/main/release/xii_linux.zip  -O xii.zip
  unzip xii.zip -d ./
  mv -f /home/xii/linux/* ./
  rm -rf /home/xii/release
  rm -f /home/xii/xii.zip
  cp -f /home/xii/env.config /home/xii/.env
  if [ Location == "world" ] ; then
      sed -i ""  's/CONTAINER_PACKAGE_URL=mirrors.ustc.edu.cn/CONTAINER_PACKAGE_URL=/g' /home/xii/.env
  fi
  chmod +x /home/xii/xii
  ln -s /home/xii/xii /usr/local/bin/xii
  ln -s /home/xii/xii /usr/local/bin/xxi

fi








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