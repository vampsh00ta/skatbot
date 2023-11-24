package handlers

import (
	"context"
	tgbotapi "github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/keyboard"
	"skat_bot/internal/repository/models"
	"strconv"
	"time"
)

func (h BotHandler) AddSkat() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		subjects, err := h.service.GetAllSubjectNames(ctx, true)
		if err != nil {
			h.log.Error(err)
			SendError(ctx, b, update)
			b.UnregisterStepHandler(ctx, update)
			return
		}
		_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Выбери учебный предмет",
			ReplyMarkup: keyboard.Strings(subjects),
		})
		if err != nil {
			h.log.Error(err)
			SendError(ctx, b, update)
			b.UnregisterStepHandler(ctx, update)

			return
		}
		var subject models.Subject
		b.RegisterStepHandler(ctx, update, h.addSkatName, subject)

	}

}
func (h BotHandler) addSkatName(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
	var err error
	data := b.GetStepData(ctx, update)
	currSubject := data.(models.Subject)
	currSubject.Name = update.Message.Text

	sems, err := h.service.GetAllSemesters(ctx, true)
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		b.UnregisterStepHandler(ctx, update)
		return
	}
	_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Выбери семестр",
		ReplyMarkup: keyboard.GetSkatSemester(sems),
	})
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		b.UnregisterStepHandler(ctx, update)

		return
	}

	b.RegisterStepHandler(ctx, update, h.addSkatSemester, currSubject)
}
func (h BotHandler) addSkatSemester(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
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

	subjectsTypes, err := h.service.GetAllSubjectTypes(ctx, true)
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		return
	}
	_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Выбери тип работы",
		ReplyMarkup: keyboard.SubjectTypes(subjectsTypes),
	})
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		return
	}

	b.RegisterStepHandler(ctx, update, h.addSkatType, currSubject)
}
func (h BotHandler) addSkatType(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
	var err error
	data := b.GetStepData(ctx, update)
	currSubject := data.(models.Subject)
	typeName := update.Message.Text

	currSubject.TypeName = typeName
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		return
	}
	subject, err := h.service.GetSubject(ctx, currSubject)
	currSubject.Variants = []models.Variant{
		models.Variant{SubjectId: subject.Id},
	}
	variantTypes, err := h.service.GetVariantTypes(ctx)
	_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Введи тип варианта",
		ReplyMarkup: keyboard.VariantsTypes(variantTypes),
	})
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		return
	}

	b.RegisterStepHandler(ctx, update, h.addSkatWorkType, currSubject)
}
func (h BotHandler) addSkatWorkType(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
	var err error
	data := b.GetStepData(ctx, update)
	currSubject := data.(models.Subject)
	variantType := update.Message.Text

	currSubject.Variants[0].TypeName = variantType

	_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Введи вариант",
		ReplyMarkup: keyboard.Empty(),
	})
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		return
	}

	b.RegisterStepHandler(ctx, update, h.addSkatVariant, currSubject)
}
func (h BotHandler) addSkatVariant(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
	var err error
	data := b.GetStepData(ctx, update)
	currSubject := data.(models.Subject)
	variant, err := strconv.Atoi(update.Message.Text)
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		b.UnregisterStepHandler(ctx, update)
		return
	}
	currSubject.Variants[0].Num = variant
	_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Введи оценку или нажми пропуск",
		ReplyMarkup: keyboard.Pass(),
	})
	b.RegisterStepHandler(ctx, update, h.addSkatGrade, currSubject)

}
func (h BotHandler) addSkatGrade(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
	var err error
	data := b.GetStepData(ctx, update)
	currSubject := data.(models.Subject)
	if update.Message.Text != "Пропуск" {
		grade, err := strconv.Atoi(update.Message.Text)
		if err != nil {
			h.log.Error(err)
			SendError(ctx, b, update)
			b.UnregisterStepHandler(ctx, update)
			return
		}
		currSubject.Variants[0].Grade = &grade
	}

	currSubject.Variants[0].CreationTime = time.Now()

	if err := h.service.AddVariant(ctx, currSubject.Variants[0]); err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		return
	}
	_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Файл добавлен",
	})
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		return
	}
	h.log.Info("AddSkat", "ok")

}
