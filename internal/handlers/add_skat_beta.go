package handlers

import (
	"context"
	tgbotapi "github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/keyboard"
	"skat_bot/internal/repository/models"
	"strconv"
	"strings"
	"time"
)

func (h BotHandler) AddSkatBeta() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		insts, err := h.service.GetAllInstitutes(ctx, true)
		if err != nil {
			h.log.Error().Str("AddSkatBeta", "error").Msg(err.Error())

			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
				ChatID: update.CallbackQuery.Message.Chat.ID,
				Text:   "Что-то пошло не так",
			})
		}
		kb := keyboard.InstituteNumsTest(insts, 1, keyboard.AddSkatInstituteCommand, keyboard.PageInstitutePaginatorData)
		h.fsm.SetKeyboard(update.CallbackQuery.Sender.ID, kb)
		_, err = b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.ID,
			ReplyMarkup: kb,
		})

		h.fsm.AddData(update.CallbackQuery.Sender.ID, models.Subject{Variants: []models.Variant{
			models.Variant{},
		}})
		h.log.Info().Str("AddSkatBeta", "ok").Msg(update.CallbackQuery.Sender.Username)

	}
}

func (h BotHandler) AddSkatBetaInstitute() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		data := h.fsm.GetData(update.CallbackQuery.Sender.ID)
		if data == nil {
			return
		}
		currSubject := data.(models.Subject)

		instituteStr := strings.Split(update.CallbackQuery.Data, "_")[1]
		inst, err := strconv.Atoi(instituteStr)
		if err != nil {
			h.log.Error().Str("AddSkatBetaInstitute", "error").Msg(err.Error())

			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
				ChatID: update.CallbackQuery.Message.Chat.ID,
				Text:   "Что-то пошло не так",
			})
		}
		currSubject.InstistuteNum = inst

		subjects, err := h.service.GetAllSubjectNames(ctx, true)
		if err != nil {
			h.log.Error().Str("AddSkatBetaInstitute", "error").Msg(err.Error())
			SendError(ctx, b, update)
			b.UnregisterStepHandler(update.Message.From.ID)
			return
		}
		kb := keyboard.SubjectNamesTest(subjects, 1, keyboard.AddSkatSubjectNameCommand, keyboard.PageSubjectNamePaginatorData)
		h.fsm.SetKeyboard(update.CallbackQuery.Sender.ID, kb)
		_, err = b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.ID,
			ReplyMarkup: keyboard.SubjectNamesTest(subjects, 1, keyboard.AddSkatSubjectNameCommand, keyboard.PageSubjectNamePaginatorData),
		})

		if err != nil {
			h.log.Error().Str("AddSkatBetaInstitute", "error").Msg(err.Error())
			SendError(ctx, b, update)
			b.UnregisterStepHandler(update.CallbackQuery.Message.From.ID)

			return
		}

		h.fsm.AddData(update.CallbackQuery.Sender.ID, currSubject)
		h.log.Info().Str("AddSkatBetaInstitute", "ok").Msg(update.CallbackQuery.Sender.Username)

	}

}
func (h BotHandler) AddSkatBetaSubjectName() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		data := h.fsm.GetData(update.CallbackQuery.Sender.ID)
		if data == nil {
			return
		}
		currSubject := data.(models.Subject)
		subjectName := strings.Split(update.CallbackQuery.Data, "_")[1]
		currSubject.Name = subjectName
		sems, err := h.service.GetAllSemesters(ctx, true)
		if err != nil {
			h.log.Error().Str("AddSkatBetaSubjectName", "error").Msg(err.Error())
			SendError(ctx, b, update)
			b.UnregisterStepHandler(update.Message.From.ID)
			return
		}
		kb := keyboard.SemesterNumsTest(sems, 1, keyboard.AddSkatSemesterCommand, keyboard.PageSemesterPaginatorData)
		h.fsm.SetKeyboard(update.CallbackQuery.Sender.ID, kb)
		_, err = b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.ID,

			ReplyMarkup: keyboard.SemesterNumsTest(sems, 1, keyboard.AddSkatSemesterCommand, keyboard.PageSemesterPaginatorData),
		})

		if err != nil {
			h.log.Error().Str("AddSkatBetaSubjectName", "error").Msg(err.Error())
			SendError(ctx, b, update)
			b.UnregisterStepHandler(update.Message.From.ID)

			return
		}

		h.fsm.AddData(update.CallbackQuery.Sender.ID, currSubject)
		h.log.Info().Str("AddSkatBetaSubjectName", "ok").Msg(update.CallbackQuery.Sender.Username)

	}
}
func (h BotHandler) AddSkatBetaSemester() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		var err error
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		data := h.fsm.GetData(update.CallbackQuery.Sender.ID)
		if data == nil {
			return
		}
		currSubject := data.(models.Subject)

		semesterStr := strings.Split(update.CallbackQuery.Data, "_")[1]
		sem, err := strconv.Atoi(semesterStr)
		if err != nil {
			h.log.Error().Str("AddSkatBetaSemester", "error").Msg(err.Error())
			SendError(ctx, b, update)
			b.UnregisterStepHandler(update.Message.From.ID)
			return
		}
		currSubject.Semester = sem

		subjectsTypes, err := h.service.GetAllSubjectTypes(ctx, true)
		if err != nil {
			h.log.Error().Str("AddSkatBetaSemester", "error").Msg(err.Error())
			SendError(ctx, b, update)
			return
		}
		kb := keyboard.SubjectTypesTest(subjectsTypes, 1, keyboard.AddSkatSubjectTypeCommand, keyboard.PageInstitutePaginatorData)
		h.fsm.SetKeyboard(update.CallbackQuery.Sender.ID, kb)
		_, err = b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.ID,

			ReplyMarkup: keyboard.SubjectTypesTest(subjectsTypes, 1, keyboard.AddSkatSubjectTypeCommand, keyboard.PageInstitutePaginatorData),
		})

		if err != nil {
			h.log.Error().Str("AddSkatBetaSemester", "error").Msg(err.Error())
			SendError(ctx, b, update)
			return
		}
		h.fsm.AddData(update.CallbackQuery.Sender.ID, currSubject)
		h.log.Info().Str("AddSkatBetaSemester", "ok").Msg(update.CallbackQuery.Sender.Username)

	}
}

func (h BotHandler) AddSkatBetaSubjectType() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		var err error
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		data := h.fsm.GetData(update.CallbackQuery.Sender.ID)
		if data == nil {
			return
		}
		currSubject := data.(models.Subject)

		typeName := strings.Split(update.CallbackQuery.Data, "_")[1]
		currSubject.TypeName = typeName

		variantTypes, err := h.service.GetVariantTypes(ctx)

		kb := keyboard.VariantsTypesTest(variantTypes, 1, keyboard.AddSkatVariantTypeCommand, keyboard.PageVariantTypePaginatorData)
		h.fsm.SetKeyboard(update.CallbackQuery.Sender.ID, kb)
		_, err = b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.ID,

			ReplyMarkup: keyboard.VariantsTypesTest(variantTypes, 1, keyboard.AddSkatVariantTypeCommand, keyboard.PageVariantTypePaginatorData),
		})
		if err != nil {
			h.log.Error().Str("AddSkatBetaSubjectType", "error").Msg(err.Error())
			SendError(ctx, b, update)
			return
		}
		h.fsm.AddData(update.CallbackQuery.Sender.ID, currSubject)
		h.log.Info().Str("AddSkatBetaSubjectType", "ok").Msg(update.CallbackQuery.Sender.Username)

	}
}

func (h BotHandler) AddSkatBetaVariantType() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		var err error
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		h.fsm.ClearKeyboard(update.CallbackQuery.Sender.ID)
		data := h.fsm.GetData(update.CallbackQuery.Sender.ID)
		if data == nil {
			return
		}
		currSubject := data.(models.Subject)

		typeName := strings.Split(update.CallbackQuery.Data, "_")[1]
		currSubject.Variants[0].TypeName = typeName

		_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.CallbackQuery.Message.Chat.ID,
			Text:        "Введи вариант или нажми пропуск",
			ReplyMarkup: keyboard.Pass(),
		})

		if err != nil {
			h.log.Error().Str("AddSkatBetaVariantType", "error").Msg(err.Error())
			SendError(ctx, b, update)
			return
		}
		h.fsm.AddData(update.CallbackQuery.Sender.ID, currSubject)
		b.RegisterStepHandler(update.CallbackQuery.Message.Chat.ID, h.AddSkatBetaVariant(), nil)
		h.log.Info().Str("AddSkatBetaVariantType", "ok").Msg(update.CallbackQuery.Sender.Username)

	}
}
func (h BotHandler) AddSkatBetaVariant() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		var err error
		data := h.fsm.GetData(update.Message.From.ID)
		if data == nil {
			return
		}
		currSubject := data.(models.Subject)

		if update.Message.Text != "Пропуск" {
			variant, err := strconv.Atoi(update.Message.Text)

			if err != nil {
				h.log.Error().Str("AddSkatBetaVariant", "error").Msg(err.Error())
				SendError(ctx, b, update)
				b.UnregisterStepHandler(update.Message.From.ID)
				return
			}
			currSubject.Variants[0].Num = &variant
		}
		if err != nil {
			h.log.Error().Str("AddSkatBetaVariant", "error").Msg(err.Error())
			SendError(ctx, b, update)
			b.UnregisterStepHandler(update.Message.From.ID)
			return
		}
		_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Введи описание,чтобы другим было легче найти нужный файл",
			ReplyMarkup: keyboard.Pass(),
		})

		h.fsm.AddData(update.Message.From.ID, currSubject)
		b.RegisterStepHandler(update.Message.From.ID, h.AddSkatBetaDesc(), nil)
		h.log.Info().Str("AddSkatBetaVariant", "ok").Msg(update.Message.From.Username)

	}
}

func (h BotHandler) AddSkatBetaDesc() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		var err error
		data := h.fsm.GetData(update.Message.From.ID)
		if data == nil {
			return
		}
		currSubject := data.(models.Subject)

		text := update.Message.Text
		//back(ctx, b, update, text, h.addSkatWorkType)
		desc := text
		if err != nil {
			h.log.Error().Str("AddSkatBetaDesc", "error").Msg(err.Error())
			SendError(ctx, b, update)
			b.UnregisterStepHandler(update.Message.From.ID)
			return
		}
		currSubject.Variants[0].Name = desc
		_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Введи оценку или нажми пропуск",
			ReplyMarkup: keyboard.Pass(),
		})
		h.fsm.AddData(update.Message.From.ID, currSubject)

		b.RegisterStepHandler(update.Message.From.ID, h.AddSkatBetaGrade(), nil)
		h.log.Info().Str("AddSkatBetaDesc", "ok").Msg(update.Message.From.Username)

	}
}

func (h BotHandler) AddSkatBetaGrade() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {

		var err error
		data := h.fsm.GetData(update.Message.From.ID)
		if data == nil {
			return
		}
		currSubject := data.(models.Subject)

		if update.Message.Text != "Пропуск" {
			grade, err := strconv.Atoi(update.Message.Text)
			if err != nil {
				h.log.Error().Str("AddSkatBetaGrade", "error").Msg(err.Error())
				SendError(ctx, b, update)
				b.UnregisterStepHandler(update.Message.From.ID)
				return
			}
			currSubject.Variants[0].Grade = &grade
		}

		_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Добавь файл(ы)",
		})
		if err != nil {
			h.log.Error().Str("AddSkatBetaGrade", "error").Msg(err.Error())
			SendError(ctx, b, update)
			return
		}
		b.RegisterStepHandler(update.Message.From.ID, h.AddSkatBetaFiles(), nil)
		h.log.Info().Str("AddSkatBetaGrade", "ok").Msg(update.Message.From.Username)

	}

}

func (h BotHandler) AddSkatBetaFiles() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		var err error
		data := h.fsm.GetData(update.Message.From.ID)
		if data == nil {
			return
		}
		currSubject := data.(models.Subject)

		var fileId string
		if update.Message.Document == nil && update.Message.Photo == nil {

			b.SendMessage(ctx, &tgbotapi.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Файл отсуствует, попробуй еще раз",
			})
			return
		}
		if update.Message.Document != nil {
			fileId = update.Message.Document.FileID
			currSubject.Variants[0].FileType = "document"

		} else {
			msgLen := len(update.Message.Photo)
			fileId = update.Message.Photo[msgLen-1].FileID
			currSubject.Variants[0].FileType = "photo"

		}
		currSubject.Variants[0].CreationTime = time.Now()
		currSubject.Variants[0].FileId = fileId
		currSubject.Variants[0].TgId = strconv.Itoa(int(update.Message.From.ID))
		currSubject.Variants[0].TgUsername = update.Message.From.Username
		subject, err := h.service.AddSkat(ctx, currSubject)
		if err != nil {
			h.log.Error().Str("AddSkatBetaFiles", "error").Msg(err.Error())
			SendError(ctx, b, update)
			return
		}

		_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Файл добавлен",
		})
		b.UnregisterStepHandler(update.Message.From.ID)
		h.log.Info().Str("AddSkatBetaFiles", "ok").Str("subjectId", strconv.Itoa(subject.Id)).Msg(update.Message.From.Username)

	}
}
