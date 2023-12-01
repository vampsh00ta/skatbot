package service

import (
	"context"
	tgbotapi "github.com/go-telegram/bot"
	"skat_bot/internal/repository/models"
)

type senderFunc func(context.Context, *tgbotapi.Bot, ResultDownload)

type taskDownload struct {
	ChatId  int64
	variant models.Variant
}
type ResultDownload struct {
	ChatId   int64
	FileName string
	File     *[]byte
	Error    error
}

type Download struct {
	Result chan ResultDownload
	bot    *tgbotapi.Bot

	taskChan chan taskDownload
	quit     chan struct{}

	s Service
}

func NewDownload(service Service) (*Download, error) {

	return &Download{
		s:        service,
		taskChan: make(chan taskDownload, 100),
		Result:   make(chan ResultDownload, 100),
		quit:     make(chan struct{}),
	}, nil
}

func (d *Download) Start(ctx context.Context, bot *tgbotapi.Bot, num int, f senderFunc) {
	d.startWorkers(ctx, num)
	go func() {
		select {
		case res := <-d.Result:
			f(ctx, bot, res)
		}
	}()
}

func (d *Download) startWorkers(ctx context.Context, num int) {
	for i := 0; i < num; i++ {
		go func(workerNum int) {
			for {
				select {
				case <-d.quit:
					return
				case task, ok := <-d.taskChan:
					if !ok {
						return
					}
					//time.Sleep(time.Second * 10)
					//d.Result <- ResultDownload{task.ChatId, "", &[]byte{}, errors.New("")}
					//
					fileName, file, err := d.s.DownloadVariant(ctx, task.variant)
					d.Result <- ResultDownload{task.ChatId, fileName, file, err}

				}
			}
		}(i)
	}
}
func (d *Download) AddWork(chatId int64, variant models.Variant) {
	go func(chatId int64, variant models.Variant) {
		select {
		case d.taskChan <- taskDownload{chatId, variant}:
		case <-d.quit:
		}
	}(chatId, variant)
}
