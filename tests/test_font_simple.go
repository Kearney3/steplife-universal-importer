package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 模拟setupChineseFont的逻辑
	fontPath := "./resource/MiSans-Regular.otf"

	// 获取绝对路径
	if absPath, err := filepath.Abs(fontPath); err == nil {
		fontPath = absPath
	}

	// 设置环境变量
	os.Setenv("FYNE_FONT", fontPath)
	os.Setenv("LANG", "zh_CN.UTF-8")
	os.Setenv("LC_ALL", "zh_CN.UTF-8")

	// 检查字体文件
	if _, err := os.Stat(fontPath); err == nil {
		fmt.Printf("✅ 字体文件存在: %s\n", fontPath)
	} else {
		fmt.Printf("❌ 字体文件不存在: %s\n", fontPath)
	}

	fmt.Printf("FYNE_FONT: %s\n", os.Getenv("FYNE_FONT"))
	fmt.Printf("LANG: %s\n", os.Getenv("LANG"))
}