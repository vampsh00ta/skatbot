package keyboard

import (
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/repository/models"
	"strconv"
)

const (
	BreakCommand = "Закончить ввод"
)

func Strings(subjects []string) *tgmodels.ReplyKeyboardMarkup {
	kb := &tgmodels.ReplyKeyboardMarkup{
		Keyboard: [][]tgmodels.KeyboardButton{},
	}

	for _, subject := range subjects {
		res := []tgmodels.KeyboardButton{
			{
				Text: subject,
			},
		}
		kb.Keyboard = append(kb.Keyboard, res)
	}

	return kb
}

func SubjectNames(subjects []models.Subject) *tgmodels.ReplyKeyboardMarkup {

	kb := &tgmodels.ReplyKeyboardMarkup{
		Keyboard: [][]tgmodels.KeyboardButton{},
	}

	for _, subject := range subjects {
		res := []tgmodels.KeyboardButton{
			{
				Text: subject.Name,
			},
		}
		kb.Keyboard = append(kb.Keyboard, res)
	}

	return kb
}

func Ints(sems []int) *tgmodels.ReplyKeyboardMarkup {

	kb := &tgmodels.ReplyKeyboardMarkup{
		Keyboard: [][]tgmodels.KeyboardButton{},
	}

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

func SubjectTypes(subjects []models.SubjectType) *tgmodels.ReplyKeyboardMarkup {

	kb := &tgmodels.ReplyKeyboardMarkup{
		Keyboard: [][]tgmodels.KeyboardButton{},
	}

	for _, subjectType := range subjects {
		res := []tgmodels.KeyboardButton{
			{
				Text: subjectType.Name,
			},
		}
		kb.Keyboard = append(kb.Keyboard, res)
	}

	return kb
}
func VariantsNums(variants []models.Variant) *tgmodels.ReplyKeyboardMarkup {

	kb := &tgmodels.ReplyKeyboardMarkup{
		Keyboard: [][]tgmodels.KeyboardButton{},
	}

	for _, variant := range variants {
		res := []tgmodels.KeyboardButton{
			{
				Text: strconv.Itoa(variant.Num),
			},
		}
		kb.Keyboard = append(kb.Keyboard, res)
	}

	return kb
}
func VariantsTypes(variants []models.VariantType) *tgmodels.ReplyKeyboardMarkup {

	kb := &tgmodels.ReplyKeyboardMarkup{
		Keyboard: [][]tgmodels.KeyboardButton{},
	}

	for _, variant := range variants {
		res := []tgmodels.KeyboardButton{
			{
				Text: variant.Name,
			},
		}
		kb.Keyboard = append(kb.Keyboard, res)
	}

	return kb
}
func Pass() *tgmodels.ReplyKeyboardMarkup {

	kb := &tgmodels.ReplyKeyboardMarkup{
		Keyboard: [][]tgmodels.KeyboardButton{
			{
				{Text: "Пропуск"},
			},
		},
	}

	return kb
}
func Empty() *tgmodels.ReplyKeyboardMarkup {

	kb := &tgmodels.ReplyKeyboardMarkup{
		Keyboard: [][]tgmodels.KeyboardButton{},
	}

	return kb
}
