package handlers

import (
	"context"

	"github.com/mymmrac/telego"
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

func (h *Handler) HelpHandle(ctx context.Context, message *telego.Message) {
	h.bot.SendMessage(ctx, tu.Message(
		tu.ID(message.Chat.ID),
		`
Простой телеграм бот

/start - начать работу с ботом
/dog - отправляет случайную фотографию с собакой
/help - справка по командам
		`,
	))
}

func (h *Handler) DogHandler(ctx context.Context, message *telego.Message) {
	url, err := h.services.DogRandom(ctx)
	if err != nil {
		h.log.Errorln("Не удалось получить ссылку на фотографию:", err)
		h.bot.SendMessage(ctx, tu.Message(
			tu.ID(message.Chat.ID),
			"Произошла ошибка, попробуйте позже",
		))
	}

	h.bot.SendPhoto(ctx, tu.Photo(
		tu.ID(message.Chat.ID),
		tu.FileFromURL(url),
	))
}
