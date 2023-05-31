# 使用轻量级Java基础镜像
FROM adoptopenjdk/openjdk11:alpine-jre

# 设置工作目录
WORKDIR /app

# 复制Java源文件到容器中的工作目录
COPY HelloWorld.class /app/HelloWorld.class

# 设置容器启动时执行的命令
CMD ["java", "HelloWorld"]
