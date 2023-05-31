# 使用基础镜像 alpine
FROM alpine:latest

# 将编译好的二进制文件复制到容器内的 /app 目录
COPY app /app/hello_world

# 设置工作目录为 /app
WORKDIR /app

# 设置容器启动时的默认命令为运行 hello_world 二进制文件
CMD ["./hello_world"]
