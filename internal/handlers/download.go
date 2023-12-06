package handlers

import (
	"context"
	tgbotapi "github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/repository/models"
	"strconv"
	"strings"
)

func (h BotHandler) DownloadFile() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		var text string
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})

		buttonData := update.CallbackQuery.Data
		idStr := strings.Split(buttonData, "_")[1]
		if idStr == "all" {
			subjectId := strings.Split(buttonData, "_")[2]
			id, err := strconv.Atoi(subjectId)
			if err != nil {
				h.log.Error().Str("Download", "error").Msgf(update.CallbackQuery.Sender.Username, err)
			}
			variants, err := h.service.GetVariantsBySubjectId(ctx, id)
			for _, variant := range variants {
				variant := variant
				SendFile(ctx, b, update, variant)
				h.log.Info().Str("Download", "ok").Str("variantId",
					strconv.Itoa(variant.Id)).Msg(update.CallbackQuery.Sender.Username)

			}

		} else {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				h.log.Error().Str("Download", "error").Msgf(update.CallbackQuery.Sender.Username, err)
			}
			variant, err := h.service.GetVariantbyId(ctx, id)
			if variant == (models.Variant{}) {
				b.SendMessage(ctx, &tgbotapi.SendMessageParams{
					ChatID: update.CallbackQuery.Message.Chat.ID,
					Text:   "Файл был удален",
				})
			}
			SendFile(ctx, b, update, variant)
			h.log.Info().Str("Download", "ok").Str("variantId",
				strconv.Itoa(variant.Id)).Msg(update.CallbackQuery.Sender.Username)

		}
		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: update.CallbackQuery.Message.Chat.ID,
			Text:   text,
		})

	}
}
func SendFile(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update, variant models.Variant) {
	if variant.FileType == "photo" {
		params := &tgbotapi.SendPhotoParams{
			ChatID:  update.CallbackQuery.Message.Chat.ID,
			Photo:   &tgmodels.InputFileString{Data: variant.FileId},
			Caption: "Твой файл",
		}
		b.SendPhoto(ctx, params)
	} else if variant.FileType == "document" {
		params := &tgbotapi.SendDocumentParams{
			ChatID:   update.CallbackQuery.Message.Chat.ID,
			Document: &tgmodels.InputFileString{Data: variant.FileId},
			Caption:  "Твой файл",
		}
		b.SendDocument(ctx, params)

	}

}
