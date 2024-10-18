FROM golang:1.23.1 AS development

WORKDIR /go/src/app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
#RUN set GOSUMDB=off
ENV CGO_ENABLED=0
RUN go mod tidy
RUN go build -o server .

# 两步可以用以下方式代替
# RUN export CGO_ENABLED=0 && go build -o app .
# RUN CGO_ENABLED=0 go build

# RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags='-w -extldflags "-static"' -a -installsuffix cgo -o c3alert .

# RUN go install -v ./...

# 如果你发现go build的二进制文件无法在alpine执行的问题：
# 后面发现网上有对比-tags netgo 和CGO_ENABLED，通过设置CGO_ENABLED=0环境变量，发现镜像启动也是没问题的。
# 有人猜测，golang默认使用cgo编译，但busybox中没有c相关的环境导致二进制文件无法正常运行。因此，将cgo编译关闭之后就没问题了。
# 个人测试：-tags netgo参数没有用

# FROM scratch AS production
FROM alpine:latest AS production
WORKDIR /root
COPY --from=development /go/src/app/server .
RUN chmod +x /root/server
EXPOSE 7777
ENTRYPOINT ["/root/server"]
#CMD ["./app"] # CMD可以被覆盖

# docker build -t go-sizeof-tips .
# docker run -it --rm -p 7777:7777 --name go-sizeof-tips go-sizeof-tips