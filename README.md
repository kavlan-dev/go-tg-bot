# go-tg-bot

[![Go](https://img.shields.io/badge/go-1.25.5-blue.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Простой Telegram бот на Go, который предоставляет базовые функции для работы с Telegram API.

## Описание

go-tg-bot - это Telegram бот, который:
- Отвечает на команды `/start`, `/help`, `/dog`
- Получает случайные фотографии собак через API [random.dog](https://random.dog)
- Использует современные практики Go и чистую архитектуру

## Установка

### Предварительные требования

- Go 1.25 или выше
- Telegram Bot Token (можно получить у [@BotFather](https://t.me/BotFather))

### Шаги установки

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/kavlan-dev/go-tg-bot.git
   cd go-tg-bot
   ```

2. Создайте конфигурационный файл:
   ```bash
   cp config/config.example.yaml config/config.yaml
   ```

3. Отредактируйте `config/config.yaml` и добавьте ваш Telegram Bot Token:
   ```yaml
   env: "dev" # dev или prod
   token: "YOUR_TELEGRAM_BOT_TOKEN"
   ```

4. Установите зависимости:
   ```bash
   go mod download
   ```

5. Запустите бота:
   ```bash
   go run cmd/bot/main.go
   ```

## Использование

### Доступные команды

- `/start` - Начать работу с ботом (показывает справку)
- `/help` - Показать справку по командам
- `/dog` - Отправить случайную фотографию собаки

### Пример взаимодействия

1. Отправьте `/start` или `/help` чтобы увидеть доступные команды
2. Отправьте `/dog` чтобы получить случайную фотографию собаки
3. Бот ответит на любое текстовое сообщение эхо-ответом

## Конфигурация

Файл конфигурации находится в `config/config.yaml`:

```yaml
env: "dev" # окружение (dev или prod)
token: "YOUR_TELEGRAM_BOT_TOKEN" # токен Telegram бота
```

### Окружения

- `dev` - Режим разработки (подробное логирование)
- `prod` - Продакшн режим (структурированное логирование)

## Архитектура

Проект следует чистой архитектуре с разделение на слои:

```
├── cmd/          # Точка входа
├── config/       # Конфигурация
├── internal/
│   ├── config/   # Загрузка конфигурации
│   ├── handlers/ # Обработчики команд
│   ├── routers/  # Маршрутизация сообщений
│   ├── services/ # Бизнес-логика
│   └── utils/    # Утилиты (логгирование)
```

## Зависимости

Основные зависимости:
- [telego](https://github.com/mymmrac/telego) - Telegram Bot API wrapper
- [viper](https://github.com/spf13/viper) - Управление конфигурацией
- [zap](https://github.com/uber-go/zap) - Логгирование

## Разработка

### Структура проекта

```bash
go-tg-bot/
├── cmd/
│   └── bot/
│       └── main.go          # Основная точка входа
├── config/
│   └── config.example.yaml  # Пример конфигурации
├── internal/
│   ├── config/              # Конфигурация
│   ├── handlers/            # Обработчики команд
│   ├── models/              # Модели данных
│   ├── routers/             # Маршрутизация сообщений
│   ├── services/            # Сервисный слой
│   └── utils/               # Утилиты
├── go.mod                   # Зависимости
├── go.sum                   # Контрольные суммы зависимостей
├── LICENSE                  # Лицензия (MIT)
└── README.md                # Документация
```

### Сборка

```bash
go build -o go-tg-bot cmd/bot/main.go
```

### Запуск

```bash
./go-tg-bot
```

## Лицензия

Проект распространяется под лицензией MIT. См. [LICENSE](LICENSE) для деталей.
