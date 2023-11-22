package main

import (
	"skat_bot/config"
	"skat_bot/internal/app"
)

func main() {

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	app.New(cfg)
}
