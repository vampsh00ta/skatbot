package app

import (
	"context"
	tgbotapi "github.com/go-telegram/bot"
	"skat_bot/config"
	handlers "skat_bot/internal/handlers"
	repository "skat_bot/internal/repository"
	"skat_bot/internal/service"
	authentication "skat_bot/internal/service/auth"
	"skat_bot/internal/step_handlers"
	"skat_bot/pkg/client"
	"skat_bot/pkg/logger"

	"os"
	"os/signal"
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
	//tx := rep.GetDb()
	//a, err := rep.GetCurrentDtpByPersonId(tx, 5)
	//fmt.Println(a, err)
	auth := &authentication.AuthMap{DB: make(map[int64]*authentication.User)}
	auth.LogIn(564764193, 955, 2)

	if err != nil {
		panic(err)
	}

	srvc := service.New(rep)

	log := logger.New(cfg.Level)
	stepH := step_handlers.New(srvc, log, auth)
	opts := []tgbotapi.Option{
		//tgbotapi.WithMiddlewares(auth.AuthMiddleware()),
	}

	bot, err := tgbotapi.New(cfg.Apitoken, opts...)
	if err != nil {
		panic(err)
	}
	handlers.New(bot, stepH)
	bot.Start(ctx)
}
