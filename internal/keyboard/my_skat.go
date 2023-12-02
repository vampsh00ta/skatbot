package keyboard

import (
	"fmt"
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/repository/models"
	"strconv"
)

func MyVariantsWithDelete(variants []models.Variant, id int64) *tgmodels.InlineKeyboardMarkup {
	fmt.Println(variants, variants == nil)
	if len(variants) == 0 {
		kb := &tgmodels.InlineKeyboardMarkup{
			InlineKeyboard: [][]tgmodels.InlineKeyboardButton{
				{
					{Text: "Пусто", CallbackData: "pass"},
				},
			},
		}
		addBack(kb)
		return kb
	}
	kb := &tgmodels.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgmodels.InlineKeyboardButton{
			{
				{Text: "Описание", CallbackData: "pass"},
				{Text: "Вариант", CallbackData: "pass"},
				{Text: "Скачать файлы", CallbackData: "variant_all_" + variants[0].TgId},
				{Text: "Удалить файлы", CallbackData: "deleteAllVariants_" + variants[0].TgId},
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
	deleteDataConst := "deleteVariant_" + strconv.Itoa(int(id)) + "_"
	for _, variant := range variants {

		deleteData := deleteDataConst
		emodjiDelete := ""
		if variant.TgId == strconv.Itoa(int(id)) {
			deleteData += strconv.Itoa(int(id))
			emodjiDelete = "✅"
		} else {
			deleteData = "pass"
			emodjiDelete = "❌"
		}

		res := []tgmodels.InlineKeyboardButton{

			{
				Text: ifNilStr(variant.Name), CallbackData: "pass",
			},
			{
				Text: ifNilNum(variant.Num), CallbackData: "pass",
			},
			{
				Text: "⬇️", CallbackData: "variant_" + strconv.Itoa(variant.Id),
			},
			{Text: emodjiDelete, CallbackData: deleteData},
		}
		kb.InlineKeyboard = append(kb.InlineKeyboard, res)

	}
	addBack(kb)

	return kb
}
