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
	c := configs.NewAppConfig()

	l := slog.New(map[bool]slog.Handler{
		true:  slog.NewJSONHandler(os.Stdout, nil),
		false: slog.NewTextHandler(os.Stdout, nil),
	}[c.Prod])

	l.Info("app config",
		"port", c.Port,
		"prod", c.Prod,
		"history", c.Services.Hist,
		"conversion", c.Services.Conv,
		"postgres", c.Infra.Psql,
	)

	db, err := dbs.NewPsqlDB(&c)
	if err != nil {
		l.Error("postgres database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	rp := repos.NewPsqlRepo(db.DB)
}
