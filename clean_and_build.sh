#!/bin/bash

# 使用方法：
#给脚本添加执行权限：chmod +x clean_and_build.sh
#运行脚本：./clean_and_build.sh
#使用 Makefile

# 设置变量
IMAGE_NAME="your_image_name"
CONTAINER_NAME="your_container_name"

# 停止并删除容器（如果存在）
echo "Stopping and removing container..."
docker stop $CONTAINER_NAME 2>/dev/null
docker rm $CONTAINER_NAME 2>/dev/null

# 删除镜像
echo "Removing image..."
docker rmi $IMAGE_NAME 2>/dev/null

# 重新构建镜像
echo "Rebuilding image..."
docker build -t $IMAGE_NAME .

# 可选：重新运行容器
# echo "Running new container..."
# docker run -d --name $CONTAINER_NAME $IMAGE_NAME

echo "Clean and build process completed."