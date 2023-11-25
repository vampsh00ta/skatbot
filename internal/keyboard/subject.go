package keyboard

import (
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/repository/models"
)

func SubjectNames(subjects []models.Subject) *tgmodels.ReplyKeyboardMarkup {

	kb := &tgmodels.ReplyKeyboardMarkup{
		Keyboard: [][]tgmodels.KeyboardButton{},
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
