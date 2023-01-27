FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
GOPROXY=https://goproxy.cn,direct \
CGO_ENABLED=0 \
GOOS=linux \
GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件 bubble
RUN go build -o bluebell .

###################
# 接下来创建一个小镜像
###################
FROM debian:stretch-slim

COPY ./wait-for.sh /
COPY ./conf /conf

# 从builder镜像中把可执行文件拷贝到当前目录
COPY --from=builder /build/bluebell /

RUN set -eux \
    && apt-get update \
    && apt-get install -y --no-install-recommends netcat \
    && chmod 755 wait-for.sh

# 需要运行的命令
#ENTRYPOINT ["/bluebell", "conf/config.ini"]