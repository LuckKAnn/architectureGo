FROM golang:latest

ENV GO111MODULE=on \
        GOPROXY=https://goproxy.cn,direct
WORKDIR /Users/liukunkun/Desktop/dev/golang/webGinDemo
COPY . /Users/liukunkun/Desktop/dev/golang/webGinDemo
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]