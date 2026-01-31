package app

import (
	"context"
	"go-tg-bot/internal/config"
	"go-tg-bot/internal/handler"
	"go-tg-bot/internal/service"
	"go-tg-bot/internal/util"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func Run() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalln(err)
	}
	log, err := util.InitLogger(cfg.Environment)
	if err != nil {
		log.Fatalf("Ошибка инициализации логгера: %v", err)
	}

	bot, err := telego.NewBot(cfg.Token, telego.WithDefaultDebugLogger())
	if err != nil {
		log.Fatalf("Ошибка создания бота: %v", err)
	}

	service := service.NewService()
	handler := handler.NewHandler(service, bot, log)

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
	hd.Handle(handler.HelpHandler, th.CommandEqual("start"))
	hd.Handle(handler.HelpHandler, th.CommandEqual("help"))
	hd.Handle(handler.DogHandler, th.CommandEqual("dog"))

	hd.Start()

	log.Info("Бот остановлен")
}
