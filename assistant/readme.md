# 助手程序
为整个环境提供一个辅助助手工具，用于协助用户快捷设置。



## 一期功能如下
```shell
用法
xii  命令 [参数]

命令
vhost            新增、删除、罗列所有网站
clear            清理镜像除配置外的所有数据

                 ---------以下为封装docker-compose相关命令---------
                 封装原因是：docker-compose通常需要进入目录或带-f参数，操作过于麻烦。
                 使用上注意所有容器名字需要带-，比如php容器就写做-php。

up               封装docker-compose up 命令,
                 首次使用请执行： xii up -h 查看下使用说明。
down             封装docker-compose down 命令,停止并删除容器，网络，图像和挂载卷
start            封装docker-compose start 命令,启动服务。范例：xii start -php
stop             封装docker-compose stop 命令,停止指定容器服务。范例：xii stop -php
restart          封装docker-compose stop 命令,重启指定容器服务。范例：xii restart -php
rm               封装docker-compose rm 命令,删除并且停止php容器
build            封装docker-compose build 命令,构建或者重新构建服务。范例：xii build -php


                 ---------以下为快速进入某个容器内部的快捷方法-------

go               快速进入容器, 直接xii go 将罗列所有存活容器选择，可以xii go -关键字加以过滤，缩小选择范围.
                 部分容器的入口不是/bin/sh，则会进入失败
nginx            快速进入nginx容器
php              快速进入*php*容器，多项目时需选择
mysql            快速进入mysql*容器，多项目时需选择
mongodb          快速进入*mongodb*容器，多项目时需选择
redis            快速进入*redis*容器，多项目时需选择
supervisor       快速进入*supervisor*容器，多项目时需选择
elasticsearch    快速进入*elasticsearch*容器，多项目时需选择
kibana           快速进入*elasticsearch*容器，多项目时需选择
logstash         快速进入*logstash*容器，多项目时需选择
node             快速进入*node*容器，多项目时需选择
golang           快速进入*golang*容器，多项目时需选择

```

## 二期功能如下
- [ ] 安装任意位置
- [ ] 自由选装组件
- [ ] 多语言
- [ ] 增加对常用的命令支持，比如mysql，sftp等