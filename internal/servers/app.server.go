package servers

import (
	"microservices-currency/internal/loggers"
	"microservices-currency/internal/repositories"
	"microservices-currency/pkg/api/currency"
)

type CurrencyServer struct {
	currency.UnimplementedCurrencyServer
	postgresRepository *repositories.PostgresRepository
	appLogger          *loggers.AppLogger
}

func NewCurrencyServer(postgresRepository *repositories.PostgresRepository, appLogger *loggers.AppLogger) *CurrencyServer {
	return &CurrencyServer{
		postgresRepository: postgresRepository,
		appLogger:          appLogger,
	}
}

/* --- --- --- */
