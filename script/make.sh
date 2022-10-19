#!/bin/bash

# 编译xii助手
cd ../assistant
gf build -n xii_mac -a amd64,arm64  -s darwin
gf build -n xii_linux -a amd64,arm,arm64 -s linux

# 迁移xii助手
cd ../
rm -rf ./release

mkdir ./release
mkdir ./release/mac
mkdir ./release/mac_arm
mkdir ./release/linux
mkdir ./release/linux_arm
mkdir ./release/linux_arm64

mv ./assistant/temp/darwin_amd64/xii_mac ./release/mac/xii
mv ./assistant/temp/darwin_arm64/xii_mac ./release/mac_arm/xii
mv ./assistant/temp/linux_amd64/xii_linux ./release/linux/xii
mv ./assistant/temp/linux_arm64/xii_linux ./release/linux_arm64/xii
mv ./assistant/temp/linux_arm/xii_linux ./release/linux_arm/xii

# 创建目录
mkdir ./release/mac/data
mkdir ./release/mac_arm/data
mkdir ./release/linux_arm64/data
mkdir ./release/linux_arm/data
mkdir ./release/linux/data


mkdir ./release/mac/www
mkdir ./release/mac_arm/www
mkdir ./release/linux_arm64/www
mkdir ./release/linux_arm/www
mkdir ./release/linux/www


mkdir ./release/mac/logs
mkdir ./release/mac_arm/logs
mkdir ./release/linux_arm64/logs
mkdir ./release/linux_arm/logs
mkdir ./release/linux/logs


mkdir ./release/mac/env
mkdir ./release/mac_arm/env
mkdir ./release/linux_arm64/env
mkdir ./release/linux_arm/env
mkdir ./release/linux/env


# 复制文件
cp -r ./repo ./release/mac/repo
cp -r ./repo ./release/mac_arm/repo
cp -r ./repo ./release/linux_arm64/repo
cp -r ./repo ./release/linux_arm/repo
cp -r ./repo ./release/linux/repo


cp -r ./script/manual.sh ./release/mac/manual.sh
cp -r ./script/manual.sh ./release/mac_arm/manual.sh
cp -r ./script/manual.sh ./release/linux_arm64/manual.sh
cp -r ./script/manual.sh ./release/linux_arm/manual.sh
cp -r ./script/manual.sh ./release/linux/manual.sh

# 打包
cd ./release
zip -r xii_mac.zip ./mac
zip -r xii_mac_arm.zip ./mac_arm
zip -r xii_linux_arm64.zip ./linux_arm64
zip -r xii_linux_arm.zip ./linux_arm
zip -r xii_linux.zip ./linux


# 清理
rm -rf ./mac
rm -rf ./mac_arm
rm -rf ./linux_arm64
rm -rf ./linux_arm
rm -rf ./linux

