package main

import (
	_ "github.com/lib/pq"
	"skat_bot/config"
	"skat_bot/internal/app"
)

// Send any text message to the bot after the bot has been started

//func main() {
//	cfg, _ := config.New()
//
//	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
//	defer cancel()
//
//	opts := []bot.Option{}
//
//	b, err := bot.New(cfg.Apitoken, opts...)
//	if nil != err {
//		// panics for the sake of simplicity.
//		// you should handle this error properly in your code.
//		panic(err)
//	}
//	b.RegisterHandler(bot.HandlerTypeMessageText, "/hello", bot.MatchTypeExact, handler)
//	b.Start(ctx)
//}
//
//func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
//	if update.InlineQuery == nil {
//		return
//	}
//
//	results := []models.InlineQueryResult{
//		&models.InlineQueryResultArticle{ID: "1", Title: "Foo 1", InputMessageContent: &models.InputTextMessageContent{MessageText: "foo 1"}},
//		&models.InlineQueryResultArticle{ID: "2", Title: "Foo 2", InputMessageContent: &models.InputTextMessageContent{MessageText: "foo 2"}},
//		&models.InlineQueryResultArticle{ID: "3", Title: "Foo 3", InputMessageContent: &models.InputTextMessageContent{MessageText: "foo 3"}},
//	}
//
//	b.AnswerInlineQuery(ctx, &bot.AnswerInlineQueryParams{
//		InlineQueryID: update.InlineQuery.ID,
//		Results:       results,
//	})
//}

func main() {

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	app.New(cfg)
}
