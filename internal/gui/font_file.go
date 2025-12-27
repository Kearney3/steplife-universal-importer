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

// loadIconFromEmbedded 从文件系统加载图标（当不使用嵌入资源时）
func loadIconFromEmbedded() fyne.Resource {
	// 从 internal/gui/resources 路径加载图标
	res, err := fyne.LoadResourceFromPath("./internal/gui/resources/icon.png")
	if err != nil {
		fyne.LogError("Error loading icon from file system", err)
		return nil
	}
	return res
}
