package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Navigation struct {
	sidebar     fyne.CanvasObject
	content     *container.Split
	mainWindow  *MainWindow
	currentPage fyne.CanvasObject
}

func NewNavigation(mainWindow *MainWindow) *Navigation {
	nav := &Navigation{
		mainWindow: mainWindow,
	}

	nav.createSidebar()
	nav.createContent()

	return nav
}

func (n *Navigation) createSidebar() {
	hostsBtn := widget.NewButton("Hosts", func() {
		n.showHostsPage()
	})

	settingsBtn := widget.NewButton("Settings", func() {
		n.showSettingsPage()
	})

	aboutBtn := widget.NewButton("About", func() {
		n.showAboutPage()
	})

	n.sidebar = container.NewVBox(
		widget.NewLabel("Navigation"),
		widget.NewSeparator(),
		hostsBtn,
		settingsBtn,
		aboutBtn,
	)
}

func (n *Navigation) createContent() {
	hostsPage := n.createHostsPage()
	n.currentPage = hostsPage

	n.content = container.NewHSplit(n.sidebar, n.currentPage)
	n.content.SetOffset(0.2)
}

func (n *Navigation) createHostsPage() fyne.CanvasObject {
	title := widget.NewLabel("SSH Hosts")
	title.Alignment = fyne.TextAlignCenter

	addHostBtn := widget.NewButton("Add Host", func() {
		n.showAddHostDialog()
	})

	hostsList := widget.NewList(
		func() int { return 0 },
		func() fyne.CanvasObject {
			return widget.NewLabel("No hosts")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
		},
	)

	content := container.NewVBox(
		title,
		widget.NewSeparator(),
		addHostBtn,
		hostsList,
	)

	return content
}

func (n *Navigation) showHostsPage() {
	hostsPage := n.createHostsPage()
	n.currentPage = hostsPage
	n.content = container.NewHSplit(n.sidebar, n.currentPage)
	n.content.SetOffset(0.2)
}

func (n *Navigation) showSettingsPage() {
	title := widget.NewLabel("Settings")
	title.Alignment = fyne.TextAlignCenter

	content := container.NewVBox(
		title,
		widget.NewSeparator(),
		widget.NewLabel("Settings page coming soon..."),
	)

	n.currentPage = content
	n.content = container.NewHSplit(n.sidebar, n.currentPage)
	n.content.SetOffset(0.2)
}

func (n *Navigation) showAboutPage() {
	title := widget.NewLabel("About")
	title.Alignment = fyne.TextAlignCenter

	content := container.NewVBox(
		title,
		widget.NewSeparator(),
		widget.NewLabel("Putty Go v1.0.0"),
		widget.NewLabel("SSH Client for Go"),
	)

	n.currentPage = content
	n.content = container.NewHSplit(n.sidebar, n.currentPage)
	n.content.SetOffset(0.2)
}

func (n *Navigation) showAddHostDialog() {
	hostnameEntry := widget.NewEntry()
	hostnameEntry.SetPlaceHolder("Hostname or IP")

	portEntry := widget.NewEntry()
	portEntry.SetPlaceHolder("Port (default: 22)")
	portEntry.SetText("22")

	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Username")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Password")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Hostname", Widget: hostnameEntry},
			{Text: "Port", Widget: portEntry},
			{Text: "Username", Widget: usernameEntry},
			{Text: "Password", Widget: passwordEntry},
		},
		OnSubmit: func() {
			n.mainWindow.window.Close()
		},
		OnCancel: func() {
			n.mainWindow.window.Close()
		},
	}

	dialog := widget.NewModalPopUp(form, n.mainWindow.window.Canvas())
	dialog.Resize(fyne.NewSize(400, 300))
	dialog.Show()
}

func (n *Navigation) GetContent() fyne.CanvasObject {
	return n.content
}
