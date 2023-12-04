package handlers

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/keyboard"
	"strconv"
)

func (h BotHandler) DeleteSkat() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		_, err := b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		userId := update.CallbackQuery.Sender.ID

		//buttonData := update.CallbackQuery.Data
		//idStr := strings.Split(buttonData, "_")git [1]
		variants, err := h.service.GetVariantbyTgid(ctx, strconv.Itoa(int(userId)))
		if err != nil {
			fmt.Println(err)
			return
		}

		kb := keyboard.MyVariantsWithDelete(variants, userId, 1, keyboard.DeleteSkatVariantCommand, keyboard.PageMyVariantsPaginatorData)
		h.fsm.SetKeyboard(update.CallbackQuery.Sender.ID, kb)
		_, err = b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.ID,
			ReplyMarkup: kb,
		})
		if err != nil {
			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
				ChatID: update.CallbackQuery.Message.Chat.ID,
				Text:   "Что-то пошло не так",
			})
		}

	}
}
