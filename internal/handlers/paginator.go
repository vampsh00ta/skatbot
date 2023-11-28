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

//func (h BotHandler) InstituteUniquePaginator() tgbotapi.HandlerFunc {
//	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
//		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
//			CallbackQueryID: update.CallbackQuery.ID,
//			ShowAlert:       false,
//		})
//		buttonData := update.CallbackQuery.Data
//		pageStr := strings.Split(buttonData, "_")[1]
//		page, err := strconv.Atoi(pageStr)
//		if err != nil {
//			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
//				ChatID: update.CallbackQuery.Message.Chat.ID,
//				Text:   "Что-то пошло не так",
//			})
//		}
//		insts, err := h.service.GetUniqueInstitutes(ctx, "", 0, "", true)
//		if err != nil {
//			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
//				ChatID: update.CallbackQuery.Message.Chat.ID,
//				Text:   "Что-то пошло не так",
//			})
//		}
//		_, err = b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
//			ChatID:    update.CallbackQuery.Message.Chat.ID,
//			MessageID: update.CallbackQuery.Message.ID,
//
//			ReplyMarkup: keyboard.InstituteNumsTest(insts, page),
//		})
//	}
//}

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
			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
				ChatID: update.CallbackQuery.Message.Chat.ID,
				Text:   "Что-то пошло не так",
			})
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
			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
				ChatID: update.CallbackQuery.Message.Chat.ID,
				Text:   "Что-то пошло не так",
			})
		}
		_, err = b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.ID,

			ReplyMarkup: keyboard.InstituteNumsTest(insts, page, command, paginator),
		})
	}
}

//func (h BotHandler) SubjectNameUniquePaginator() tgbotapi.HandlerFunc {
//	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
//		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
//			CallbackQueryID: update.CallbackQuery.ID,
//			ShowAlert:       false,
//		})
//		buttonData := update.CallbackQuery.Data
//
//		data := h.fsm.GetData(update.CallbackQuery.Message.From.ID)
//		subject := data.(models.Subject)
//		pageStr := strings.Split(buttonData, "_")[1]
//		page, err := strconv.Atoi(pageStr)
//		if err != nil {
//			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
//				ChatID: update.CallbackQuery.Message.Chat.ID,
//				Text:   "Что-то пошло не так",
//			})
//		}
//		subjects, err := h.service.GetUniqueSubjects(ctx, subject.Id, 0, "", true)
//		if err != nil {
//			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
//				ChatID: update.CallbackQuery.Message.Chat.ID,
//				Text:   "Что-то пошло не так",
//			})
//		}
//		_, err = b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
//			ChatID:    update.CallbackQuery.Message.Chat.ID,
//			MessageID: update.CallbackQuery.Message.ID,
//
//			ReplyMarkup: keyboard.SubjectNamesTest(subjects, page),
//		})
//	}
//}

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
			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
				ChatID: update.CallbackQuery.Message.Chat.ID,
				Text:   "Что-то пошло не так",
			})
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
			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
				ChatID: update.CallbackQuery.Message.Chat.ID,
				Text:   "Что-то пошло не так",
			})
		}
		_, err = b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.ID,

			ReplyMarkup: keyboard.SubjectNamesTest(subjects, page, command, paginator),
		})
	}
}

//func (h BotHandler) SemesterUniquePaginator() tgbotapi.HandlerFunc {
//	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
//		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
//			CallbackQueryID: update.CallbackQuery.ID,
//			ShowAlert:       false,
//		})
//		buttonData := update.CallbackQuery.Data
//
//		data := h.fsm.GetData(update.CallbackQuery.Message.From.ID)
//		subject := data.(models.Subject)
//		pageStr := strings.Split(buttonData, "_")[1]
//		page, err := strconv.Atoi(pageStr)
//		if err != nil {
//			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
//				ChatID: update.CallbackQuery.Message.Chat.ID,
//				Text:   "Что-то пошло не так",
//			})
//		}
//		sems, err := h.service.GetUniqueSemesters(ctx, subject.Name, subject.InstistuteNum, "", true)
//		if err != nil {
//			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
//				ChatID: update.CallbackQuery.Message.Chat.ID,
//				Text:   "Что-то пошло не так",
//			})
//		}
//		_, err = b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
//			ChatID:    update.CallbackQuery.Message.Chat.ID,
//			MessageID: update.CallbackQuery.Message.ID,
//
//			ReplyMarkup: keyboard.SemesterNumsTest(sems, page),
//		})
//	}
//}

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
			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
				ChatID: update.CallbackQuery.Message.Chat.ID,
				Text:   "Что-то пошло не так",
			})
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
			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
				ChatID: update.CallbackQuery.Message.Chat.ID,
				Text:   "Что-то пошло не так",
			})
		}
		_, err = b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.ID,

			ReplyMarkup: keyboard.SemesterNumsTest(sems, page, command, paginator),
		})
	}
}

//func (h BotHandler) SubjecTypeUniquePaginator() tgbotapi.HandlerFunc {
//	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
//		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
//			CallbackQueryID: update.CallbackQuery.ID,
//			ShowAlert:       false,
//		})
//		buttonData := update.CallbackQuery.Data
//
//		data := h.fsm.GetData(update.CallbackQuery.Message.From.ID)
//		subject := data.(models.Subject)
//
//		pageStr := strings.Split(buttonData, "_")[1]
//		page, err := strconv.Atoi(pageStr)
//		if err != nil {
//			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
//				ChatID: update.CallbackQuery.Message.Chat.ID,
//				Text:   "Что-то пошло не так",
//			})
//		}
//		types, err := h.service.GetUniqueSubjectTypes(ctx, subject.Name, subject.Semester, subject.InstistuteNum, true)
//		if err != nil {
//			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
//				ChatID: update.CallbackQuery.Message.Chat.ID,
//				Text:   "Что-то пошло не так",
//			})
//		}
//		_, err = b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
//			ChatID:    update.CallbackQuery.Message.Chat.ID,
//			MessageID: update.CallbackQuery.Message.ID,
//
//			ReplyMarkup: keyboard.SubjectTypesTest(types, page),
//		})
//	}
//}

func (h BotHandler) SubjecTypePaginator() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		buttonData := update.CallbackQuery.Data

		pageStr := strings.Split(buttonData, "_")[1]
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
				ChatID: update.CallbackQuery.Message.Chat.ID,
				Text:   "Что-то пошло не так",
			})
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
			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
				ChatID: update.CallbackQuery.Message.Chat.ID,
				Text:   "Что-то пошло не так",
			})
		}
		_, err = b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.ID,

			ReplyMarkup: keyboard.SubjectTypesTest(types, page, command, paginator),
		})
	}
}
