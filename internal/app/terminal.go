package app

import (
	"fmt"
	"io"
	"putty-go/internal/ssh"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Terminal struct {
	window    fyne.Window
	conn      *ssh.Connection
	output    *widget.RichText
	input     *widget.Entry
	container fyne.CanvasObject
	hostname  string
	username  string
	prompt    string
}

func NewTerminal(window fyne.Window, conn *ssh.Connection, hostname, username string) *Terminal {
	terminal := &Terminal{
		window:   window,
		conn:     conn,
		hostname: hostname,
		username: username,
		prompt:   fmt.Sprintf("%s@%s:~$ ", username, hostname),
	}

	terminal.createUI()
	return terminal
}

func (t *Terminal) createUI() {
	t.output = widget.NewRichText()
	t.output.Wrapping = fyne.TextWrapOff
	t.output.ParseMarkdown("")

	t.input = widget.NewEntry()
	t.input.SetPlaceHolder(t.prompt)
	t.input.OnSubmitted = func(cmd string) {
		t.sendCommand(cmd)
		t.input.SetText("")
	}

	disconnectBtn := widget.NewButton("Disconnect", func() {
		t.disconnect()
	})
	disconnectBtn.Importance = widget.DangerImportance

	statusBar := widget.NewLabel("SSH Terminal - " + t.username + "@" + t.hostname)
	statusBar.Alignment = fyne.TextAlignCenter

	t.container = container.NewBorder(
		statusBar,
		container.NewHBox(t.input, disconnectBtn),
		nil,
		nil,
		t.output,
	)

	t.appendWelcome()
}

func (t *Terminal) appendWelcome() {
	welcome := fmt.Sprintf(`
**Welcome to SSH Terminal**

Connected to: **%s@%s**
Time: %s

Type commands and press Enter to execute them.
Type 'exit' to disconnect or use the Disconnect button.

---

`, t.username, t.hostname, time.Now().Format("2006-01-02 15:04:05"))

	t.appendOutput(welcome)
}

func (t *Terminal) sendCommand(command string) {
	if command == "" {
		return
	}

	if command == "exit" || command == "quit" {
		t.disconnect()
		return
	}

	t.appendOutput(fmt.Sprintf("**%s** %s\n", t.prompt, command))

	output, err := t.conn.ExecuteCommand(command)
	if err != nil {
		t.appendOutput(fmt.Sprintf("**Error:** %s\n", err.Error()))
		return
	}

	if strings.TrimSpace(output) != "" {
		t.appendOutput(fmt.Sprintf("```\n%s\n```\n", output))
	}
}

func (t *Terminal) appendOutput(text string) {
	currentText := t.output.String()
	t.output.ParseMarkdown(currentText + text)
}

func (t *Terminal) disconnect() {
	t.appendOutput("\n**Disconnecting...**\n")
	t.conn.Close()
	t.window.Close()
}

func (t *Terminal) GetContent() fyne.CanvasObject {
	return t.container
}

func (t *Terminal) StartInteractive() {
	session, err := t.conn.Client.NewSession()
	if err != nil {
		t.appendOutput(fmt.Sprintf("**Error creating session:** %s\n", err.Error()))
		return
	}
	defer session.Close()

	session.Stdout = &terminalWriter{t}
	session.Stderr = &terminalWriter{t}

	stdin, err := session.StdinPipe()
	if err != nil {
		t.appendOutput(fmt.Sprintf("**Error creating stdin pipe:** %s\n", err.Error()))
		return
	}

	if err := session.Shell(); err != nil {
		t.appendOutput(fmt.Sprintf("**Error starting shell:** %s\n", err.Error()))
		return
	}

	t.input.OnSubmitted = func(cmd string) {
		io.WriteString(stdin, cmd+"\n")
		t.input.SetText("")
	}

	session.Wait()
}

type terminalWriter struct {
	terminal *Terminal
}

func (tw *terminalWriter) Write(p []byte) (n int, err error) {
	text := string(p)
	tw.terminal.appendOutput(fmt.Sprintf("```\n%s\n```\n", text))
	return len(p), nil
}
