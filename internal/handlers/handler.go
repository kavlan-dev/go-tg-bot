package handlers

import (
	"context"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"go.uber.org/zap"
)

type servicesInterface interface {
	DogRandom(ctx context.Context) (string, error)
}

type handler struct {
	services servicesInterface
	bot      *telego.Bot
	log      *zap.SugaredLogger
}

func NewHandler(services servicesInterface, bot *telego.Bot, log *zap.SugaredLogger) *handler {
	return &handler{
		services: services,
		bot:      bot,
		log:      log,
	}
}

func (h *handler) HelpHandle(ctx *th.Context, update telego.Update) error {
	chatID := tu.ID(update.Message.Chat.ID)
	msg := "Простой телеграм бот\n\n/start - начать работу с ботом\n/dog - отправляет случайную фотографию с собакой\n/help - справка по командам"

	_, err := h.bot.SendMessage(ctx, tu.Message(
		chatID,
		msg,
	))

	return err
}

func (h *handler) DogHandler(ctx *th.Context, update telego.Update) error {
	chatID := tu.ID(update.Message.Chat.ID)

	url, err := h.services.DogRandom(ctx)
	if err != nil {
		h.bot.SendMessage(ctx, tu.Message(
			chatID,
			"Произошла ошибка, попробуйте позже",
		))
		return err
	}

	_, err = h.bot.SendPhoto(ctx, tu.Photo(
		chatID,
		tu.FileFromURL(url),
	))
	return err
}
