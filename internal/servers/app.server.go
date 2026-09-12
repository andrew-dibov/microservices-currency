package servers

import (
	"context"
	"database/sql"
	"errors"
	"microservices-currency/internal/loggers"
	"microservices-currency/internal/repositories"
	"microservices-currency/pkg/api/currency"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewAppServer(postgresRepository *repositories.PostgresRepository, appLogger *loggers.AppLogger) *AppServer {
	return &AppServer{
		postgresRepository: postgresRepository,
		appLogger:          appLogger,
	}
}

/* --- --- --- */

func (server *AppServer) Rate(ctx context.Context, request *currency.RateRequest) (*currency.RateResponse, error) {
	if request.FromCurrency == "" || request.ToCurrency == "" {
		return nil, status.Error(codes.InvalidArgument, "FromCurrency or ToCurrency is empty")
	}

	if len(request.FromCurrency) != 3 || len(request.ToCurrency) != 3 {
		return nil, status.Error(codes.InvalidArgument, "FromCurrency or ToCurrency is not 3 chars")
	}

	for _, char := range request.FromCurrency {
		if char < 'A' || char > 'Z' {
			return nil, status.Error(codes.InvalidArgument, "FromCurrency has invalid chars")
		}
	}

	for _, char := range request.ToCurrency {
		if char < 'A' || char > 'Z' {
			return nil, status.Error(codes.InvalidArgument, "ToCurrency has invalid chars")
		}
	}

	/* --- --- --- */

	rate, err := server.postgresRepository.Rate(ctx, request.FromCurrency, request.ToCurrency)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "AppServer rate not found")
		}

		server.appLogger.Error("AppServer get rate failed", "error", err)
		return nil, status.Error(codes.Internal, "AppServer get rate failed")
	}

	/* --- --- --- */

	return &currency.RateResponse{
		Rate:         rate,
		FromCurrency: request.FromCurrency,
		ToCurrency:   request.ToCurrency,
	}, nil
}

func (server *AppServer) Rates(ctx context.Context, request *currency.RatesRequest) (*currency.RatesResponse, error) {
	if request.BaseCurrency == "" {
		return nil, status.Error(codes.InvalidArgument, "BaseCurrency is empty")
	}

	if len(request.BaseCurrency) != 3 {
		return nil, status.Error(codes.InvalidArgument, "BaseCurrency is not 3 chars")
	}

	for _, char := range request.BaseCurrency {
		if char < 'A' || char > 'Z' {
			return nil, status.Error(codes.InvalidArgument, "BaseCurrency has invalid chars")
		}
	}

	/* --- --- --- */

	rates, err := server.postgresRepository.Rates(ctx, "USD")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "AppServer rates not found")
		}

		server.appLogger.Error("AppServer get rates failed", "error", err)
		return nil, status.Error(codes.Internal, "AppServer get rates failed")
	}

	/* --- --- --- */

	if request.BaseCurrency == "USD" {
		if _, ok := rates["USD"]; !ok {
			rates["USD"] = 1.0
		}

		return &currency.RatesResponse{
			Rates:        rates,
			BaseCurrency: request.BaseCurrency,
		}, nil
	}

	/* --- --- --- */

	baseRate, ok := rates[request.BaseCurrency]
	if !ok {
		server.appLogger.Warn("AppServer BaseCurrency not found", "base_currency", request.BaseCurrency)
		return nil, status.Error(codes.NotFound, "AppServer BaseCurrency not found")
	}

	newRates := make(map[string]float64, len(rates))
	for code, rateUSD := range rates {
		newRates[code] = rateUSD / baseRate
	}

	return &currency.RatesResponse{
		Rates:        newRates,
		BaseCurrency: request.BaseCurrency,
	}, nil
}
