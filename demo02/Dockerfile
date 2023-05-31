# 使用基础镜像 alpine
FROM alpine:latest

# 将编译好的二进制文件复制到容器内的 /app 目录
COPY main /app/main

# 设置工作目录为 /app
WORKDIR /app

# 设置启动容器时运行的命令
CMD ["/app/main"]
