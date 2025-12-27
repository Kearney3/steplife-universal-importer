@echo off
REM Windows构建脚本 - 一生足迹数据导入器

setlocal enabledelayedexpansion

REM 解析命令行参数
set EMBED_FONT=true
:parse_args
if "%~1"=="" goto end_parse
if "%~1"=="--no-embed-font" (
    set EMBED_FONT=false
    shift
    goto parse_args
)
if "%~1"=="-h" goto show_help
if "%~1"=="--help" goto show_help
if "%~1"=="" goto end_parse
echo 未知参数: %~1
echo 使用 %~nx0 --help 查看帮助信息
exit /b 1
:show_help
echo 用法: %~nx0 [选项]
echo.
echo 选项:
echo   --no-embed-font    不嵌入字体，使用文件系统路径加载字体
echo   -h, --help         显示此帮助信息
echo.
echo 默认情况下会嵌入字体到可执行文件中。
exit /b 0
:end_parse

echo 构建一生足迹数据导入器...

REM 清理之前的构建
if exist main.exe del main.exe
if exist main del main

REM 构建参数
REM -trimpath: 移除文件系统中的路径信息，使构建更可重现
REM -ldflags="-s -w": 减小二进制文件大小
REM   -s: 省略符号表和调试信息
REM   -w: 省略DWARF符号表

REM 根据参数决定是否使用 build tag
set BUILD_TAGS=
if "!EMBED_FONT!"=="true" (
    set BUILD_TAGS=-tags=embed_font
    echo 字体嵌入: 启用（字体将嵌入到可执行文件中）
) else (
    echo 字体嵌入: 禁用（将从文件系统加载字体）
)

REM 构建
go build %BUILD_TAGS% -trimpath -ldflags="-s -w" -o main.exe ./cmd

if %errorlevel% equ 0 (
    echo 构建完成: main.exe
    dir main.exe
    echo 构建成功！
) else (
    echo 构建失败！
    pause
    exit /b 1
)

pause
