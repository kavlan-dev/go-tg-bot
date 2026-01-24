package main

import (
	"context"
	"go-tg-bot/internal/config"
	"go-tg-bot/internal/handlers"
	"go-tg-bot/internal/services"
	"go-tg-bot/internal/utils"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}
	log, err := utils.InitLogger(cfg.Environment)
	if err != nil {
		log.Fatalf("Ошибка инициализации логгера: %v", err)
	}

	bot, err := telego.NewBot(cfg.Token, telego.WithDefaultDebugLogger())
	if err != nil {
		log.Fatalf("Ошибка создания бота: %v", err)
	}

	services := services.New()
	handlers := handlers.New(services, bot, log)

	botUser, err := bot.GetMe(ctx)
	if err != nil {
		log.Fatalf("Ошибка при проверке работоспособности бота: %v", err)
	}
	log.Info("Бот запущен: @%s (%s)", botUser.Username, botUser.FirstName)

	updates, err := bot.UpdatesViaLongPolling(ctx, nil)
	if err != nil {
		log.Fatalf("Ошибка обновления: %v", err)
	}

	hd, err := th.NewBotHandler(bot, updates)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Info("Остановка приложения...")
		cancel()
	}()

	log.Info("Чтение сообщений...")
	hd.Handle(handlers.HelpHandle, th.CommandEqual("start"))
	hd.Handle(handlers.HelpHandle, th.CommandEqual("help"))
	hd.Handle(handlers.DogHandler, th.CommandEqual("dog"))

	hd.Start()

	log.Info("Бот остановлен")
}
