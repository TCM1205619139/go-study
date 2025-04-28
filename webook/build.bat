@echo off
REM 整理模块依赖
go mod tidy

REM 设置环境变量（Linux + ARM）
set GOOS=linux
set GOARCH=arm

REM 编译为 Linux ARM 架构的可执行文件（输出文件名为当前目录的名称）
go build -tags=k8s -o webook-exe ./internal/main.go

REM 构建 Docker 镜像
docker build -t flycash/webook:v0.0.1 .

REM 清除环境变量（可选）
set GOOS=
set GOARCH=

echo Build completed!
pause
