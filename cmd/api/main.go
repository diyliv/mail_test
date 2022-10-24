package main

import (
	"github.com/diyliv/mailganer/config"
	"github.com/diyliv/mailganer/internal/repository"
	"github.com/diyliv/mailganer/internal/server"
	"github.com/diyliv/mailganer/pkg/logger"
	"github.com/diyliv/mailganer/pkg/storage/postgres"
)

func main() {
	cfg := config.ReadConfig()
	logger := logger.InitLogger()
	psql := postgres.ConnPostgres(cfg)
	psqlRepo := repository.NewPostgres(logger, psql)

	server := server.NewServer(logger, psqlRepo, cfg)
	server.StartHttp()
}
