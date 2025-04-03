package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type myTheme struct{}

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.White
	}

	return theme.DefaultTheme().Color(name, variant)
}

// Icon implements fyne.Theme.
func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	panic("unimplemented")
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

// 确保实现所有接口方法
var _ fyne.Theme = (*myTheme)(nil)

// 导出主题实例
var CustomTheme = &myTheme{}
