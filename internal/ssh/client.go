package ssh

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"golang.org/x/crypto/ssh"
)

type Connection struct {
	Hostname string
	Port     int
	Username string
	Password string
	Client   *ssh.Client
	Session  *ssh.Session
}

func NewConnection(hostname, username, password string, port int) *Connection {
	return &Connection{
		Hostname: hostname,
		Port:     port,
		Username: username,
		Password: password,
	}
}

func (c *Connection) Connect() error {
	config := &ssh.ClientConfig{
		User: c.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(c.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}

	address := net.JoinHostPort(c.Hostname, strconv.Itoa(c.Port))

	client, err := ssh.Dial("tcp", address, config)
	if err != nil {
		return fmt.Errorf("failed to connect: %v", err)
	}

	c.Client = client
	return nil
}

func (c *Connection) CreateSession() error {
	if c.Client == nil {
		return fmt.Errorf("not connected")
	}

	session, err := c.Client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %v", err)
	}

	c.Session = session
	return nil
}

func (c *Connection) ExecuteCommand(command string) (string, error) {
	if c.Session == nil {
		return "", fmt.Errorf("no active session")
	}

	output, err := c.Session.CombinedOutput(command)
	if err != nil {
		return "", fmt.Errorf("command execution failed: %v", err)
	}

	return string(output), nil
}

func (c *Connection) StartShell() error {
	if c.Session == nil {
		return fmt.Errorf("no active session")
	}

	c.Session.Stdout = os.Stdout
	c.Session.Stderr = os.Stderr
	c.Session.Stdin = os.Stdin

	if err := c.Session.Shell(); err != nil {
		return fmt.Errorf("failed to start shell: %v", err)
	}

	return c.Session.Wait()
}

func (c *Connection) Close() error {
	if c.Session != nil {
		c.Session.Close()
	}
	if c.Client != nil {
		return c.Client.Close()
	}
	return nil
}
