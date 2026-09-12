package jobs

import (
	"microservices-currency/internal/clients"
	"microservices-currency/internal/configs"
	"microservices-currency/internal/loggers"
	"microservices-currency/internal/repositories"
)

type UpdaterJob struct {
	appConfig          *configs.AppConfig
	appLogger          *loggers.AppLogger
	exchangeClient     *clients.ExchangeClient
	postgresRepository *repositories.PostgresRepository
}
