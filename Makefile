# Makefile
# init environment variables
export PATH        := $(shell go env GOPATH)/bin:$(PATH)
export GOPATH      := $(shell go env GOPATH)
export GO111MODULE := on

print-env:
	@echo "PATH: $(PATH)"
	@echo "GOROOT: $(GOROOT)"
	@echo "GOPATH: $(GOPATH)"
	@which go
	@go version

# 定义变量
BINARY_NAME=meow_backend_artifact
DOCKER_COMPOSE=docker-compose

# 默认目标
all: build

# 构建 Go 应用
build:
	@echo "Building..."
	@go build -o $(BINARY_NAME) main.go

# 运行应用
run: build
	@echo "Running..."
	@./$(BINARY_NAME)

# 清理构建产物
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@go clean

# 运行测试
test:
	@echo "Testing..."
	@go test ./...

# Docker 相关命令
docker-build:
	@echo "Building Docker image..."
	@docker build -t $(BINARY_NAME) .

docker-run: docker-build
	@echo "Running Docker container..."
	@docker run -p 8080:8080 $(BINARY_NAME)

# Docker Compose 命令
up:
	@echo "Starting services with Docker Compose..."
	@$(DOCKER_COMPOSE) up -d

down:
	@echo "Stopping services..."
	@$(DOCKER_COMPOSE) down

logs:
	@echo "Showing logs..."
	@$(DOCKER_COMPOSE) logs -f


IMAGE_NAME = your_image_name
CONTAINER_NAME = your_container_name

clean:
	@echo "Cleaning up..."
	-docker stop $(CONTAINER_NAME)
	-docker rm $(CONTAINER_NAME)
	-docker rmi $(IMAGE_NAME)

# 帮助信息
help:
	@echo "Available commands:"
	@echo "  make build      - Build the Go application"
	@echo "  make run        - Run the application"
	@echo "  make clean      - Remove build artifacts"
	@echo "  make test       - Run tests"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-run   - Run Docker container"
	@echo "  make up         - Start services with Docker Compose"
	@echo "  make down       - Stop services"
	@echo "  make logs       - Show service logs"

.PHONY: all build run clean test docker-build docker-run up down logs help clean