package query_handlers

import (
	"context"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"skat_bot/internal/keyboard"
	"skat_bot/internal/step_handlers"
)

type BotHandler struct {
	step *step_handlers.StepHandler
	back *BackSession
}

func New(bot *tgbotapi.Bot, step *step_handlers.StepHandler) {
	back := &BackSession{user: make(map[int64]*Back)}
	botHandler := &BotHandler{step, back}
	NewMain(bot, botHandler)
	NewGaishnik(bot, botHandler)
	NewGai(bot, botHandler)
	NewCheckVehicle(bot, botHandler)
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		"/start", tgbotapi.MatchTypeExact, Start())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.BackCommand, tgbotapi.MatchTypeExact, back.undo())

}
func Start() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Главное меню",
			ReplyMarkup: keyboard.Main(),
		})
	}
}
