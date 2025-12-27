//go:build embed_font
// +build embed_font

package gui

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed resources/MiSans-Regular.otf
var embeddedFont []byte

//go:embed resources/icon.png
var embeddedIcon []byte

// loadCustomFontFromEmbedded 从嵌入资源加载字体
func loadCustomFontFromEmbedded() fyne.Resource {
	if len(embeddedFont) == 0 {
		fyne.LogError("Embedded font is empty", nil)
		return nil
	}
	res := fyne.NewStaticResource("MiSans-Regular.otf", embeddedFont)
	return res
}

// loadIconFromEmbedded 从嵌入资源加载图标
func loadIconFromEmbedded() fyne.Resource {
	if len(embeddedIcon) == 0 {
		fyne.LogError("Embedded icon is empty", nil)
		return nil
	}
	res := fyne.NewStaticResource("icon.png", embeddedIcon)
	return res
}
