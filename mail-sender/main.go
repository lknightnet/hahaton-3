package main

import (
	"email-service/config"
	"email-service/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	app.Run(cfg)
}
