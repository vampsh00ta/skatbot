package keyboard

import (
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/repository/models"
	"strconv"
)

func MainBeta() *tgmodels.InlineKeyboardMarkup {
	kb := &tgmodels.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgmodels.InlineKeyboardButton{
			{
				{Text: "Получить скат", CallbackData: GetSkatCommand},
			},

			{
				{Text: "Добавить скат", CallbackData: AddSkatCommand},
			},
			{
				{Text: "Мои скаты", CallbackData: GetMySkatsCommand},
			},
		},
	}
	return kb
}

func SemesterNumsTest(sems []int, page int, command, paddingCommand string) *tgmodels.InlineKeyboardMarkup {
	kb := &tgmodels.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgmodels.InlineKeyboardButton{
			{
				{Text: "Выбери семестр 🧑‍🎓", CallbackData: "pass"},
			},
		},
	}

	for i := (page - 1) * pageAmount; i < (page-1)*pageAmount+minNum(pageAmount, len(sems)-(page-1)*pageAmount); i++ {
		if i == len(sems) {
			break
		}
		inst := sems[i]

		res := []tgmodels.InlineKeyboardButton{

			{
				Text: strconv.Itoa(inst), CallbackData: command + strconv.Itoa(inst),
			},
		}
		kb.InlineKeyboard = append(kb.InlineKeyboard, res)
	}
	addPadding(kb, page, len(sems), paddingCommand)
	addBack(kb)

	return kb
}
func SubjectNamesTest(subjects []models.Subject, page int, command, paddingCommand string) *tgmodels.InlineKeyboardMarkup {

	kb := &tgmodels.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgmodels.InlineKeyboardButton{
			{
				{Text: "Выбери учебный предмет 📚", CallbackData: "pass"},
			},
		},
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

func SubjectTypesTest(subjects []models.Subject, page int, command, paddingCommand string) *tgmodels.InlineKeyboardMarkup {

	kb := &tgmodels.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgmodels.InlineKeyboardButton{
			{
				{Text: "Выбери тип работы  ✍", CallbackData: "pass"},
			},
		},
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

func InstituteNumsTest(insts []int, page int, command, paddingCommand string) *tgmodels.InlineKeyboardMarkup {
	kb := &tgmodels.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgmodels.InlineKeyboardButton{
			{
				{Text: "Выбери институт  🧑‍🎓", CallbackData: "pass"},
			},
		},
	}

	for i := (page - 1) * pageAmount; i < (page-1)*pageAmount+minNum(pageAmount, len(insts)-(page-1)*pageAmount); i++ {
		if i == len(insts) {
			break
		}
		inst := insts[i]

		res := []tgmodels.InlineKeyboardButton{

			{
				Text: strconv.Itoa(inst), CallbackData: command + strconv.Itoa(inst),
			},
		}
		kb.InlineKeyboard = append(kb.InlineKeyboard, res)
	}
	addPadding(kb, page, len(insts), paddingCommand)
	addBack(kb)
	return kb
}
