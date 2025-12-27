//go:build embed_font
// +build embed_font

package gui

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed resources/MiSans-Regular.otf
var embeddedFont []byte

// loadCustomFontFromEmbedded 从嵌入资源加载字体
func loadCustomFontFromEmbedded() fyne.Resource {
	if len(embeddedFont) == 0 {
		fyne.LogError("Embedded font is empty", nil)
		return nil
	}
	res := fyne.NewStaticResource("MiSans-Regular.otf", embeddedFont)
	return res
}
