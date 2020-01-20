#!/bin/sh
#设置host.docker.internal域名访问宿主机网络
ip -4 route list match 0/0 | awk '{print $3" host.docker.internal"}' >> /etc/hosts
#设置时区
cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
echo 'Asia/Shanghai' >/etc/timezone

#设置编码为中文
yum -y install kde-l10n-Chinese glibc-common
LANG="zh_CN.UTF-8"
echo "export LC_ALL=zh_CN.UTF-8"  >>  /etc/profile

./life coolq