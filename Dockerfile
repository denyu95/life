#源镜像
FROM centos:latest
#作者
MAINTAINER DENYU "extrnight@126.com"
#设置工作目录
WORKDIR /root/go/src/github.com/denyu95/life
#将服务器的go工程代码加入到docker容器中
ADD . /root/go/src/github.com/denyu95/life
#安装go运行环境
RUN yum install -y go
#go构建可执行文件
RUN go build .
#暴露端口
EXPOSE 8080
#最终运行docker的命令
ENTRYPOINT ["./start.sh"]
#创建挂载点
VOLUME ["/root/go/src/github.com/denyu95/life/logs"]