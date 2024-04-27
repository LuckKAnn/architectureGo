FROM golang:latest

ENV GO111MODULE=on \
        GOPROXY=https://goproxy.cn,direct
WORKDIR /Users/liukunkun/Desktop/dev/golang/webGinDemo
# 这里copy . 代表的是什么呢
# 把这个文件下的东西copy到指定的文件夹里面吗
COPY . /Users/liukunkun/Desktop/dev/golang/webGinDemo
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./ginDemo"]