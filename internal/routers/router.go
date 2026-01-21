package routers

import (
	"context"
	"strings"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type HandlersInterface interface {
	HelpHandle(ctx context.Context, message *telego.Message)
	DogHandler(ctx context.Context, message *telego.Message)
}

func StartRouters(ctx context.Context, updates <-chan telego.Update, bot *telego.Bot, handlers HandlersInterface) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		message := update.Message
		if !strings.HasPrefix(message.Text, "/") {
			bot.SendMessage(ctx, tu.Message(
				tu.ID(message.Chat.ID),
				message.Text,
			))
		}

		switch {
		case strings.HasPrefix(message.Text, "/start"):
			handlers.HelpHandle(ctx, message)
		case strings.HasPrefix(message.Text, "/help"):
			handlers.HelpHandle(ctx, message)
		case strings.HasPrefix(message.Text, "/dog"):
			handlers.DogHandler(ctx, message)
		default:
			bot.SendMessage(ctx, tu.Message(
				tu.ID(message.Chat.ID),
				"Unknown command",
			))
		}
	}
}
