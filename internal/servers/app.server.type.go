package servers

import (
	"microservices-currency/internal/loggers"
	"microservices-currency/internal/repositories"
	"microservices-currency/pkg/api/currency"
)

type AppServer struct {
	currency.UnimplementedCurrencyServer
	postgresRepository *repositories.PostgresRepository
	appLogger          *loggers.AppLogger
}
