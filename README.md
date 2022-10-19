
# 最新changelog（1-2条）
2022-10-19 增加对arm、arm64 架构的支持。

# 群聊沟通
Tg: [https://t.me/xii_app](https://t.me/xii_app)

# 简介

<font color=red> 详细文档请查看： </font>
[https://xii.app](https://xii.app)


Docker 化的 lnmp 环境及更多软件包，同时提供助手程序 xii 用来快捷管理。

> **注：** xii 是罗马数字 12 的意思，正好有域名 xii.app ，就很随意的用了。 这会征求名。



## 内置软件(Docker 镜像)清单
| 软件          | 具体版本                                 | 备注                        |
| ------------- | ---------------------------------------- | --------------------------- |
| **PHP**       | php7.3 , php7.4, php8.0, php8.1, php8.2  |                             |
| **Web服务器** | nginx                                    | 内嵌acme.sh,用于ssl证书申请 |
| **数据库**    | mysql5.7, mysql8.0 , mysql5.6 , mysql5.5 |                             |
| **Nosql**     | mongoldb , redis                         |                             |
| **缓存**      | memcached                                |                             |
| **消息队列**  | rabitmq                                  |                             |
| **其他**      | supervisor                               |                             |
| **工具**      | phpmyadmin                               |                             |


**注 1：** 可自行轻松扩展软件清单

**注 2：** php 支持快捷安装扩展，<u>已默认安装 compose</u>，php 扩展等请查看左边 php 节点说明。 nginx 的配置已做优化处理，默认自带 acme.sh，可免费申请 ssl 证书。phpmyadmin 实际运行体积大改 450Mb，不太建议启用。

**注3：** Mysql尽可能建议使用5.7或以上版本，能获得更多的性能。 使用前请查看左边mysql说明了解端口、导出导入数据等信息。


## 助手程序 Xii

主要提供了易用性功能，具体如下：

1. 便捷的网站增删查，使用 xii vhost add 可以创建安装，安装 nginx 的优化配置，申请免费的 ssl 证书等（该功能借鉴自 lnmp 套件）。
2. 便捷的 docker 镜像管理功能， 使用 `` xii init` 可以随时随地增删 docker 镜像。
3. 封装了 docker-compose 的大部分命令，docker-compose 平时要的时候，要么需要进入到 docker-compose.yml 所在的目录，要么需要带一大串参数。 而且任何危险操作都没有二次确认，这导致不是很熟悉的人容易误操作。 xii 的封装解决了这些问题。可以在任何目录下执行，危险动作会有说明或二次确认。
4. 封装了 docker 的几个命令，主要也是危险动作二次确认或便捷操作。
5. 后续拟增加的功能：升级功能，继续提供一些便捷操作组合等，具体参照“路线图”


# 安装

**当前仅在 linux 和 mac 测试完全。**

注：暂没在 window 下充分测试完，项目的快捷操作严重依赖 xii 助手，所以如需要 window 下使用，建议使用 yeszao/dnmp 这个开源项目(本项目的配置项大量复用了它的代码段，并进行优化配置或提升版本等)。

## 一键安装

<font color=red>海外服务器：</font>

```sh
wget -c https://raw.githubusercontent.com/xiiapp/xii/main/script/install.sh && chmod +x install.sh && ./install.sh
```

<font color=red>境内服务器(加速)：</font>

```sh
wget -c https://raw.githubusercontent.com/xiiapp/xii/main/script/install.sh && chmod +x install.sh && ./install.sh china
```

**注：** <u>当前仅支持 mac 和 linux，暂不支持 win 系统一键安装。</u>
**注：** <u> arm 或者arm64 用户，请在一键安装后面追加一个参数，` arm` 或` arm64`</u>


## 手动安装

### mac 系统

1. 确保系统安装好 docker 和 docker-compose，建议直接安装 docker 官方的 docker-desktop。

2. 下载安装包,下载地址二选一

   > 苹果M1/M2芯片用户请下载 arm64 版本
   > [https://github.com/xiiapp/xii/raw/main/release/xii_mac_arm.zip ](https://github.com/xiiapp/xii/raw/main/release/xii_mac_arm.zip)

   > 苹果Intel芯片用户请下载 arm64 版本
   > [https://github.com/xiiapp/xii/raw/main/release/xii_mac.zip ](https://github.com/xiiapp/xii/raw/main/release/xii_mac.zip)


3. 解压安装后后，手动执行 `chmod +x manual.sh && ./manual.sh` 完成安装。

4. 注意检测 docker 是否启动，在 docker 启动的情况下，可以执行后续命令。

### Linux 系统

1. 确保系统安装好 docker 和 docker-compose。

2. 下载安装包,下载地址二选一

   > Linux 版本
   > [https://github.com/xiiapp/xii/raw/main/release/xii_linux.zip ](https://github.com/xiiapp/xii/raw/main/release/xii_linux.zip)
   > 
   > Linux arm64 版本
   > [https://github.com/xiiapp/xii/raw/main/release/xii_linux_arm64.zip ](https://github.com/xiiapp/xii/raw/main/release/xii_linux_arm64.zip)
   >
   > Linux arm 版本
   > [https://github.com/xiiapp/xii/raw/main/release/xii_linux_arm.zip ](https://github.com/xiiapp/xii/raw/main/release/xii_linux_arm.zip)


3. 解压安装后后，手动执行 `chmod +x manual.sh && ./manual.sh` 完成安装。

4. 注意检测 docker 是否启动，在 docker 启动的情况下，可以执行后续命令。

# 卸载





## 快捷卸载

```sh
wget -c https://raw.githubusercontent.com/xiiapp/xii/main/script/uninstall.sh && chmod +x uninstall.sh && ./uninstall.sh
```



## 手动卸载

- 备份好所需数据，网站数据放在```www目录```，配置放在```env/容器类型```, 容器产生的数据，比如msyql，一般在```data/容器名``` 下

- 执行 ```xii down```  停止所有容器

- 执行 ```xii rmall``` 删除所有容器、镜像、卷

- Linux用户执行 ``` rm -rf /home/xii ```，Mac用户执行 ``` rm -rf ~/xii``` 删除所有数据

- 执行 ``` rm  -f /usr/local/bin/xii ```  和 ``` rm  -f /usr/local/bin/xxi``` 删除软链接


# 快速选择组件

无论什么时候，只要你想变更 docker 容器，都可以执行以下命令进行组件的变更。

```
xii init
```

**演示动画**

<script id="asciicast-ag7Woq3p2wc9hsFnbI4c1EzM9" src="https://asciinema.org/a/ag7Woq3p2wc9hsFnbI4c1EzM9.js" async></script>

**友情提示，一个可能更好的操作流程**

第一次建议执行 `xii up -d` 然后再执行 `xii ps` 看下你的容器是否都起来了。如果没起来，可以再执行一次 `xii up ` 或 `xii up -容器名` 来看一下报了什么错误



# 网站管理



## 添加网站或nginx反代

> xii vhost add



**注：** 添加网站或者任意nginx反代的过程中，会提示很多选项，根据需要创建即可，此外过程中，会提示是否申请免费的SSL证书。

**注：**当前不会创建mysql数据库，拟后续版本推出。



**演示:**

<script id="asciicast-528842" src="https://asciinema.org/a/528842.js" async></script>

**注：**国内某些云服务器间歇性屏蔽掉github的访问，有可能会导致用来生成证书的acme.sh无法安装上，xii vhost如果遇到这种情况，会尝试自己安装一次，还不行会将安装命令显示出来，并提示用户复制黏贴后自己执行几次安装。



## 删除网站或反代

>  xii vhost del



注：其本质是检查nginx的配置目录vhost里的配置文件，一个文件就是一个网站

**演示：**

<script id="asciicast-528843" src="https://asciinema.org/a/528843.js" async></script>



## 查看网站列表(仅查看)

> xii vhost list





## 修改网站配置

解析匹配关系，暂不提供。

请自行编辑 ```env/nginx/vhost/域名.conf```文件

# 有关 xii 的操作问答

## 问题 1:如何重新编译镜像，即修改 docker-compose.yml ，dockerfile，env 文件等，要如何让他生效？

**场景：**服务器已经在跑了一个 nginx 容器（也就是说这个容器被 `docker-compose up` 或者 `xii up` 过一次了）， 此时修改`docker-compose.yml`或者`dockerfile`文件、.env 文件后，要如何让他生效？

**关键理解：**

- docker-compose.yml 的的改变，只跟 2 个 docker-compose up 和 docker-compose build 有关，其他命令都无法让他生效
- 默认**<font color=blue>/data 、/www、/env 里的数据是不会被删除掉</font>**的。如有需要强制清理，请`docker-compose down --volumes`。

**<font color=blue>可以有几种如下操作来确保生效，但推荐使用`xii rebuild -镜像名`来快捷操作：</font>**

- **最粗暴做法，先 `xii down` 后 `xii up` 一次：**

```sh
  # 做法 1:使用原始命令，必须在 xii 的项目下执行
  docker-compose down
  docker-compose up -d

  # 做法 2: 使用 xii 助手，可以在任何目录下执行
  xii down
  xii up -d
```

注意：`docker-compose down` 会删除容器、镜像、网络、映射的卷（我们通过 dockerfile 里 volume 命令挂载上去的不会删除）。

- **标准流程**

```sh


    #做法 1:原始命令方式，必须在 docker-compose.yml 所在目录下执行
    docker stop 容器名 #1.先停止容器
    docker rm 容器名 #2.删除掉容器
    systemctl restart docker #重启 docker 服务（测试过，新版似乎无需，需要进一步测试，可以暂时跳过观察一下）
    docker-compose up -d --no-deps --build 容器名 #4.重新 docker up 一下

    # 做法 2:使用 xii 助手

    xii stop -容器名
    xii rm -容器名 #重启 docker，同做法 1 第 3 步一样 #在 up 之前，也可以直接先 xii build 一次，多了一条操作命令，
    xii up -d -容器名

    # 做法 3: 使用 xii rebuild 一步到位。强烈推荐这个，最简单
    xii rebuild -容器名

```

- **粗暴做法 2，适合同样不想手动删除已经建立的 container**

```sh
    # --force-recreate 即使配置或者镜像（images）没有改变也强制重启

    docker-compose up --force-recreate -d

    # 参数挺长的，推荐使用 xii rebuild -容器名

```

## 问题 2: 在 Linux 上非 root 安装，每次启动 docker 都需要增加 sudo 的解决方法：

```
sudo gpasswd -a ${USER} docker
```

执行完毕后，重新登陆一次 ssh 即可

## 问题 3: 国内使用 docker 慢

安装的时候，复制专用国内的安装脚本。其实就是 install.sh 后面加一个参数 china 就会默认使用国内源。


# 授权

**MIT**


## Contributors

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-13-orange.svg?style=flat-square)](#contributors)
<!-- ALL-CONTRIBUTORS-BADGE:END -->