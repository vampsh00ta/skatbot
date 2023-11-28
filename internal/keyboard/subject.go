package keyboard

import (
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/repository/models"
)

func SubjectNamesTest(subjects []models.Subject, page int, command, paddingCommand string) *tgmodels.InlineKeyboardMarkup {

	kb := &tgmodels.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgmodels.InlineKeyboardButton{
			{
				{Text: "Выбери учебный предмет 📚", CallbackData: "pass"},
			},
		},
	}
	minNum := func(a, b int) int {
		if a < pageAmount {
			return a
		}
		if b < pageAmount {
			return b
		}
		return pageAmount
	}
	for i := (page - 1) * pageAmount; i < (page-1)*pageAmount+minNum(pageAmount, len(subjects)-(page-1)*pageAmount); i++ {
		if i == len(subjects) {
			break
		}
		subject := subjects[i]

		res := []tgmodels.InlineKeyboardButton{

			{
				Text: subject.Name, CallbackData: command + subject.Name,
			},
		}
		kb.InlineKeyboard = append(kb.InlineKeyboard, res)
	}
	addPadding(kb, page, len(subjects), paddingCommand)
	addBack(kb)

	return kb
}
func SubjectNames(subjects []models.Subject) *tgmodels.ReplyKeyboardMarkup {

	kb := &tgmodels.ReplyKeyboardMarkup{
		ResizeKeyboard: false,
		Keyboard:       [][]tgmodels.KeyboardButton{},
	}

	for _, subject := range subjects {
		res := []tgmodels.KeyboardButton{
			{
				Text: subject.Name,
			},
		}
		kb.Keyboard = append(kb.Keyboard, res)
	}

	return kb
}

func SubjectTypesTest(subjects []models.Subject, page int, command, paddingCommand string) *tgmodels.InlineKeyboardMarkup {

	kb := &tgmodels.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgmodels.InlineKeyboardButton{
			{
				{Text: "Выбери тип работы  ✍", CallbackData: "pass"},
			},
		},
	}
	minNum := func(a, b int) int {
		if a < pageAmount {
			return a
		}
		if b < pageAmount {
			return b
		}
		return pageAmount
	}
	for i := (page - 1) * pageAmount; i < (page-1)*pageAmount+minNum(pageAmount, len(subjects)-(page-1)*pageAmount); i++ {
		if i == len(subjects) {
			break
		}
		subject := subjects[i]

		res := []tgmodels.InlineKeyboardButton{

			{
				Text: subject.TypeName, CallbackData: command + subject.TypeName,
			},
		}
		kb.InlineKeyboard = append(kb.InlineKeyboard, res)
	}
	addPadding(kb, page, len(subjects), paddingCommand)
	addBack(kb)

	return kb
}
func SubjectTypes(subjects []models.Subject) *tgmodels.ReplyKeyboardMarkup {
	kb := &tgmodels.ReplyKeyboardMarkup{
		Keyboard: [][]tgmodels.KeyboardButton{},
	}
	res := []tgmodels.KeyboardButton{
		{
			Text: BackCommand,
		},
	}

	kb.Keyboard = append(kb.Keyboard, res)

	for _, subject := range subjects {
		res := []tgmodels.KeyboardButton{
			{
				Text: subject.TypeName,
			},
		}
		kb.Keyboard = append(kb.Keyboard, res)
	}

	return kb
}
