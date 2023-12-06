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

func (h BotHandler) InstitutePaginator() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		buttonData := update.CallbackQuery.Data
		pageStr := strings.Split(buttonData, "_")[1]

		page, err := strconv.Atoi(pageStr)
		if err != nil {
			h.log.Error().Str("InstitutePaginator", "error").Msg(err.Error())
			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
		}

		var insts []int
		var paginator string
		var command string
		if len(strings.Split(buttonData, "#unique#")) > 1 {

			insts, err = h.service.GetUniqueInstitutes(ctx, "", 0, "", true)
			paginator = keyboard.PageInstituteUniquePaginatorData
			command = keyboard.GetSkatInstituteCommand

		} else {
			insts, err = h.service.GetAllInstitutes(ctx, true)
			paginator = keyboard.PageInstitutePaginatorData
			command = keyboard.AddSkatInstituteCommand

		}
		if err != nil {
			h.log.Error().Str("InstitutePaginator", "error").Msg(err.Error())
			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
		}
		b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.ID,

			ReplyMarkup: keyboard.InstituteNumsTest(insts, page, command, paginator),
		})
		h.log.Info().Str("Institute", "ok").Msg("Paginator")

	}
}

func (h BotHandler) SubjectNamePaginator() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		buttonData := update.CallbackQuery.Data

		pageStr := strings.Split(buttonData, "_")[1]
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			h.log.Error().Str("SubjectNamePaginator", "error").Msg(err.Error())
			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
		}

		var subjects []models.Subject
		var paginator string
		var command string

		if len(strings.Split(pageStr, "#")) > 1 {
			data := h.fsm.GetData(update.CallbackQuery.Message.Chat.ID)
			subject := data.(models.Subject)
			subjects, err = h.service.GetUniqueSubjects(ctx, subject.InstistuteNum, 0, "", true)
			paginator = keyboard.PageSubjectNameUniquePaginatorData
			command = keyboard.AddSkatSubjectNameCommand

		} else {
			subjects, err = h.service.GetAllSubjectNames(ctx, true)
			paginator = keyboard.PageSubjectNamePaginatorData
			command = keyboard.GetSkatSubjectNameCommand

		}
		if err != nil {
			h.log.Error().Str("SubjectNamePaginator", "error").Msg(err.Error())
			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
		}
		b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.ID,

			ReplyMarkup: keyboard.SubjectNamesTest(subjects, page, command, paginator),
		})
		h.log.Info().Str("SubjectName", "ok").Msg("Paginator")

	}
}

func (h BotHandler) SemesterPaginator() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		buttonData := update.CallbackQuery.Data

		pageStr := strings.Split(buttonData, "_")[1]
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			h.log.Error().Str("SemesterPaginator", "error").Msg(err.Error())
			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
		}

		var sems []int
		var paginator string
		var command string

		if len(strings.Split(pageStr, "#")) > 1 {
			data := h.fsm.GetData(update.CallbackQuery.Message.Chat.ID)
			subject := data.(models.Subject)
			sems, err = h.service.GetUniqueSemesters(ctx, subject.Name, subject.InstistuteNum, "", true)
			paginator = keyboard.PageSemesterUniquePaginatorData
			command = keyboard.GetSkatSemesterCommand

		} else {
			sems, err = h.service.GetAllSemesters(ctx, true)
			paginator = keyboard.PageSemesterPaginatorData
			command = keyboard.AddSkatSemesterCommand

		}
		if err != nil {
			h.log.Error().Str("SemesterPaginator", "error").Msg(err.Error())
			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
		}
		b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.ID,

			ReplyMarkup: keyboard.SemesterNumsTest(sems, page, command, paginator),
		})
		h.log.Info().Str("Semester", "ok").Msg("Paginator")

	}
}

func (h BotHandler) SubjectTypePaginator() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		buttonData := update.CallbackQuery.Data

		pageStr := strings.Split(buttonData, "_")[1]
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			h.log.Error().Str("SubjectTypePaginator", "error").Msg(err.Error())
			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
		}
		var types []models.Subject
		var paginator string
		var command string
		if len(strings.Split(pageStr, "#")) > 1 {
			data := h.fsm.GetData(update.CallbackQuery.Message.Chat.ID)
			subject := data.(models.Subject)
			types, err = h.service.GetUniqueSubjectTypes(ctx, subject.Name, subject.Semester, subject.InstistuteNum, true)
			paginator = keyboard.PageSubjectTypeUniquePaginatorData
			command = keyboard.GetSkatSubjectTypeCommand

		} else {
			types, err = h.service.GetAllSubjectTypes(ctx, true)
			paginator = keyboard.PageSubjectTypePaginatorData
			command = keyboard.AddSkatSubjectTypeCommand

		}
		if err != nil {
			h.log.Error().Str("SubjectTypePaginator", "error").Msg(err.Error())
			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
		}
		b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.ID,

			ReplyMarkup: keyboard.SubjectTypesTest(types, page, command, paginator),
		})
		h.log.Info().Str("SubjectType", "ok").Msg("Paginator")

	}
}

func (h BotHandler) VariantsPaginator() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		buttonData := update.CallbackQuery.Data
		var err error
		splited := strings.Split(buttonData, "_")

		page, err := strconv.Atoi(splited[1])
		if err != nil {
			h.log.Error().Str("VariantsPaginator", "error").Msg(err.Error())
			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
		}
		userId := update.CallbackQuery.Sender.ID
		var paginator string
		var command string
		var variants []models.Variant
		var kb tgmodels.ReplyMarkup
		if len(strings.Split(splited[0], "#")) > 1 {
			variants, err = h.service.GetVariantbyTgid(ctx, strconv.Itoa(int(userId)))
			paginator = keyboard.PageMyVariantsPaginatorData
			command = keyboard.DeleteMySkatVariantCommand
			kb = keyboard.MyVariantsWithDelete(variants, userId,
				page, command, paginator)

		} else {
			data := h.fsm.GetData(update.CallbackQuery.Message.Chat.ID)
			subject := data.(models.Subject)

			variants, err = h.service.GetVariantsBySubject(ctx, subject)
			paginator = keyboard.PageVariantsPaginatorData
			command = keyboard.DeleteSkatVariantCommand
			kb = keyboard.VariantsWithDelete(variants, userId,
				page, command, paginator)

		}
		if err != nil {
			h.log.Error().Str("VariantsPaginator", "error").Msg(err.Error())
			SendError(ctx, b, update.CallbackQuery.Message.Chat.ID)
			return
		}
		b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.ID,

			ReplyMarkup: kb,
		})
		h.log.Info().Str("Variants", "ok").Msg("Paginator")

	}
}
