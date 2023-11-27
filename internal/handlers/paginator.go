package handlers

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
	"skat_bot/internal/keyboard"
	"strconv"
	"strings"
)

func (h BotHandler) PageInstitutePaginator() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *tgmodels.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		buttonData := update.CallbackQuery.Data
		pageStr := strings.Split(buttonData, "_")[1]
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
				ChatID: update.CallbackQuery.Message.Chat.ID,
				Text:   "Что-то пошло не так",
			})
		}
		insts, err := h.service.GetAllInstitutes(ctx, true)
		insts = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
		if err != nil {
			_, err = b.SendMessage(ctx, &tgbotapi.SendMessageParams{
				ChatID: update.CallbackQuery.Message.Chat.ID,
				Text:   "Что-то пошло не так",
			})
		}
		fmt.Println(page)
		_, err = b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.ID,

			ReplyMarkup: keyboard.InstituteNumsTest(insts, page),
		})
	}
}
