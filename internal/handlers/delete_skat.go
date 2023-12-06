package handlers

import (
	"context"
	tgbotapi "github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/keyboard"
	"skat_bot/internal/repository/models"
	"strconv"
	"strings"
)

func (h BotHandler) DeleteSkat() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		var err error
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		userId := update.CallbackQuery.Sender.ID
		splited := strings.Split(update.CallbackQuery.Data, "_")
		variantUserId, err := strconv.Atoi(splited[1])
		if err != nil {
			h.log.Error().Str("DeleteSkat", "error").Msg(err.Error())
			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
			return

		}
		variantId, err := strconv.Atoi(splited[2])
		if err != nil {
			h.log.Error().Str("DeleteSkat", "error").Msg(err.Error())

			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
			return

		}
		if variantUserId != int(userId) {

			h.log.Error().Str("DeleteSkat", "error").Msg("no such variant to delete")

			b.SendMessage(ctx, &tgbotapi.SendMessageParams{
				ChatID: update.CallbackQuery.Message.Chat.ID,
				Text:   "Ты не можешь удалить чужой вариант",
			})
			return
		}
		if err := h.service.DeleteVariantById(ctx, variantId); err != nil {
			h.log.Error().Str("DeleteSkat", "error").Msg(err.Error())
			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
			return

		}
		var paginator string
		var command string
		var variants []models.Variant

		if len(strings.Split(splited[0], "#")) > 1 {
			variants, err = h.service.GetVariantbyTgid(ctx, strconv.Itoa(int(userId)))
			command = keyboard.DeleteMySkatVariantCommand
			paginator = keyboard.PageMyVariantsPaginatorData

		} else {
			data := h.fsm.GetData(update.CallbackQuery.Message.Chat.ID)
			subject := data.(models.Subject)
			variants, err = h.service.GetVariantsBySubjectId(ctx, subject.Id)
			command = keyboard.DeleteSkatVariantCommand
			paginator = keyboard.PageVariantsPaginatorData

		}
		kb := keyboard.MyVariantsWithDelete(variants, userId, 1, command, paginator)
		_, err = b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.ID,
			ReplyMarkup: kb,
		})
		_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: update.CallbackQuery.Message.Chat.ID,
			Text:   "Вариант удален",
		})
		h.log.Info().Str("DeleteSkat", "ok").Str("variantId", strconv.Itoa(variantId)).
			Msg(update.CallbackQuery.Sender.Username)

	}
}
