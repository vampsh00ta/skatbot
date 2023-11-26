package response

import (
	"fmt"
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/repository/models"
	"strconv"
)

func Variants(variants []models.Variant) *tgmodels.InlineKeyboardMarkup {
	kb := &tgmodels.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgmodels.InlineKeyboardButton{
			{
				{Text: "Описание", CallbackData: "pass"},
				{Text: "Вариант", CallbackData: "pass"},
				{Text: "Скачать файлы", CallbackData: "variant_all_" + strconv.Itoa(variants[0].SubjectId)},
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
				Text: "⬇️", CallbackData: "variant_" + strconv.Itoa(variant.Id),
			},
		}
		kb.InlineKeyboard = append(kb.InlineKeyboard, res)

	}

	return kb
}
func VariantsWithDelete(variants []models.Variant, id int64) *tgmodels.InlineKeyboardMarkup {
	kb := &tgmodels.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgmodels.InlineKeyboardButton{
			{
				{Text: "Описание", CallbackData: "pass"},
				{Text: "Вариант", CallbackData: "pass"},
				{Text: "Скачать файлы", CallbackData: "variant_all_" + strconv.Itoa(variants[0].SubjectId)},
				{Text: "Удалить файл", CallbackData: "pass" + strconv.Itoa(variants[0].SubjectId)},
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
		fmt.Println(variant.TgId)

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

	return kb
}
