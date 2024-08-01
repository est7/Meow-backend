FROM golang:1.22.4-alpine AS builder

# 添加必要的工具
RUN apk add --no-cache git make bash ca-certificates tzdata

# 设置环境变量
ENV PROJECT_DIR=/app \
    BINARY_NAME=meow_backend_artifact \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn" \
    TZ=Asia/Shanghai \
    APP_ENV=docker

# 创建并设置工作目录
RUN mkdir -p ${PROJECT_DIR}
WORKDIR ${PROJECT_DIR}

# 复制依赖文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 确认 Makefile 的内容
RUN cat Makefile

# 编译应用
RUN make build


# 确保构建产物存在
RUN ls -l ${PROJECT_DIR}/${BINARY_NAME}

# 第二阶段：运行环境
FROM alpine:latest

# 添加必要的工具
RUN apk add --no-cache git make bash ca-certificates tzdata curl

# 设置二进制文件名称
ENV BINARY_NAME=meow_backend_artifact

# 从 builder 阶段复制编译好的二进制文件
COPY --from=builder /app/${BINARY_NAME} /bin/${BINARY_NAME}

# 设置工作目录为 /bin
WORKDIR /bin

# 确保文件被正确复制并设置执行权限
RUN ls -la /bin && \
    file /bin/${BINARY_NAME} && \
    chmod +x /bin/${BINARY_NAME}

# 暴露应用端口
EXPOSE 8080

# 运行应用
CMD ["/bin/sh", "-c", "echo Binary name is: ${BINARY_NAME} && ls -l /bin/${BINARY_NAME} && exec /bin/${BINARY_NAME}"]
