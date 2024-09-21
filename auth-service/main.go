package main

import (
	"auth-service/config"
	"auth-service/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	app.Run(cfg)
}
