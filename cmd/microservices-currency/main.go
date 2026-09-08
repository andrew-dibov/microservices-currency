package main

import (
	"database/sql"
	"microservices-currency/internal/clients"
	"microservices-currency/internal/configs"
	"microservices-currency/internal/loggers"
	"microservices-currency/internal/repositories"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	appConfig := configs.NewAppConfig()
	appLogger := loggers.NewAppLogger(appConfig)

	appLogger.Info("config",
		"port", appConfig.App.Port,
		"prod", appConfig.App.Prod,
		"postgres", appConfig.PostgresDatabase.Address,
	)

	/* --- --- --- */

	postgresDatabase, err := sql.Open("postgres", appConfig.PostgresDatabase.Address)
	if err != nil {
		appLogger.Error("postgresDatabase returned error", "error", err)
		os.Exit(1)
	}
	defer postgresDatabase.Close()

	postgresRepository := repositories.NewPostgresRepository(postgresDatabase)

	/* --- --- --- */

	exchangeClient := clients.NewExchangeClient(appConfig)

}
