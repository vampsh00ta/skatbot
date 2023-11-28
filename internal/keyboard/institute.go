package keyboard

import (
	tgmodels "github.com/go-telegram/bot/models"
	"strconv"
)

func SemesterNumsTest(sems []int, page int, command, paddingCommand string) *tgmodels.InlineKeyboardMarkup {
	kb := &tgmodels.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgmodels.InlineKeyboardButton{
			{
				{Text: "Выбери семестр 🧑‍🎓", CallbackData: "pass"},
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
func InstituteNumsTest(insts []int, page int, command, paddingCommand string) *tgmodels.InlineKeyboardMarkup {
	kb := &tgmodels.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgmodels.InlineKeyboardButton{
			{
				{Text: "Выбери институт  🧑‍🎓", CallbackData: "pass"},
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
