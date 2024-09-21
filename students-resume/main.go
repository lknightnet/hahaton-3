package main

import (
	"student-resume/config"
	"student-resume/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	app.Run(cfg)
}
