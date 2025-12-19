@echo off
REM Windows构建脚本 - 一生足迹数据导入器

echo 构建一生足迹数据导入器...

REM 清理之前的构建
if exist main.exe del main.exe
if exist main del main

REM 构建
go build -o main.exe ./cmd

if %errorlevel% equ 0 (
    echo 构建完成: main.exe
    echo 构建成功！
) else (
    echo 构建失败！
    pause
    exit /b 1
)

pause
