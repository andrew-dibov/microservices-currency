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
		"postgres", cfg.Infrastructure.Psql,
	)

	psql, err := sql.Open("postgres", cfg.Infrastructure.Psql)
	if err != nil {
		log.Error("postgres database",
			"error", err,
		)
		os.Exit(1)
	}
	defer psql.Close()

	repo := repos.NewPsqlRepo(psql)

	/* --- */

}
