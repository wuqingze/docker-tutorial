# 使用官方的Go镜像作为基础镜像
FROM golang:1.17-alpine as builder

# 设置工作目录
WORKDIR /app

# 将Go代码复制到工作目录中
COPY server.go .

# 编译Go HTTP服务器程序
RUN go build -o hello-world-server server.go

# 使用一个新的轻量级Alpine镜像
FROM alpine:latest

# 复制由上一个阶段生成的可执行文件到新的镜像中
COPY --from=builder /app/hello-world-server /app/

# 将index.html文件复制到新的镜像中
COPY index.html /app/

# 设置工作目录
WORKDIR /app

# 暴露HTTP服务器的端口
EXPOSE 8080

# 设置启动容器时运行的命令
CMD ["/app/hello-world-server"]
