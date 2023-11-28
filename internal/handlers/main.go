package handlers

import (
	"context"
	tgbotapi "github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/keyboard"
)

func (h BotHandler) Menu() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		var id int64
		var chatId int64

		if update.Message != nil {
			id = update.Message.From.ID
			chatId = update.Message.Chat.ID
		} else if update.CallbackQuery != nil {
			id = update.CallbackQuery.Sender.ID
			chatId = update.CallbackQuery.Message.Chat.ID

		}
		b.UnregisterStepHandler(id)
		kb := keyboard.MainBeta()
		h.fsm.ClearKeyboard(id)
		h.fsm.DeleteData(id)
		h.fsm.InitKeyboard(id, kb)

		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      chatId,
			Text:        "Главное меню",
			ReplyMarkup: kb,
		})
	}
}
