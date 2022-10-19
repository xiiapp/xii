#!/bin/bash

# 编译xii助手
cd ../assistant
gf build -n xii_mac -a amd64  -s darwin
gf build -n xii_linux -a amd64  -s linux

# 迁移xii助手
cd ../
rm -rf ./release_dev/linux
rm -rf ./release_dev/mac
rm -f ./release_dev/xii_linux.zip
rm -f ./release_dev/xii_mac.zip

mkdir ./release_dev/linux
mkdir ./release_dev/mac

mv ./assistant/temp/darwin_amd64/xii_mac ./release_dev/mac/xii
mv ./assistant/temp/linux_amd64/xii_linux ./release_dev/linux/xii

# 创建目录
mkdir ./release_dev/linux/data
mkdir ./release_dev/mac/data

mkdir ./release_dev/linux/www
mkdir ./release_dev/mac/www

mkdir ./release_dev/linux/logs
mkdir ./release_dev/mac/logs

mkdir ./release_dev/linux/env
mkdir ./release_dev/mac/env

# 复制文件
cp -r ./repo ./release_dev/linux/repo
cp -r ./repo ./release_dev/mac/repo
cp -r ./script/manual.sh ./release_dev/mac/manual.sh
cp -r ./script/manual.sh ./release_dev/linux/manual.sh

# 打包
cd ./release_dev
zip -r xii_linux.zip ./linux
zip -r xii_mac.zip ./mac
rm -rf ./linux
rm -rf ./mac
