package handlers

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-co-op/gocron"
	tgbotapi "github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/repository/models"
	"strconv"
	"strings"
)

var scheduler *gocron.Scheduler

func (h BotHandler) doJob(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update, variant models.Variant, job gocron.Job) {

	h.getFile(ctx, b, update, variant)

}
func (h BotHandler) getFile(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update, variant models.Variant) {
	fileName, bs, err := h.service.DownloadVariant(ctx, variant)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fileName)
	params := &tgbotapi.SendDocumentParams{
		ChatID:   update.CallbackQuery.Message.Chat.ID,
		Document: &tgmodels.InputFileUpload{Filename: fileName, Data: bytes.NewReader(*bs)},
		Caption:  "Твой файл",
	}

	b.SendDocument(ctx, params)
}
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
			text = "Подожди,файлы грузятся ⌛"
			subjectId := strings.Split(buttonData, "_")[2]
			id, err := strconv.Atoi(subjectId)
			if err != nil {
				return
			}
			variants, err := h.service.GetVariantsBySubjectId(ctx, id)
			for _, variant := range variants {
				variant := variant

				go h.getFile(ctx, b, update, variant)
			}

		} else {
			text = "Подожди,файл грузится ⌛"
			id, err := strconv.Atoi(idStr)
			if err != nil {
				return
			}
			variant, err := h.service.GetVariantbyId(ctx, id)

			go h.getFile(ctx, b, update, variant)
		}
		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: update.CallbackQuery.Message.Chat.ID,
			Text:   text,
		})

	}
}
