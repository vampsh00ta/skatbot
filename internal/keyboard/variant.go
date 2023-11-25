package keyboard

import (
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/repository/models"
)

func VariantsTypes(variants []models.Variant) *tgmodels.ReplyKeyboardMarkup {

	kb := &tgmodels.ReplyKeyboardMarkup{
		Keyboard: [][]tgmodels.KeyboardButton{},
	}
	res := []tgmodels.KeyboardButton{
		{
			Text: BackCommand,
		},
	}

	kb.Keyboard = append(kb.Keyboard, res)
	for _, variant := range variants {
		res := []tgmodels.KeyboardButton{
			{
				Text: variant.TypeName,
			},
		}
		kb.Keyboard = append(kb.Keyboard, res)
	}

	return kb
}
