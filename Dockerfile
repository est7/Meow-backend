FROM golang:1.22.4-alpine AS builder
# 最新的 alpine 镜像没有一些工具，例如（`git` 和 `bash`）。
# 将 git、bash 和 openssh 添加到镜像中
RUN apk add --no-cache git make bash ca-certificates tzdata

# 镜像设置必要的环境变量
# 工作目录名称 app

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn" \
    TZ=Asia/Shanghai \
    APP_ENV=docker

# 移动到工作目录
RUN mkdir -p ${PROJECT_DIR}

# 设置工作目录# 使用官方的 Go 镜像作为基础镜像
WORKDIR ${PROJECT_DIR}

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod go.sum ./
# 下载依赖
RUN go mod download

# 复制项目的源代码
COPY . .

# 编译应用
RUN make build
#RUN go build -o main .

# 使用轻量级的 alpine 镜像作为运行环境,这里因为国内网络问题，使用阿里云的镜像
FROM alpine:latest

# 安装 curl 用于健康检查
RUN apk --no-cache add curl

WORKDIR /root/

# 从 builder 阶段复制编译好的二进制文件
COPY --from=builder /app/main .

# 暴露应用端口
EXPOSE 8080

# 运行应用
CMD ["./main"]