# 下载的是完整的golang编译和运行时环境，还包括gcc、build之类的工具
FROM golang:latest


ENV GO111MODULE=on \
        GOPROXY=https://goproxy.cn,direct
WORKDIR /app
# 这里copy . 代表的是什么呢
# 把这个文件下的东西copy到指定的文件夹里面吗
COPY . /app
RUN go build .

# 声明运行时容器提供的服务端口，可以修改
EXPOSE 8000

# 容器启动程序及参数
ENTRYPOINT ["./ginDemo"]