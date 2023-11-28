package keyboard

import (
	tgmodels "github.com/go-telegram/bot/models"
)

const (
	//GetSkatCommand = "Получить скат"
	//AddSkatCommand = "Добавить скат"
	BackCommand = "Назад"
)
const (
	AddSkatCommand            = "addSkat_"
	AddSkatSemesterCommand    = "addSkatSemester_"
	AddSkatInstituteCommand   = "addSkatInstitute_"
	AddSkatSubjectNameCommand = "addSkatSubjectName_"
	AddSkatSubjectTypeCommand = "addSkatSubjectType_"
	AddSkatVariantTypeCommand = "addSkatVariantType_"
)
const (
	GetSkatCommand            = "getSkat_"
	GetSkatSemesterCommand    = "getSkatSemester_"
	GetSkatInstituteCommand   = "getSkatInstitute_"
	GetSkatSubjectNameCommand = "getSkatSubjectName_"
	GetSkatSubjectTypeCommand = "getSkatSubjectType_"
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
func MainBeta() *tgmodels.InlineKeyboardMarkup {
	kb := &tgmodels.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgmodels.InlineKeyboardButton{
			{
				{Text: "Получить скат", CallbackData: GetSkatCommand},
			},
			{
				{Text: "Добавить скат", CallbackData: AddSkatCommand},
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
				{Text: "Пропуск"},
			},
		},
	}

	return kb
}
