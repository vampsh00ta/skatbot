package handlers

import (
	"context"
	tgbotapi "github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/keyboard"
	"skat_bot/internal/repository/models"
	"skat_bot/internal/response"
	"strconv"
)

func (h BotHandler) GetSkat() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {

		insts, err := h.service.GetUniqueInstitutes(ctx, "", 0, "", true)
		if err != nil {
			h.log.Error(err)
			SendError(ctx, b, update)
			b.UnregisterStepHandler(ctx, update)
			return
		}
		if len(insts) == 0 {
			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Пока что тут пусто,но твоя работа может стать первой :)",
				//ReplyMarkup: keyboard.SubjectsTypes(subjectsTypes),
			})
			b.UnregisterStepHandler(ctx, update)
			return
		}
		_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Выбери доступный институт",
			ReplyMarkup: keyboard.InstituteNums(insts),
		})
		if err != nil {
			h.log.Error(err)
			SendError(ctx, b, update)
			b.UnregisterStepHandler(ctx, update)

			return
		}
		var subject models.Subject
		b.RegisterStepHandler(ctx, update, h.getSkatInstitute, subject)

	}

}
func (h BotHandler) getSkatInstitute(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
	var err error

	data := b.GetStepData(ctx, update)
	currSubject := data.(models.Subject)
	inst, err := strconv.Atoi(update.Message.Text)
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		b.UnregisterStepHandler(ctx, update)
		return
	}
	currSubject.InstistuteNum = inst

	subjects, err := h.service.GetUniqueSubjects(ctx, inst, 0, "", true)
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		b.UnregisterStepHandler(ctx, update)
		return
	}
	_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Выбери доступный учебный предмет",
		ReplyMarkup: keyboard.SubjectNames(subjects),
	})
	b.RegisterStepHandler(ctx, update, h.getSkatName, currSubject)

}
func (h BotHandler) getSkatName(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
	var err error
	data := b.GetStepData(ctx, update)
	currSubject := data.(models.Subject)
	currSubject.Name = update.Message.Text

	sems, err := h.service.GetUniqueSemesters(ctx, currSubject.Name, currSubject.InstistuteNum, "", true)
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		b.UnregisterStepHandler(ctx, update)
		return
	}
	_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Выбери доступный семестр",
		ReplyMarkup: keyboard.SemesterNums(sems),
	})
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		b.UnregisterStepHandler(ctx, update)

		return
	}

	b.RegisterStepHandler(ctx, update, h.getSkatSemester, currSubject)
}

func (h BotHandler) getSkatSemester(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
	var err error
	data := b.GetStepData(ctx, update)
	currSubject := data.(models.Subject)
	semester, err := strconv.Atoi(update.Message.Text)
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		b.UnregisterStepHandler(ctx, update)
		return
	}
	currSubject.Semester = semester

	subjectsTypes, err := h.service.GetUniqueSubjectTypes(ctx, currSubject.Name, currSubject.Semester, currSubject.InstistuteNum, true)
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		return
	}

	_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Выбери доступный тип работы",
		ReplyMarkup: keyboard.SubjectTypes(subjectsTypes),
	})
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		return
	}

	b.RegisterStepHandler(ctx, update, h.getSkatType, currSubject)
}

func (h BotHandler) getSkatType(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
	var err error
	data := b.GetStepData(ctx, update)
	currSubject := data.(models.Subject)
	subjectType := update.Message.Text
	currSubject.TypeName = subjectType

	variants, err := h.service.GetVariantsBySubject(ctx, currSubject)
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		return
	}

	_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Доступные файлы",
		ReplyMarkup: response.Variants(variants),
	})
	_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Ты в главном меню",
		ReplyMarkup: keyboard.Main(),
	})
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		return
	}
}
