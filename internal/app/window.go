package app

import (
	"putty-go/internal/models"
	"putty-go/internal/ssh"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type MainWindow struct {
	window       fyne.Window
	app          fyne.App
	hostManager  *models.HostManager
	hostsList    *widget.List
	selectedHost int
}

func NewMainWindow(app fyne.App) *MainWindow {
	window := app.NewWindow("Putty Go")
	window.Resize(fyne.NewSize(1000, 700))
	window.CenterOnScreen()

	return &MainWindow{
		window:       window,
		app:          app,
		hostManager:  models.NewHostManager(),
		selectedHost: -1,
	}
}

func (mw *MainWindow) Show() {
	mw.window.Show()
}

func (mw *MainWindow) SetContent(content fyne.CanvasObject) {
	mw.window.SetContent(content)
}

func (mw *MainWindow) CreateMainContent() fyne.CanvasObject {
	title := widget.NewLabel("SSH Hosts")
	title.Alignment = fyne.TextAlignCenter

	addHostBtn := widget.NewButton("Add Host", func() {
		mw.showAddHostDialog()
	})

	connectBtn := widget.NewButton("Connect", func() {
		mw.connectToSelectedHost()
	})

	mw.hostsList = widget.NewList(
		func() int { return len(mw.hostManager.GetHosts()) },
		func() fyne.CanvasObject {
			return widget.NewLabel("No hosts")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			hosts := mw.hostManager.GetHosts()
			if id < len(hosts) {
				obj.(*widget.Label).SetText(hosts[id].Name)
			}
		},
	)

	mw.hostsList.OnSelected = func(id widget.ListItemID) {
		mw.selectedHost = int(id)
	}

	content := container.NewVBox(
		title,
		widget.NewSeparator(),
		addHostBtn,
		connectBtn,
		mw.hostsList,
	)

	return content
}

func (mw *MainWindow) showAddHostDialog() {
	hostnameEntry := widget.NewEntry()
	hostnameEntry.SetPlaceHolder("Hostname or IP")

	portEntry := widget.NewEntry()
	portEntry.SetPlaceHolder("Port (default: 22)")
	portEntry.SetText("22")

	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Username")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Password")

	var dialog *widget.PopUp

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Hostname", Widget: hostnameEntry},
			{Text: "Port", Widget: portEntry},
			{Text: "Username", Widget: usernameEntry},
			{Text: "Password", Widget: passwordEntry},
		},
		OnSubmit: func() {
			if hostnameEntry.Text == "" {
				hostnameEntry.SetText("")
				hostnameEntry.SetPlaceHolder("Hostname is required!")
				return
			}
			if usernameEntry.Text == "" {
				usernameEntry.SetText("")
				usernameEntry.SetPlaceHolder("Username is required!")
				return
			}
			if passwordEntry.Text == "" {
				passwordEntry.SetText("")
				passwordEntry.SetPlaceHolder("Password is required!")
				return
			}

			port, err := strconv.Atoi(portEntry.Text)
			if err != nil || port <= 0 {
				port = 22
			}

			mw.hostManager.AddHost(hostnameEntry.Text, usernameEntry.Text, passwordEntry.Text, port)
			mw.hostsList.Refresh()
			dialog.Hide()
		},
		OnCancel: func() {
			dialog.Hide()
		},
	}

	dialog = widget.NewModalPopUp(form, mw.window.Canvas())
	dialog.Resize(fyne.NewSize(400, 300))
	dialog.Show()
}

func (mw *MainWindow) connectToSelectedHost() {
	if mw.selectedHost == -1 {
		dialog.ShowInformation("No Selection", "Please select a host to connect to.", mw.window)
		return
	}

	hosts := mw.hostManager.GetHosts()
	if mw.selectedHost >= len(hosts) {
		dialog.ShowInformation("Error", "Selected host not found.", mw.window)
		return
	}

	host := hosts[mw.selectedHost]
	conn := ssh.NewConnection(host.Hostname, host.Username, host.Password, host.Port)

	progress := dialog.NewProgressInfinite("Connecting", "Connecting to "+host.Name, mw.window)
	progress.Show()

	go func() {
		err := conn.Connect()
		progress.Hide()

		if err != nil {
			dialog.ShowError(err, mw.window)
			return
		}

		err = conn.CreateSession()
		if err != nil {
			dialog.ShowError(err, mw.window)
			conn.Close()
			return
		}

		terminal := NewTerminal(mw.window, conn, host.Hostname, host.Username)
		terminalWindow := mw.app.NewWindow("SSH Terminal - " + host.Name)
		terminalWindow.SetContent(terminal.GetContent())
		terminalWindow.Resize(fyne.NewSize(900, 700))
		terminalWindow.CenterOnScreen()
		terminalWindow.Show()

		go terminal.StartInteractive()
	}()
}
