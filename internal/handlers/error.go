package handlers

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
)

func SendError(ctx context.Context, bot *tgbotapi.Bot, chatId int64) {
	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: chatId,
		Text:   fmt.Sprintf("Что-то пошло не так"),
	})

}
