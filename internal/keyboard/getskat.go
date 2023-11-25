package keyboard

import (
	tgmodels "github.com/go-telegram/bot/models"
	"strconv"
)

const (
	BreakCommand = "Закончить ввод"
)

func InstituteNums(sems []int) *tgmodels.ReplyKeyboardMarkup {

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
