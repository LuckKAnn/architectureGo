FROM golang:latest

ENV GO111MODULE=on \
        GOPROXY=https://goproxy.cn,direct
WORKDIR /app
# 这里copy . 代表的是什么呢
# 把这个文件下的东西copy到指定的文件夹里面吗
COPY . /app
RUN go build /app

EXPOSE 8080
ENTRYPOINT ["./ginDemo"]