package handlers

import (
	"context"
	tgbotapi "github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/keyboard"
	"skat_bot/internal/service"
	"skat_bot/internal/service/auth"
	log "skat_bot/pkg/logger"
)

type BotHandler struct {
	service service.Service
	log     *log.Logger
	auth    auth.Auth
	back    BackSession
}
type GroupHandler struct {
}

func New(bot *tgbotapi.Bot, s service.Service, log *log.Logger, auth auth.Auth) {
	back := BackSession{user: make(map[int64]Back)}
	botHandler := &BotHandler{s, log, auth, back}
	//get skat
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.GetSkatCommand, tgbotapi.MatchTypeExact, botHandler.GetSkat())
	//add skat
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.AddSkatCommand, tgbotapi.MatchTypeExact, botHandler.AddSkat())
	//break
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		"/"+"break", tgbotapi.MatchTypeExact, BreakSkat())
	//start
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		"/start", tgbotapi.MatchTypeExact, Start())
	//back
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.BackCommand, tgbotapi.MatchTypeExact, back.undo())

}
func Start() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Главное меню",
			ReplyMarkup: keyboard.Main(),
		})
	}
}
func BreakSkat() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		b.UnregisterStepHandler(ctx, update)

		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Главное меню",
			ReplyMarkup: keyboard.Main(),
		})

	}

}
