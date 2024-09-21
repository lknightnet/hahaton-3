package app

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"os"
	"os/signal"
	"student-resume/config"
	"student-resume/internal/controller/http"
	"student-resume/internal/repository"
	"student-resume/internal/service"
	"student-resume/pkg/database/postgres"
	"student-resume/pkg/server"
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
		DBResume:  repositories.Resume,
		DBStudent: repositories.Student,
	}

	services := service.NewServices(deps)

	rout := mux.NewRouter()

	http.NewResumeRoutes(services.Resume, rout)
	http.NewStudentRoutes(services.Student, rout)
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
