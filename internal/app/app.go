package app

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"os"
	"os/signal"
	"skat_bot/config"
	handlers "skat_bot/internal/handlers"
	repository "skat_bot/internal/repository"
	"skat_bot/internal/service"
	"skat_bot/pkg/client"
	"skat_bot/pkg/logger"
	"syscall"
)

func New(cfg *config.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	db, err := client.NewPostgresClient(ctx, 5, cfg.PG)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	rep := repository.New(db)
	//fmt.Println(res, a)
	//tx := rep.GetDb()
	//a, err := rep.GetAllSubjectsOrderByName(ctx, true)
	//fmt.Println(a, err)
	//auth := &authentication.AuthMap{DB: make(map[int64]*authentication.User)}
	//auth.LogIn(564764193, 955, 2)

	if err != nil {
		panic(err)
	}

	srvc := service.New(rep)
	//err = srvc.DownloadVariant(ctx, models.Variant{FilePath: "documents/file_0.docx", Name: "xyu"})
	fmt.Println(err)
	log := logger.New(cfg.Level)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	opts := []tgbotapi.Option{
		//tgbotapi.WithDefaultHandler(handler),
	}

	bot, err := tgbotapi.New(cfg.Apitoken, opts...)
	if err != nil {
		panic(err)
	}

	handlers.New(bot, srvc, log)

	//bot.SetWebhook(ctx, &tgbotapi.SetWebhookParams{
	//	URL: fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cfg.Apitoken),
	//})

	//go func() {
	//	http.ListenAndServe(":2000", bot.WebhookHandler())
	//}()

	// Use StartWebhook instead of Start

	bot.Start(ctx)

}

func handler(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}
