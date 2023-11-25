package handlers

import (
	"context"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"skat_bot/internal/keyboard"
)

const (
	BackCommand = "Предыдущий шаг"
)

func back(ctx context.Context, bot *tgbotapi.Bot, update *models.Update, text string, f tgbotapi.HandlerFunc) bool {
	if text != keyboard.BackCommand {
		return false
	}
	f(ctx, bot, update)
	return true

}
