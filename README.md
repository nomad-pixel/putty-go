# Termius Go

Графическое приложение SSH клиента, написанное на Go с использованием библиотеки Fyne.

## Возможности

- Кроссплатформенный GUI (Windows, macOS, Linux)
- Подключение к SSH серверам
- Сохранение настроек
- Современный интерфейс

## Требования

- Go 1.21 или выше
- CGO (для компиляции Fyne)

## Установка зависимостей

```bash
go mod tidy
```

## Запуск приложения

```bash
go run main.go
```

## Сборка

### Для текущей платформы:
```bash
go build -o putty-go
```

### Для других платформ:
```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o putty-go.exe

# Linux
GOOS=linux GOARCH=amd64 go build -o putty-go

# macOS
GOOS=darwin GOARCH=amd64 go build -o putty-go
```

## Структура проекта

```
putty-go/
├── main.go                 # Точка входа приложения
├── go.mod                  # Модуль Go
├── internal/
│   ├── app/
│   │   └── window.go       # Главное окно приложения
│   └── config/
│       └── config.go       # Конфигурация приложения
└── README.md               # Документация
```

## Разработка

Приложение использует архитектуру с разделением на пакеты:
- `internal/app` - логика приложения и GUI
- `internal/config` - управление конфигурацией

## Лицензия

MIT
