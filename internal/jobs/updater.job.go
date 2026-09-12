package jobs

import (
	"context"
	"microservices-currency/internal/clients"
	"microservices-currency/internal/configs"
	"microservices-currency/internal/loggers"
	"microservices-currency/internal/repositories"
	"time"
)

func NewUpdaterJob(appConfig *configs.AppConfig, appLogger *loggers.AppLogger, exchangeClient *clients.ExchangeClient, postgresRepository *repositories.PostgresRepository) *UpdaterJob {
	return &UpdaterJob{
		appConfig:          appConfig,
		appLogger:          appLogger,
		exchangeClient:     exchangeClient,
		postgresRepository: postgresRepository,
	}
}

/* --- --- --- */

func (updater *UpdaterJob) Start(ctx context.Context) {
	ticker := time.NewTicker(updater.appConfig.UpdaterJob.UpdateInterval)
	defer ticker.Stop()

	updater.update(ctx)

	for {
		select {
		case <-ctx.Done():
			updater.appLogger.Info("UpdaterJob stopped")
			return
		case <-ticker.C:
			updater.update(ctx)
		}
	}
}

func (updater *UpdaterJob) update(ctx context.Context) {
	fetchCtx, cancel := context.WithTimeout(ctx, updater.appConfig.UpdaterJob.FetchTimeout)
	defer cancel()

	rates, err := updater.exchangeClient.Rates(fetchCtx, "USD")
	if err != nil {
		updater.appLogger.Error("UpdaterJob failed to fetch", "error", err)
		return
	}

	/* --- --- --- */

	storeCtx, cancel := context.WithTimeout(ctx, updater.appConfig.UpdaterJob.StoreTimeout)
	defer cancel()

	if err := updater.postgresRepository.Update(storeCtx, "USD", rates); err != nil {
		updater.appLogger.Error("UpdaterJob failed to update", "error", err)
		return
	}

	/* --- --- --- */

	updater.appLogger.Info("UpdaterJob completed", "count", len(rates))
}
