package app

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"skat_bot/config"
	"skat_bot/internal/handlers"
	repository "skat_bot/internal/repository"
	"skat_bot/internal/service"
	"skat_bot/pkg/client"
	"skat_bot/pkg/logger"
	"syscall"
)

func NewWebhook(cfg *config.Config) {
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

	if err != nil {
		panic(err)
	}

	srvc := service.New(rep)
	log := logger.New(cfg.Level)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	opts := []tgbotapi.Option{}

	bot, err := tgbotapi.New(cfg.Apitoken, opts...)
	if err != nil {
		panic(err)
	}
	bot.SetWebhook(ctx, &tgbotapi.SetWebhookParams{
		URL: cfg.BaseURL + "/webhook" + cfg.Apitoken,
	})
	if err != nil {
		panic(err)
	}
	// gin router
	router := gin.New()
	router.Use(gin.Logger())
	handlers.New(bot, srvc, log)

	// telegram

	go func() {
		http.ListenAndServe(":"+cfg.Http.Port, bot.WebhookHandler())
	}()

	// Use StartWebhook instead of Start
	bot.StartWebhook(ctx)

}

func webhookHandler(c *gin.Context) {
	defer c.Request.Body.Close()

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var update tgmodels.Update
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		log.Println(err)
		return
	}

	// to monitor changes run: heroku logs --tail
	log.Printf("From: %+v Text: %+v\n", update.Message.From, update.Message.Text)
}

func NewPooling(cfg *config.Config) {
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
	log := logger.New(cfg.Level)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	opts := []tgbotapi.Option{
		//tgbotapi.WithHTTPClient(time.Millisecond*20,httpserver.Port(cfg.Http.Port))

		//tgbotapi.WithMiddlewares(handlers.BreakSkat),
	}

	bot, err := tgbotapi.New(cfg.Apitoken, opts...)
	if err != nil {
		panic(err)
	}
	handlers.New(bot, srvc, log)
	bot.DeleteWebhook(ctx, &tgbotapi.DeleteWebhookParams{
		true,
	})
	bot.Start(ctx)

}
