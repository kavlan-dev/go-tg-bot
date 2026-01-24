package handlers

import (
	"context"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"go.uber.org/zap"
)

type ServicesInterface interface {
	DogRandom(ctx context.Context) (string, error)
}

type Handler struct {
	services ServicesInterface
	bot      *telego.Bot
	log      *zap.SugaredLogger
}

func New(services ServicesInterface, bot *telego.Bot, log *zap.SugaredLogger) *Handler {
	return &Handler{
		services: services,
		bot:      bot,
		log:      log,
	}
}

func (h *Handler) HelpHandle(ctx *th.Context, update telego.Update) error {
	chatID := tu.ID(update.Message.Chat.ID)
	msg := "Простой телеграм бот\n\n/start - начать работу с ботом\n/dog - отправляет случайную фотографию с собакой\n/help - справка по командам"

	_, err := h.bot.SendMessage(ctx, tu.Message(
		chatID,
		msg,
	))

	return err
}

func (h *Handler) DogHandler(ctx *th.Context, update telego.Update) error {
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
