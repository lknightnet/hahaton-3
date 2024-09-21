package app

import (
	"email-service/config"
	"email-service/internal/controller/http"
	"email-service/internal/service"
	"email-service/pkg/server"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"net/smtp"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {

	fmt.Println("\n  ██████ ▄▄▄█████▓ ██▓  █████▒██▓     ██▓ ███▄    █   ▄████     ███▄ ▄███▓ ▄▄▄       ███▄    █   ██████ \n▒██    ▒ ▓  ██▒ ▓▒▓██▒▓██   ▒▓██▒    ▓██▒ ██ ▀█   █  ██▒ ▀█▒   ▓██▒▀█▀ ██▒▒████▄     ██ ▀█   █ ▒██    ▒ \n░ ▓██▄   ▒ ▓██░ ▒░▒██▒▒████ ░▒██░    ▒██▒▓██  ▀█ ██▒▒██░▄▄▄░   ▓██    ▓██░▒██  ▀█▄  ▓██  ▀█ ██▒░ ▓██▄   \n  ▒   ██▒░ ▓██▓ ░ ░██░░▓█▒  ░▒██░    ░██░▓██▒  ▐▌██▒░▓█  ██▓   ▒██    ▒██ ░██▄▄▄▄██ ▓██▒  ▐▌██▒  ▒   ██▒\n▒██████▒▒  ▒██▒ ░ ░██░░▒█░   ░██████▒░██░▒██░   ▓██░░▒▓███▀▒   ▒██▒   ░██▒ ▓█   ▓██▒▒██░   ▓██░▒██████▒▒\n▒ ▒▓▒ ▒ ░  ▒ ░░   ░▓   ▒ ░   ░ ▒░▓  ░░▓  ░ ▒░   ▒ ▒  ░▒   ▒    ░ ▒░   ░  ░ ▒▒   ▓▒█░░ ▒░   ▒ ▒ ▒ ▒▓▒ ▒ ░\n░ ░▒  ░ ░    ░     ▒ ░ ░     ░ ░ ▒  ░ ▒ ░░ ░░   ░ ▒░  ░   ░    ░  ░      ░  ▒   ▒▒ ░░ ░░   ░ ▒░░ ░▒  ░ ░\n░  ░  ░    ░       ▒ ░ ░ ░     ░ ░    ▒ ░   ░   ░ ░ ░ ░   ░    ░      ░     ░   ▒      ░   ░ ░ ░  ░  ░  \n      ░            ░             ░  ░ ░           ░       ░           ░         ░  ░         ░       ░  \n                                                                                                        \n")

	deps := &service.Dependencies{
		SmtpHost: cfg.Mail.Host,
		SmtpPort: cfg.Mail.Port,
		Auth:     smtp.PlainAuth("", cfg.Mail.From, cfg.Mail.Password, cfg.Mail.Host),
		From:     cfg.Mail.From,
	}

	services := service.NewServices(deps)

	rout := mux.NewRouter()

	http.NewMailRoute(rout, services.Mail)
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

	err := srv.Shutdown()
	if err != nil {
		log.Println(errors.Wrap(err, "Run: server shutdown"))
	}

}
