//go:build !embed_font
// +build !embed_font

package gui

import (
	"fyne.io/fyne/v2"
)

// loadCustomFontFromEmbedded 从文件系统加载字体（当不使用嵌入字体时）
func loadCustomFontFromEmbedded() fyne.Resource {
	// 回退到文件系统路径
	return loadCustomFont("./resource/MiSans-Regular.otf")
}

