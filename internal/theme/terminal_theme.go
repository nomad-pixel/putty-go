package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type TerminalTheme struct {
	fyne.Theme
}

func NewTerminalTheme() *TerminalTheme {
	return &TerminalTheme{
		Theme: theme.DefaultTheme(),
	}
}

func (t *TerminalTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{R: 26, G: 26, B: 46, A: 255}
	case theme.ColorNameForeground:
		return color.RGBA{R: 160, G: 230, B: 160, A: 255}
	case theme.ColorNameInputBackground:
		return color.RGBA{R: 26, G: 26, B: 46, A: 255}
	case theme.ColorNameInputBorder:
		return color.RGBA{R: 160, G: 230, B: 160, A: 255}
	case theme.ColorNameButton:
		return color.RGBA{R: 46, G: 46, B: 66, A: 255}
	case theme.ColorNameHover:
		return color.RGBA{R: 66, G: 66, B: 86, A: 255}
	case theme.ColorNamePressed:
		return color.RGBA{R: 36, G: 36, B: 56, A: 255}
	case theme.ColorNameFocus:
		return color.RGBA{R: 230, G: 160, B: 230, A: 255}
	case theme.ColorNameSelection:
		return color.RGBA{R: 230, G: 160, B: 230, A: 100}
	case theme.ColorNameScrollBar:
		return color.RGBA{R: 160, G: 230, B: 160, A: 255}
	case theme.ColorNameShadow:
		return color.RGBA{R: 0, G: 0, B: 0, A: 150}
	default:
		return t.Theme.Color(name, variant)
	}
}

func (t *TerminalTheme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Monospace {
		return theme.DefaultTheme().Font(style)
	}
	return theme.DefaultTheme().Font(style)
}

func (t *TerminalTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *TerminalTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
