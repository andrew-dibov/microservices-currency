package main

import (
	"microservices-currency/internal/configs"
	"microservices-currency/internal/repos"

	"database/sql"
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

	psql, err := sql.Open("postgres", cfg.Infra.Psql)
	if err != nil {
		log.Error("postgres database",
			"error", err,
		)
		os.Exit(1)
	}
	defer psql.Close()

	psql.SetMaxOpenConns(cfg.Psql.MaxOpen)
	psql.SetMaxIdleConns(cfg.Psql.MaxIdle)
	psql.SetConnMaxLifetime(cfg.Psql.MaxLifetime)

	repo := repos.NewPsqlRepo(psql)
}
