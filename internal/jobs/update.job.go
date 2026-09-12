package jobs

// import (
// 	"context"
// 	"microservices-currency/internal/clients"
// 	"microservices-currency/internal/configs"
// 	"microservices-currency/internal/loggers"
// 	"microservices-currency/internal/repositories"
// 	"time"
// )

// type Updater struct {
// 	appConfig          *configs.AppConfig
// 	appLogger          *loggers.AppLogger
// 	exchangeClient     *clients.ExchangeClient
// 	postgresRepository *repositories.PostgresRepository
// }

// func NewUpdater(appConfig *configs.AppConfig, appLogger *loggers.AppLogger, exchangeClient *clients.ExchangeClient, postgresRepository *repositories.PostgresRepository) Updater {
// 	return Updater{
// 		appConfig:          appConfig,
// 		appLogger:          appLogger,
// 		exchangeClient:     exchangeClient,
// 		postgresRepository: postgresRepository,
// 	}
// }

// func (updater *Updater) Start(ctx context.Context) {
// 	ticker := time.NewTicker(updater.appConfig.UpdaterJob.UpdateInterval)
// 	defer ticker.Stop()

// 	updater.update(ctx)

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			updater.appLogger.Info("updater stopped")
// 			return
// 		case <-ticker.C:
// 			updater.update(ctx)
// 		}
// 	}
// }

// func (updater *Updater) update(ctx context.Context) {
// 	ctx, cancel := context.WithTimeout(ctx, updater.appConfig.UpdaterJob.FetchTimeout)
// 	defer cancel()
// 	f
// 	rates, err := updater.exchangeClient.GetRates(ctx, updater.appConfig.UpdaterJob.BaseCurrency)
// 	if err != nil {
// 		updater.appLogger.Error("fetch rates failed", "error", err)
// 		return
// 	}

// 	ctx, cancel = context.WithTimeout(ctx, updater.appConfig.UpdaterJob.StoreTimeout)
// 	defer cancel()

// 	if err := updater.postgresRepository.UpdateRates(ctx, updater.appConfig.UpdaterJob.BaseCurrency, rates.ConversionRates); err != nil {
// 		updater.appLogger.Error("update rates failed", "error", err)
// 		return
// 	}

// 	updater.appLogger.Info("rates updated", "count", len(rates.ConversionRates))
// }
