package handlers

import (
	"context"
	tgbotapi "github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
)

const (
	BackCommand = "Предыдущий шаг"
)
const (
	MainStep = iota
	InstituteStep
	SubjectNameStep
	SemesterStep
	SubjectTypeStep
	VariantTypeStep
)

func (h BotHandler) Back() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})

		userId := update.CallbackQuery.Sender.ID
		_, err := b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.ID,
			ReplyMarkup: h.fsm.GetKeyboard(userId),
		})
		if err != nil {
			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
				ChatID: update.CallbackQuery.Message.Chat.ID,
				Text:   "Что-то пошло не так",
			})
		}

	}
}
