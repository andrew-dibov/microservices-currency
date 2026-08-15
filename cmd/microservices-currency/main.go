package main

import (
	"microservices-currency/internal/configs"
	"microservices-currency/internal/dbs"
	"microservices-currency/internal/repos"

	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	cfg := configs.NewAppConfig()

	log := slog.New(map[bool]slog.Handler{
		true:  slog.NewJSONHandler(os.Stdout, nil),
		false: slog.NewTextHandler(os.Stdout, nil),
	}[cfg.Prod])

	log.Info("app config",
		"port", cfg.Port,
		"prod", cfg.Prod,
		"history", cfg.Services.Hist,
		"conversion", cfg.Services.Conv,
		"postgres", cfg.Infra.Psql,
	)

	db, err := dbs.NewPsqlDb(&cfg)
	if err != nil {
		log.Error("postgres database",
			"error", err,
		)
		os.Exit(1)
	}
	defer db.Close()

	repo := repos.NewPsqlRepo(db.DB)
}
