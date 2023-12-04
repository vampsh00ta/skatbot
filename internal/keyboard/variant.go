package keyboard

import (
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/repository/models"
)

func VariantsTypesTest(variants []models.Variant, page int, command, paddingCommand string) *tgmodels.InlineKeyboardMarkup {

	kb := &tgmodels.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgmodels.InlineKeyboardButton{
			{
				{Text: "Выбери тип варинта ✍", CallbackData: "pass"},
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
	for i := (page - 1) * pageAmount; i < (page-1)*pageAmount+minNum(pageAmount, len(variants)-(page-1)*pageAmount); i++ {
		if i == len(variants) {
			break
		}
		variant := variants[i]

		res := []tgmodels.InlineKeyboardButton{

			{
				Text: variant.TypeName, CallbackData: command + variant.TypeName,
			},
		}
		kb.InlineKeyboard = append(kb.InlineKeyboard, res)
	}
	addPadding(kb, page, len(variants), paddingCommand)
	addBack(kb)

	return kb
}
