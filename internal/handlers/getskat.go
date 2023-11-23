package query_handlers

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"skat_bot/internal/keyboard"
)

type Gaishnik struct {
	*BotHandler
}

func NewGaishnik(bot *tgbotapi.Bot, handler *BotHandler) {
	gaishnik := Gaishnik{handler}
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.GetSkatCommand,
		tgbotapi.MatchTypeExact,
		gaishnik.GetSkat())

}

func (g Gaishnik) GetSkat() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		subjects, err := g.step.S.GetAllSubjectsOrderByName(ctx, true)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Учебные предметы",
			ReplyMarkup: keyboard.GetSkat(subjects),
		})
		fmt.Println(err)

	}

}
