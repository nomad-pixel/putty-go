package app

import (
	"fyne.io/fyne/v2"
)

type MainWindow struct {
	window fyne.Window
	app    fyne.App
}

func NewMainWindow(app fyne.App) *MainWindow {
	window := app.NewWindow("Putty Go")
	window.Resize(fyne.NewSize(1000, 700))
	window.CenterOnScreen()

	return &MainWindow{
		window: window,
		app:    app,
	}
}

func (mw *MainWindow) Show() {
	mw.window.Show()
}

func (mw *MainWindow) SetContent(content fyne.CanvasObject) {
	mw.window.SetContent(content)
}

func (mw *MainWindow) CreateMainContent() fyne.CanvasObject {
	navigation := NewNavigation(mw)
	return navigation.GetContent()
}
