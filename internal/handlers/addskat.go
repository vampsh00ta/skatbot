package handlers

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/keyboard"
	"skat_bot/internal/repository/models"
	"strconv"
	"time"
)

func (h BotHandler) AddSkat() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		insts, err := h.service.GetAllInstitutes(ctx, true)
		if err != nil {
			h.log.Error(err)
			SendError(ctx, b, update)
			b.UnregisterStepHandler(ctx, update)
			return
		}
		_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Выбери институт",
			ReplyMarkup: keyboard.InstituteNums(insts),
		})
		if err != nil {
			h.log.Error(err)
			SendError(ctx, b, update)
			b.UnregisterStepHandler(ctx, update)

			return
		}
		subject := models.Subject{Variants: []models.Variant{
			{},
		}}
		b.RegisterStepHandler(ctx, update, h.addSkatInstitute, subject)

	}

}
func (h BotHandler) addSkatInstitute(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
	var err error
	data := b.GetStepData(ctx, update)
	currSubject := data.(models.Subject)

	institute, err := strconv.Atoi(update.Message.Text)
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		b.UnregisterStepHandler(ctx, update)
		return
	}

	currSubject.InstistuteNum = institute
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		return
	}

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
		ReplyMarkup: keyboard.SubjectNames(subjects),
	})
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		b.UnregisterStepHandler(ctx, update)

		return
	}

	b.RegisterStepHandler(ctx, update, h.addSkatName, currSubject)
}
func (h BotHandler) addSkatName(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
	var err error
	data := b.GetStepData(ctx, update)
	currSubject := data.(models.Subject)
	subjectName := update.Message.Text
	currSubject.Name = subjectName
	ok, err := h.service.CheckSubjectName(ctx, subjectName)
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		b.UnregisterStepHandler(ctx, update)
		return
	}
	if !ok {
		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Такой предмент отсутствует",
		})
	}
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
		ReplyMarkup: keyboard.SemesterNums(sems),
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
	//text := update.Message.Text
	//if back(ctx, b, update, update.Message.Text, h.addSkatName) {
	//	return
	//}
	typeName := update.Message.Text

	currSubject.TypeName = typeName
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		return
	}

	variantTypes, err := h.service.GetVariantTypes(ctx)
	_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Введи тип варианта ",
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
	//if back(ctx, b, update, update.Message.Text, h.addSkatType) {
	//	return
	//}
	variantType := update.Message.Text

	currSubject.Variants[0].TypeName = variantType

	_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Введи вариант или нажми пропуск",
		ReplyMarkup: keyboard.Pass(),
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
	if back(ctx, b, update, update.Message.Text, h.addSkatType) {
		return
	}
	//back(ctx, b, update, text, h.addSkatInstitute)

	if update.Message.Text != "Пропуск" {
		variant, err := strconv.Atoi(update.Message.Text)

		if err != nil {
			h.log.Error(err)
			SendError(ctx, b, update)
			b.UnregisterStepHandler(ctx, update)
			return
		}
		currSubject.Variants[0].Num = &variant
	}
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		b.UnregisterStepHandler(ctx, update)
		return
	}
	_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Введи описание,чтобы другим было легче найти нужный файл",
		ReplyMarkup: keyboard.Pass(),
	})
	b.RegisterStepHandler(ctx, update, h.addSkatDesc, currSubject)

}

func (h BotHandler) addSkatDesc(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
	var err error
	data := b.GetStepData(ctx, update)
	currSubject := data.(models.Subject)
	text := update.Message.Text
	//back(ctx, b, update, text, h.addSkatWorkType)
	desc := text
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		b.UnregisterStepHandler(ctx, update)
		return
	}
	currSubject.Variants[0].Name = desc
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

	_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Добавь файл(ы)",
	})
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		return
	}
	b.RegisterStepHandler(ctx, update, h.addSkatFiles, currSubject)

}
func (h BotHandler) addSkatFiles(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
	var err error
	data := b.GetStepData(ctx, update)
	currSubject := data.(models.Subject)
	//&& update.Message.Photo != nil
	var fileId string
	if update.Message.Document == nil && update.Message.Photo == nil {

		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Файл отсуствует, попробуй еще раз",
			ReplyMarkup: keyboard.Main(),
		})
		return
	}
	if update.Message.Document != nil {
		fileId = update.Message.Document.FileID
	} else {
		fmt.Println(update.Message.Photo)
		msgLen := len(update.Message.Photo)
		fileId = update.Message.Photo[msgLen-1].FileID
	}

	subject, err := h.service.AddOrGetSubject(ctx, currSubject)
	if err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		return
	}
	currSubject.Variants[0].CreationTime = time.Now()
	currSubject.Variants[0].SubjectId = subject.Id
	currSubject.Variants[0].FileId = fileId
	if err := h.service.AddVariant(ctx, currSubject.Variants[0]); err != nil {
		h.log.Error(err)
		SendError(ctx, b, update)
		return
	}
	_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Файл добавлен",
		ReplyMarkup: keyboard.Main(),
	})
	h.log.Info("AddSkat", "ok")
	b.UnregisterStepHandler(ctx, update)

}
