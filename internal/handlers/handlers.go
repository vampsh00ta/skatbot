package handlers

import (
	"bytes"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/fsm"
	"skat_bot/internal/keyboard"
	"skat_bot/internal/service"
	log "skat_bot/pkg/logger"
	"strconv"
	"strings"
)

type BotHandler struct {
	service service.Service
	log     *log.Logger
	fsm     fsm.Fsm
}
type GroupHandler struct {
}

func New(bot *tgbotapi.Bot, s service.Service, log *log.Logger) {
	f := fsm.New()
	botHandler := &BotHandler{s, log, f}
	//get skat
	//bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
	//	keyboard.GetSkatCommand, tgbotapi.MatchTypeExact, botHandler.GetSkat())
	////add skat
	//bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
	//	keyboard.AddSkatCommand, tgbotapi.MatchTypeExact, botHandler.AddSkat())
	//back

	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.Back,
		tgbotapi.MatchTypeContains,
		botHandler.Back())
	//add skat beta
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.AddSkatCommand,
		tgbotapi.MatchTypeContains,
		botHandler.AddSkatBeta())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.AddSkatInstituteCommand,
		tgbotapi.MatchTypeContains,
		botHandler.AddSkatBetaInstitute())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.AddSkatSubjectNameCommand,
		tgbotapi.MatchTypeContains,
		botHandler.AddSkatBetaSubjectName())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.AddSkatSemesterCommand,
		tgbotapi.MatchTypeContains,
		botHandler.AddSkatBetaSemester())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.AddSkatSubjectTypeCommand,
		tgbotapi.MatchTypeContains,
		botHandler.AddSkatBetaSubjectType())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.AddSkatVariantTypeCommand,
		tgbotapi.MatchTypeContains,
		botHandler.AddSkatBetaVariantType())

	//get skat beta
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.GetSkatCommand,
		tgbotapi.MatchTypeContains,
		botHandler.GetSkatBeta())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.GetSkatInstituteCommand,
		tgbotapi.MatchTypeContains,
		botHandler.GetSkatBetaInstitute())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.GetSkatSubjectNameCommand,
		tgbotapi.MatchTypeContains,
		botHandler.GetSkatBetaSubjectName())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.GetSkatSemesterCommand,
		tgbotapi.MatchTypeContains,
		botHandler.GetSkatBetaSemester())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.GetSkatSubjectTypeCommand,
		tgbotapi.MatchTypeContains,
		botHandler.GetSkatBetaSubjectType())

	//download skat

	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		"variant", tgbotapi.MatchTypePrefix, botHandler.DownloadFile())
	//paginators
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.PageInstitutePaginatorData, tgbotapi.MatchTypeContains, botHandler.InstitutePaginator())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.PageSemesterPaginatorData, tgbotapi.MatchTypeContains, botHandler.SemesterPaginator())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.PageSubjectNamePaginatorData, tgbotapi.MatchTypeContains, botHandler.SubjectNamePaginator())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.PageSubjectTypePaginatorData, tgbotapi.MatchTypeContains, botHandler.SubjecTypePaginator())
	//pass callback
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		"pass", tgbotapi.MatchTypePrefix, botHandler.Pass())
	//start
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		"/start", tgbotapi.MatchTypeExact, Start())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		"/startbeta", tgbotapi.MatchTypeExact, botHandler.StartBeta())
	////back
	//bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
	//	keyboard.BackCommand, tgbotapi.MatchTypeExact, back.undo())

}
func Start() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		b.UnregisterStepHandler(update.Message.From.ID)
		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Главное меню",
			ReplyMarkup: keyboard.Main(),
		})
	}
}

func (h BotHandler) StartBeta() tgbotapi.HandlerFunc {
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
		h.StartBeta()(ctx, b, update)

	}
}
func (h BotHandler) Pass() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
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
