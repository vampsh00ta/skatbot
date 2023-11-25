package handlers

import (
	"bytes"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/keyboard"
	"skat_bot/internal/service"
	"skat_bot/internal/service/auth"
	log "skat_bot/pkg/logger"
	"strconv"
	"strings"
)

type BotHandler struct {
	service service.Service
	log     *log.Logger
	auth    auth.Auth
	back    BackSession
}
type GroupHandler struct {
}

func New(bot *tgbotapi.Bot, s service.Service, log *log.Logger, auth auth.Auth) {
	back := BackSession{user: make(map[int64]Back)}
	botHandler := &BotHandler{s, log, auth, back}
	//get skat
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.GetSkatCommand, tgbotapi.MatchTypeExact, botHandler.GetSkat())
	//add skat
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.AddSkatCommand, tgbotapi.MatchTypeExact, botHandler.AddSkat())
	//download skat
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		"variant", tgbotapi.MatchTypePrefix, botHandler.DownloadFile())
	//start
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		"/start", tgbotapi.MatchTypeExact, Start())
	//back
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.BackCommand, tgbotapi.MatchTypeExact, back.undo())

}
func Start() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		b.UnregisterStepHandler(ctx, update)
		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Главное меню",
			ReplyMarkup: keyboard.Main(),
		})
	}
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
				go func() {

					fileName, bs, err := h.service.DownloadVariant(ctx, variant)
					if err != nil {
						fmt.Println(err)
						return
					}
					params := &tgbotapi.SendDocumentParams{
						ChatID:   update.CallbackQuery.Message.Chat.ID,
						Document: &tgmodels.InputFileUpload{Filename: fileName, Data: bytes.NewReader(*bs)},
						Caption:  "Твой файл",
					}

					b.SendDocument(ctx, params)
				}()
			}

		} else {
			text = "Подожди,файл грузится ⌛"
			id, err := strconv.Atoi(idStr)
			if err != nil {
				return
			}
			go func() {

				fileName, bs, err := h.service.DownloadVariantById(ctx, id)
				if err != nil {
					fmt.Println(err)
					return
				}
				params := &tgbotapi.SendDocumentParams{
					ChatID:   update.CallbackQuery.Message.Chat.ID,
					Document: &tgmodels.InputFileUpload{Filename: fileName, Data: bytes.NewReader(*bs)},
					Caption:  "Твой файл",
				}

				b.SendDocument(ctx, params)
			}()

		}
		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: update.CallbackQuery.Message.Chat.ID,
			Text:   text,
		})

	}
}

//func ReadFiles() tgbotapi.HandlerFunc {
//	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
//
//
//		b.UnregisterStepHandler(ctx, update)
//
//		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
//			ChatID:      update.Message.Chat.ID,
//			Text:        "Главное меню",
//			ReplyMarkup: keyboard.Main(),
//		})
//
//	}
//
//}
