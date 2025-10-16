package main

import (
	"log"
	"putty-go/internal/app"
	"putty-go/internal/config"

	fyneapp "fyne.io/fyne/v2/app"
)

func main() {
	myApp := fyneapp.NewWithID("com.termius.go")

	cfg := config.DefaultConfig()
	if err := cfg.Load(); err != nil {
		log.Printf("Ошибка загрузки конфигурации: %v", err)
	}

	mainWindow := app.NewMainWindow(myApp)

	content := mainWindow.CreateMainContent()
	mainWindow.SetContent(content)

	mainWindow.Show()
	myApp.Run()
}
