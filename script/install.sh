#!/bin/bash

red() {
  echo -e "\033[31m\033[01m$1\033[0m"
}

green() {
  echo -e "\033[32m\033[01m$1\033[0m"
}

Linfo(){
  green "Checking System..."
  CPU=$(uname -a)
if [[ "$CPU" =~ "aarch64" ]]; then
  S=_arm64
  url=https://github.com/xiiapp/xii/raw/main/release/xii_linux_arm64.zip
elif [[ "$CPU" =~ "arm" ]]; then
  S=_arm
  url=https://github.com/xiiapp/xii/raw/main/release/xii_linux_arm.zip
elif [[ "$CPU" =~ "x86_64" ]]; then
  S=
  url=https://github.com/xiiapp/xii/raw/main/release/xii_linux.zip
else
  red "脚本不支持此服务器架构,请尝试手动安装"
  exit 1
fi
}
xiiapp
Minfo(){
  CPU=$(uname -a)
if [[ "$CPU" =~ "aarch64" ]]; then
  S=_arm
  url=https://github.com/xiiapp/xii/raw/main/release/xii_mac_arm.zip
elif  [[ "$CPU" =~ "x86_64" ]]; then
  S=
  url=https://github.com/xiiapp/xii/raw/main/release/xii_mac.zip
fi
}

cncheck(){
  green "Checking Internet"
  ipcheck=$(curl -s "ipinfo.io")
  cn=$(echo $ipcheck | grep "CN" )
if [ -n "$cn" ];then
  echo "当前为大陆环境，使用镜像源安装"
  durl="get.daocloud.io/docker"
  curl="dn-dao-github-mirror.daocloud.io"
else
  durl="get.docker.com"
  curl="github.com"
  Location="world"
fi
}

docker(){
  green "Checking docker"
  docker=$(command -v docker)
  compose=$(command -v docker-compose)
  # [[ -z $(docker -v 2>/dev/null) ]] && docker="False"
  # [[ -n $(docker -v 2>/dev/null) ]] && green "docker:Installed"
  # [[ -z $(docker-compose -v 2>/dev/null) ]] && docker-compose="False"
  # [[ -n $(docker-compose -v 2>/dev/null) ]] && green "docker-compose:Installed"
if [[ -z "$docker" ]]; then
  green "Start installing docker"
  curl -sSL $durl | sh
  systemctl enable docker
  systemctl start docker
  groupadd docker
  if [ "root" !=  "$USER" ]; then
      usermod -aG docker $USER
  fi
fi
if [[ -z "$compose" ]]; then
    curl -L "${curl}/docker/compose/releases/download/v2.12.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    chmod +x /usr/local/bin/docker-compose
fi
}

depend(){
  green "Checking depend"
depends=("curl" "wget" "unzip" "sudo")
depend=""
for i in "${!depends[@]}"; do
  now_depend="${depends[$i]}"
  if [ ! -x "$(command -v $now_depend 2>/dev/null)" ]; then
    echo "$now_depend 未安装"
    depend="$now_depend $depend"
  fi
done
if [ "$depend" ]; then
  if [ -x "$(command -v apk 2>/dev/null)" ]; then
    echo "apk包管理器,正在尝试安装依赖:$depend"
    apk --no-cache add $depend $proxy >>/dev/null 2>&1
  elif [ -x "$(command -v apt-get 2>/dev/null)" ]; then
    echo "apt-get包管理器,正在尝试安装依赖:$depend"
    apt -y install $depend >>/dev/null 2>&1
  elif [ -x "$(command -v yum 2>/dev/null)" ]; then
    echo "yum包管理器,正在尝试安装依赖:$depend"
    yum -y install $depend >>/dev/null 2>&1
  else
    red "未找到合适的包管理工具,请手动安装:$depend"
    exit 1
  fi
  for i in "${!depends[@]}"; do
    now_depend="${depends[$i]}"
    if [ ! -x "$(command -v $now_depend)" ]; then
      red "$now_depend 未成功安装,请尝试手动安装!"
      exit 1
    fi
  done
fi
}

Minstall(){
  green "MacOS:Start installing xii"
  cd ~

  if [  -d "xii" ] ; then
    mv xii xii.bakup.$RANDOM
  fi

  mkdir xii
  cd xii
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
}

Linstall(){
  green "Linux:Start installing xii"
  cd /home

  if [  -d "xii" ] ; then
    mv xii xii.bakup.$RANDOM
  fi
  mkdir xii
  cd xii
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
}

sys=$(uname -a)
if [[ "$sys" =~ "Linux" ]]; then
  green "System:Linux"
Linfo
depend
cncheck
docker
Linstall
fi
if [[ "$sys" =~ "Darwin" ]]; then
  green "System:MacOS"
Minfo
depend
cncheck
docker
Minstall
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
