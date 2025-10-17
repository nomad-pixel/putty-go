# Termius Go

A graphical SSH client application written in Go using the Fyne library.

## Features

- Cross-platform GUI (Windows, macOS, Linux)
- SSH server connections
- Settings persistence
- Modern interface

## Requirements

- Go 1.21 or higher
- CGO (for Fyne compilation)

## Installing Dependencies

```bash
go mod tidy
```

## Running the Application

```bash
go run main.go
```

## Building

### For current platform:
```bash
go build -o putty-go
```

### For other platforms:
```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o putty-go.exe

# Linux
GOOS=linux GOARCH=amd64 go build -o putty-go

# macOS
GOOS=darwin GOARCH=amd64 go build -o putty-go
```

## Project Structure

```
putty-go/
├── main.go                 # Application entry point
├── go.mod                  # Go module
├── internal/
│   ├── app/
│   │   └── window.go       # Main application window
│   └── config/
│       └── config.go       # Application configuration
└── README.md               # Documentation
```

## Development

The application uses a package-based architecture:
- `internal/app` - application logic and GUI
- `internal/config` - configuration management

## License

MIT
