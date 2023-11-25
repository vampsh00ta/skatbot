package response

import (
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/repository/models"
	"strconv"
)

func Variants(variants []models.Variant) *tgmodels.InlineKeyboardMarkup {
	kb := &tgmodels.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgmodels.InlineKeyboardButton{
			{
				{Text: "Описание", CallbackData: "pass"},
				{Text: "Номер варианта", CallbackData: "pass"},
				{Text: "Скачать pdf", CallbackData: "pass"},
				{Text: "Скачать word", CallbackData: "pass"},
			},
		},
	}
	ifNilNum := func(num *int) string {
		if num == nil {
			return "Пусто"
		}
		return strconv.Itoa(*num)
	}
	ifNilStr := func(str string) string {
		if str == "" {
			return "Пусто"
		}
		return str
	}
	for _, variant := range variants {
		res := []tgmodels.InlineKeyboardButton{

			{
				Text: ifNilStr(variant.Name), CallbackData: "pass",
			},
			{
				Text: ifNilNum(variant.Num), CallbackData: "pass",
			},
			{
				Text: "pdf", CallbackData: "pass",
			},
			{
				Text: "docx", CallbackData: "pass",
			},
		}
		kb.InlineKeyboard = append(kb.InlineKeyboard, res)

	}

	return kb
}
