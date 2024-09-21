package app

import (
	"auth-service/config"
	"auth-service/internal/controller/http"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"auth-service/pkg/database/postgres"
	"auth-service/pkg/mail"
	"auth-service/pkg/server"
	"auth-service/pkg/token"
	"context"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	ctx := context.Background()

	pg, err := postgres.New(ctx, cfg.PG.ConnURI)
	if err != nil {
		log.Println(err)
	}

	repositories := repository.NewRepositories(pg)

	deps := &service.Dependencies{
		Salt:    cfg.Hash.Salt,
		DBUser:  repositories.Auth,
		DBToken: repositories.Token,
		JWT:     token.NewJWTDependencies(cfg.JWT.Singkey),
		Email:   mail.NewMailSender(cfg.Api.Url),
		Topic:   cfg.Email.Topic,
		Link:    cfg.Email.Link,
	}

	services := service.NewServices(deps)

	rout := mux.NewRouter()

	http.NewAuthRoutes(services.Auth, rout)
	srv := server.NewServer(rout, server.Port(cfg.HTTP.Port), server.ReadTimeout(cfg.ReadTimeout),
		server.WriteTimeout(cfg.WriteTimeout), server.ShutdownTimeout(cfg.ShutdownTimeout))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("Run: " + s.String())
	case err := <-srv.Notify():
		log.Println(errors.Wrap(err, "Run: signal.Notify"))
	}

	err = srv.Shutdown()
	if err != nil {
		log.Println(errors.Wrap(err, "Run: server shutdown"))
	}
}
