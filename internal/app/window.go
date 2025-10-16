package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MainWindow struct {
	window fyne.Window
	app    fyne.App
}

func NewMainWindow(app fyne.App) *MainWindow {
	window := app.NewWindow("Termius Go")
	window.Resize(fyne.NewSize(800, 600))
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
	title := widget.NewLabel("Termius Go - SSH клиент")
	title.Alignment = fyne.TextAlignCenter

	connectBtn := widget.NewButton("Подключиться к серверу", mw.onConnect)
	settingsBtn := widget.NewButton("Настройки", mw.onSettings)
	aboutBtn := widget.NewButton("О программе", mw.onAbout)

	content := container.NewVBox(
		title,
		widget.NewSeparator(),
		connectBtn,
		settingsBtn,
		aboutBtn,
	)

	return content
}

func (mw *MainWindow) onConnect() {
}

func (mw *MainWindow) onSettings() {
}

func (mw *MainWindow) onAbout() {
}
