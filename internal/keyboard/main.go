package keyboard

import (
	tgmodels "github.com/go-telegram/bot/models"
)

const (
	GetSkatCommand = "Получить скат"
	AddSkatCommand = "Добавить скат"
	BackCommand    = "Назад"
)

func Main() *tgmodels.ReplyKeyboardMarkup {
	kb := &tgmodels.ReplyKeyboardMarkup{
		Keyboard: [][]tgmodels.KeyboardButton{
			{
				{Text: GetSkatCommand},
			}, {
				{Text: AddSkatCommand},
			},
		},
	}
	return kb
}

func Pass() *tgmodels.ReplyKeyboardMarkup {

	kb := &tgmodels.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard: [][]tgmodels.KeyboardButton{
			{
				{Text: BackCommand},
			},
			{
				{Text: "Пропуск"},
			},
		},
	}

	return kb
}
