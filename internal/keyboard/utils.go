package keyboard

import (
	tgmodels "github.com/go-telegram/bot/models"
	"math"
	"strconv"
)

const (
	Back = "back"
)

func addPadding(kb *tgmodels.InlineKeyboardMarkup, page, objLen int, paddingCommand string) {
	d := float64(objLen) / float64(pageAmount)
	allPages := int(math.Ceil(d))
	var res []tgmodels.InlineKeyboardButton
	if allPages == 1 {
		res = []tgmodels.InlineKeyboardButton{
			{
				Text: strconv.Itoa(page) + "/" + strconv.Itoa(allPages), CallbackData: "pass",
			},
		}

	} else if page == 1 {
		res = []tgmodels.InlineKeyboardButton{
			{
				Text: strconv.Itoa(page) + "/" + strconv.Itoa(allPages), CallbackData: "pass",
			},
			{
				Text: "Вперед", CallbackData: paddingCommand + strconv.Itoa(page+1),
			},
		}

	} else if page == allPages {
		res = []tgmodels.InlineKeyboardButton{
			{
				Text: "Назад", CallbackData: paddingCommand + strconv.Itoa(page-1),
			},
			{
				Text: strconv.Itoa(page) + "/" + strconv.Itoa(allPages), CallbackData: "pass",
			},
		}
	} else if page > 1 && page != allPages {
		res = []tgmodels.InlineKeyboardButton{
			{
				Text: "Назад", CallbackData: paddingCommand + strconv.Itoa(page-1),
			},
			{
				Text: strconv.Itoa(page) + "/" + strconv.Itoa(allPages), CallbackData: "pass",
			},
			{
				Text: "Вперед", CallbackData: paddingCommand + strconv.Itoa(page+1),
			},
		}
	}
	(*kb).InlineKeyboard = append((*kb).InlineKeyboard, res)
}
func addBack(kb *tgmodels.InlineKeyboardMarkup) {

	res := []tgmodels.InlineKeyboardButton{
		{
			Text: "Прошлый шаг", CallbackData: Back,
		},
	}
	(*kb).InlineKeyboard = append((*kb).InlineKeyboard, res)
}
func minNum(a, b int) int {
	if a < pageAmount {
		return a
	}
	if b < pageAmount {
		return b
	}
	return pageAmount
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
