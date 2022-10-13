#!/bin/bash

echo "  手动安装：将xii_mac.zip或xii_linux.zip压缩包跟本脚本放在一起，然后执行本脚本将自动安装"

# Create Floder
if [ "$(uname)" == "Darwin" ] ; then
  mv xii_mac.zip ~/xii.zip
  cd ~

  if [  -d "xii" ] ; then
    mv xii xii.bakup.$RANDOM
  fi

  mkdir xii
  cd xii
  unzip xii.zip -d ./
  mv -f ~/xii/mac/* ./
  rm -rf ~/xii/release
  rm -f ~/xii/xii.zip
  cp -f ~/xii/env.config ~/xii/.env
  chmod +x ~/xii/xii
  echo "创建软链接需要输入密码授权"
  sudo ln -s ~/xii/xii /usr/local/bin/xii
  sudo ln -s ~/xii/xii /usr/local/bin/xxi

else
   mv xii_linux.zip /home/xii.zip
  cd /home
  if [  -d "xii" ] ; then
    mv xii xii.bakup.$RANDOM
  fi
  mkdir xii
  mv xii.zip xii/
  cd xii
  unzip xii.zip -d ./
  mv -f /home/xii/linux/* ./
  rm -rf /home/xii/release
  rm -f /home/xii/xii.zip
  cp -f /home/xii/env.config /home/xii/.env
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