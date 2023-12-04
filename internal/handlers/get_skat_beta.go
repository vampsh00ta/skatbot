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

func (h BotHandler) GetSkatBeta() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		insts, err := h.service.GetUniqueInstitutes(ctx, "", 0, "", true)
		if err != nil {
			h.log.Error().Str("GetSkatBeta", "error").Msg(err.Error())
			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)

		}
		kb := keyboard.InstituteNumsTest(insts, 1, keyboard.GetSkatInstituteCommand, keyboard.PageInstituteUniquePaginatorData)
		h.fsm.SetKeyboard(update.CallbackQuery.Sender.ID, kb)

		_, err = b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.ID,
			ReplyMarkup: kb,
		})
		h.fsm.AddData(update.CallbackQuery.Message.Chat.ID, models.Subject{Variants: []models.Variant{
			models.Variant{},
		}})
		h.log.Info().Str("GetSkatBeta", "ok").Msg(update.CallbackQuery.Sender.Username)
	}
}

func (h BotHandler) GetSkatBetaInstitute() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		data := h.fsm.GetData(update.CallbackQuery.Message.Chat.ID)
		if data == nil {
			return
		}
		currSubject := data.(models.Subject)

		instituteStr := strings.Split(update.CallbackQuery.Data, "_")[1]
		inst, err := strconv.Atoi(instituteStr)
		if err != nil {
			h.log.Error().Str("GetSkatBetaInstitute", "error").Msg(err.Error())

			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)

		}
		currSubject.InstistuteNum = inst

		subjects, err := h.service.GetUniqueSubjects(ctx, inst, 0, "", true)
		if err != nil {
			h.log.Error().Str("GetSkatBetaInstitute", "error").Msg(err.Error())

			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)

			b.UnregisterStepHandler(update.Message.From.ID)
			return
		}
		kb := keyboard.SubjectNamesTest(subjects, 1, keyboard.GetSkatSubjectNameCommand, keyboard.PageSubjectNameUniquePaginatorData)
		h.fsm.SetKeyboard(update.CallbackQuery.Sender.ID, kb)
		_, err = b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.ID,
			ReplyMarkup: kb,
		})

		if err != nil {
			h.log.Error().Str("GetSkatBetaInstitute", "error").Msg(err.Error())

			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
			b.UnregisterStepHandler(update.CallbackQuery.Message.From.ID)

			return
		}

		h.fsm.AddData(update.CallbackQuery.Message.Chat.ID, currSubject)
		h.log.Info().Str("GetSkatBetaInstitute", "ok").Msg(update.CallbackQuery.Sender.Username)

	}

}
func (h BotHandler) GetSkatBetaSubjectName() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		data := h.fsm.GetData(update.CallbackQuery.Message.Chat.ID)
		if data == nil {
			return
		}
		currSubject := data.(models.Subject)

		subjectName := strings.Split(update.CallbackQuery.Data, "_")[1]
		currSubject.Name = subjectName
		sems, err := h.service.GetUniqueSemesters(ctx, subjectName, currSubject.InstistuteNum, "", true)
		if err != nil {
			h.log.Error().Str("GetSkatBetaSubjectName", "error").Msg(err.Error())
			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
			b.UnregisterStepHandler(update.CallbackQuery.Message.Chat.ID)
			return
		}
		kb := keyboard.SemesterNumsTest(sems, 1, keyboard.GetSkatSemesterCommand, keyboard.PageSemesterUniquePaginatorData)
		h.fsm.SetKeyboard(update.CallbackQuery.Sender.ID, kb)
		_, err = b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.ID,

			ReplyMarkup: kb,
		})

		if err != nil {
			h.log.Error().Str("GetSkatBetaSubjectName", "error").Msg(err.Error())

			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
			b.UnregisterStepHandler(update.Message.From.ID)

			return
		}

		h.fsm.AddData(update.CallbackQuery.Message.Chat.ID, currSubject)
		h.log.Info().Str("GetSkatBetaSubjectName", "ok").Msg(update.CallbackQuery.Sender.Username)

	}
}
func (h BotHandler) GetSkatBetaSemester() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		var err error
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		data := h.fsm.GetData(update.CallbackQuery.Message.Chat.ID)
		if data == nil {
			return
		}
		currSubject := data.(models.Subject)

		semesterStr := strings.Split(update.CallbackQuery.Data, "_")[1]
		sem, err := strconv.Atoi(semesterStr)
		if err != nil {
			h.log.Error().Str("GetSkatBetaSemester", "error").Msg(err.Error())

			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
			b.UnregisterStepHandler(update.CallbackQuery.Message.Chat.ID)
			return
		}
		currSubject.Semester = sem

		subjectsTypes, err := h.service.GetUniqueSubjectTypes(ctx, currSubject.Name, currSubject.Semester, currSubject.InstistuteNum, true)
		if err != nil {
			h.log.Error().Str("GetSkatBetaSemester", "error").Msg(err.Error())

			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
			return
		}
		kb := keyboard.SubjectTypesTest(subjectsTypes, 1, keyboard.GetSkatSubjectTypeCommand, keyboard.PageSubjectTypeUniquePaginatorData)
		h.fsm.SetKeyboard(update.CallbackQuery.Sender.ID, kb)
		_, err = b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.ID,

			ReplyMarkup: kb,
		})

		if err != nil {
			h.log.Error().Str("GetSkatBetaSemester", "error").Msg(err.Error())

			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
			return
		}
		h.fsm.AddData(update.CallbackQuery.Message.Chat.ID, currSubject)
		h.log.Info().Str("GetSkatBetaSemester", "ok").Msg(update.CallbackQuery.Sender.Username)

	}
}

func (h BotHandler) GetSkatBetaSubjectType() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		var err error
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		h.fsm.ClearKeyboard(update.CallbackQuery.Sender.ID)
		data := h.fsm.GetData(update.CallbackQuery.Message.Chat.ID)
		if data == nil {
			return
		}
		currSubject := data.(models.Subject)

		typeName := strings.Split(update.CallbackQuery.Data, "_")[1]
		currSubject.TypeName = typeName
		variants, err := h.service.GetVariantsBySubject(ctx, currSubject)
		if err != nil {
			h.log.Error().Str("GetSkatBetaSubjectType", "error").Msg(err.Error())

			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
			return
		}

		_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.CallbackQuery.Message.Chat.ID,
			Text:        "Доступные файлы",
			ReplyMarkup: keyboard.VariantsWithDelete(variants, update.CallbackQuery.Message.Chat.ID),
		})

		if err != nil {
			h.log.Error().Str("GetSkatBetaSubjectType", "error").Msg(err.Error())

			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
			return
		}
		h.log.Info().Str("GetSkatBetaSubjectType", "ok").Msg(update.CallbackQuery.Sender.Username)

	}
}
