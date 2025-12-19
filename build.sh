#!/bin/bash

# 构建脚本 - 一生足迹数据导入器

echo "构建一生足迹数据导入器..."

# 检测操作系统
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    OS="linux"
elif [[ "$OSTYPE" == "darwin"* ]]; then
    OS="darwin"
elif [[ "$OSTYPE" == "msys" ]] || [[ "$OSTYPE" == "win32" ]]; then
    OS="windows"
else
    echo "不支持的操作系统: $OSTYPE"
    exit 1
fi

# 检测架构
ARCH=$(uname -m)
if [[ "$ARCH" == "x86_64" ]]; then
    ARCH="amd64"
elif [[ "$ARCH" == "arm64" ]] || [[ "$ARCH" == "aarch64" ]]; then
    ARCH="arm64"
else
    echo "不支持的架构: $ARCH"
    exit 1
fi

echo "目标平台: $OS/$ARCH"

# 清理之前的构建
rm -f main main.exe

# 构建
if [[ "$OS" == "windows" ]]; then
    GOOS=$OS GOARCH=$ARCH go build -o main.exe ./cmd
    echo "构建完成: main.exe"
else
    GOOS=$OS GOARCH=$ARCH go build -o main ./cmd
    echo "构建完成: main"
fi

echo "构建成功！"
