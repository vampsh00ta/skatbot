package keyboard

import (
	"fmt"
	tgmodels "github.com/go-telegram/bot/models"
	"math"
	"strconv"
)

const (
	BreakCommand = "Закончить ввод"
)
const (
	pageAmount = 5
)
const (
	PageInstitutePaginatorData = "institutePage_"
)

func InstituteNumsTest(insts []int, page int) *tgmodels.InlineKeyboardMarkup {
	kb := &tgmodels.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgmodels.InlineKeyboardButton{
			{
				{Text: "Выбери институт", CallbackData: "pass"},
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
	fmt.Println((len(insts) - (page-1)*pageAmount) % pageAmount)
	for i := (page - 1) * pageAmount; i < (page-1)*pageAmount+minNum(pageAmount, len(insts)-(page-1)*pageAmount); i++ {
		if i == len(insts) {
			break
		}
		inst := insts[i]

		res := []tgmodels.InlineKeyboardButton{

			{
				Text: strconv.Itoa(inst), CallbackData: "institute_" + strconv.Itoa(inst),
			},
		}
		kb.InlineKeyboard = append(kb.InlineKeyboard, res)
	}
	allPages := int(math.Ceil(float64((len(insts))) / float64(pageAmount)))
	var res []tgmodels.InlineKeyboardButton

	if page == 1 {
		res = []tgmodels.InlineKeyboardButton{
			{
				Text: strconv.Itoa(page) + "/" + strconv.Itoa(allPages), CallbackData: "pass",
			},
			{
				Text: "Вперед", CallbackData: PageInstitutePaginatorData + strconv.Itoa(page+1),
			},
		}

	} else if page == allPages {
		res = []tgmodels.InlineKeyboardButton{
			{
				Text: "Назад", CallbackData: PageInstitutePaginatorData + strconv.Itoa(page-1),
			},
			{
				Text: strconv.Itoa(page) + "/" + strconv.Itoa(allPages), CallbackData: "pass",
			},
		}

	} else if page > 1 && page != allPages {
		res = []tgmodels.InlineKeyboardButton{
			{
				Text: "Назад", CallbackData: PageInstitutePaginatorData + strconv.Itoa(page-1),
			},
			{
				Text: strconv.Itoa(page) + "/" + strconv.Itoa(allPages), CallbackData: "pass",
			},
			{
				Text: "Вперед", CallbackData: PageInstitutePaginatorData + strconv.Itoa(page+1),
			},
		}
	}
	kb.InlineKeyboard = append(kb.InlineKeyboard, res)

	return kb
}
func InstituteNums(sems []int) *tgmodels.ReplyKeyboardMarkup {

	kb := &tgmodels.ReplyKeyboardMarkup{
		ResizeKeyboard: false,
		Keyboard:       [][]tgmodels.KeyboardButton{},
	}
	res := []tgmodels.KeyboardButton{
		{
			Text: BackCommand,
		},
	}
	kb.Keyboard = append(kb.Keyboard, res)

	for _, sem := range sems {
		res := []tgmodels.KeyboardButton{
			{
				Text: strconv.Itoa(sem),
			},
		}
		kb.Keyboard = append(kb.Keyboard, res)
	}

	return kb
}
func SemesterNums(sems []int) *tgmodels.ReplyKeyboardMarkup {

	kb := &tgmodels.ReplyKeyboardMarkup{
		Keyboard: [][]tgmodels.KeyboardButton{},
	}
	res := []tgmodels.KeyboardButton{
		{
			Text: BackCommand,
		},
	}
	kb.Keyboard = append(kb.Keyboard, res)

	for _, sem := range sems {
		res := []tgmodels.KeyboardButton{
			{
				Text: strconv.Itoa(sem),
			},
		}
		kb.Keyboard = append(kb.Keyboard, res)
	}

	return kb
}

func Empty() *tgmodels.ReplyKeyboardMarkup {

	kb := &tgmodels.ReplyKeyboardMarkup{
		Keyboard: [][]tgmodels.KeyboardButton{},
	}

	return kb
}
