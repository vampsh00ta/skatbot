package handlers

import (
	"context"
	tgbotapi "github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
)

func (h BotHandler) Pass() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})

	}
}
