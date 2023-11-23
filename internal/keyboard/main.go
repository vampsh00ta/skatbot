package keyboard

import (
	"github.com/go-telegram/bot/models"
)

const (
	GetSkatCommand    = "Получить скат"
	AddSkatCommand    = "Добавить скат"
	GaiCommand        = "ГАИ"
	GaishnikCommand   = "Сотрудник ГИБДД"
	MainСommand       = "Главное меню"
	SpravkiCommand    = "Справки"
	MasterDataCommand = "Master Data"
)

func Main() *models.ReplyKeyboardMarkup {
	kb := &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: GetSkatCommand},
			}, {
				{Text: AddSkatCommand},
			},
		},
	}
	return kb
}
