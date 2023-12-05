package keyboard

import (
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/repository/models"
	"strconv"
)

func MyVariantsWithDelete(variants []models.Variant, id int64, page int, command, paddingCommand string) *tgmodels.InlineKeyboardMarkup {
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
				{Text: "Скачать файлы", CallbackData: DownloadVariant + "all_" + variants[0].TgId},
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
	deleteDataConst := command + strconv.Itoa(int(id)) + "_"
	for i := (page - 1) * pageAmount; i < (page-1)*pageAmount+minNum(pageAmount, len(variants)-(page-1)*pageAmount); i++ {
		if i == len(variants) {
			break
		}
		variant := variants[i]
		deleteData := deleteDataConst
		emodjiDelete := ""
		if variant.TgId == strconv.Itoa(int(id)) {

			deleteData += strconv.Itoa(variant.Id)

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
				Text: "⬇️", CallbackData: DownloadVariant + strconv.Itoa(variant.Id),
			},
			{Text: emodjiDelete, CallbackData: deleteData},
		}
		kb.InlineKeyboard = append(kb.InlineKeyboard, res)
	}

	addPadding(kb, page, len(variants), paddingCommand)

	addBack(kb)
	return kb
}

func Variants(variants []models.Variant) *tgmodels.InlineKeyboardMarkup {
	kb := &tgmodels.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgmodels.InlineKeyboardButton{
			{
				{Text: "Описание", CallbackData: "pass"},
				{Text: "Вариант", CallbackData: "pass"},
				{Text: "Скачать файлы", CallbackData: DownloadVariant + "all_" + strconv.Itoa(variants[0].SubjectId)},
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
				Text: "⬇️", CallbackData: DownloadVariant + strconv.Itoa(variant.Id),
			},
		}
		kb.InlineKeyboard = append(kb.InlineKeyboard, res)

	}

	return kb
}
func VariantsWithDelete(variants []models.Variant, id int64, page int, command, paddingCommand string) *tgmodels.InlineKeyboardMarkup {
	kb := &tgmodels.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgmodels.InlineKeyboardButton{
			{
				{Text: "Описание", CallbackData: "pass"},
				{Text: "Вариант", CallbackData: "pass"},
				{Text: "Скачать файлы", CallbackData: DownloadVariant + "all_" + strconv.Itoa(variants[0].SubjectId)},
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
	deleteDataConst := command + strconv.Itoa(int(id)) + "_"
	for i := (page - 1) * pageAmount; i < (page-1)*pageAmount+minNum(pageAmount, len(variants)-(page-1)*pageAmount); i++ {
		if i == len(variants) {
			break
		}
		variant := variants[i]

		deleteData := deleteDataConst
		emodjiDelete := ""
		if variant.TgId == strconv.Itoa(int(id)) {
			deleteData += strconv.Itoa(variant.Id)
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
				Text: "⬇️", CallbackData: DownloadVariant + strconv.Itoa(variant.Id),
			},
			{Text: emodjiDelete, CallbackData: deleteData},
		}
		kb.InlineKeyboard = append(kb.InlineKeyboard, res)

	}
	addPadding(kb, page, len(variants), paddingCommand)

	return kb
}
