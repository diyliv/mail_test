package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/diyliv/mailganer/config"
	httpservice "github.com/diyliv/mailganer/internal/delivery/http"
	"github.com/diyliv/mailganer/internal/models"
	"go.uber.org/zap"
)

type server struct {
	logger        *zap.Logger
	psqlInterface models.EmailInterface
	cfg           *config.Config
}

func NewServer(logger *zap.Logger, psqlInterface models.EmailInterface, cfg *config.Config) *server {
	return &server{
		logger:        logger,
		psqlInterface: psqlInterface,
		cfg:           cfg,
	}
}

func (s *server) StartHttp() {
	ctx := context.Background()

	s.logger.Info(fmt.Sprintf("Starting HTTP server on port %s", s.cfg.HttpServer.Port))

	handlers := httpservice.NewHttpService(s.cfg, s.logger, s.psqlInterface)

	http.Handle("/send_email", handlers.SendEmail(ctx))
	http.HandleFunc("/delay_email", handlers.DelayEmail(ctx))

	go func() {
		if err := http.ListenAndServe(s.cfg.HttpServer.Port, nil); err != nil {
			s.logger.Error("Error while serving http server: " + err.Error())
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	<-done
}
