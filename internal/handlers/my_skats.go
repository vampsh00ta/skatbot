package handlers

import (
	"context"
	tgbotapi "github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/keyboard"
	"strconv"
)

func (h BotHandler) MySkats() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		_, err := b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		userId := update.CallbackQuery.Sender.ID
		variants, err := h.service.GetVariantbyTgid(ctx, strconv.Itoa(int(userId)))
		if err != nil {
			h.log.Error().Str("MySkats", "error").Msg(err.Error())
			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
			return
		}

		kb := keyboard.MyVariantsWithDelete(variants, userId, 1, keyboard.DeleteMySkatVariantCommand, keyboard.PageMyVariantsPaginatorData)
		h.fsm.SetKeyboard(update.CallbackQuery.Sender.ID, kb)
		b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.ID,
			ReplyMarkup: kb,
		})

	}
}
