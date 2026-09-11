package configs

import (
	"microservices-currency/internal/modules"
	"time"
)

func NewAppConfig() AppConfig {
	return AppConfig{
		App: App{
			Name: modules.GetStringEnv("APP_NAME", "microservices-currency"),

			Prod: modules.GetBooleanEnv("APP_PROD", false),
			Port: modules.GetStringEnv("APP_PORT", "50052"),

			KeepaliveTime:    modules.GetDurationEnv("GRPC_KEEPALIVE_TIME", 10*time.Second),
			KeepaliveTimeout: modules.GetDurationEnv("GRPC_KEEPALIVE_TIMEOUT", 1*time.Second),

			ShutdownTimeout: modules.GetDurationEnv("APP_SHUTDOWN_TIMEOUT", 5*time.Second),
		},

		PostgresDatabase: PostgresDatabase{
			Address: modules.GetStringEnv("POSTGRES_ADDRESS", "postgres://app:password@postgres-currency:5432/currency?sslmode=disable"),
		},

		ExchangeClient: ExchangeClient{
			Address: modules.GetStringEnv("EXCHANGE_ADDRESS", "https://v6.exchangerate-api.com/v6/"),
			Token:   modules.GetStringEnv("EXCHANGE_TOKEN", ""),

			Timeout: modules.GetDurationEnv("EXCHANGE_TIMEOUT", 10*time.Second),
		},

		Updater: Updater{
			UpdateInterval: modules.GetDurationEnv("UPDATE_INTERVAL", 1*time.Hour),
			BaseCurrency:   modules.GetStringEnv("UPDATE_BASE_CURRENCY", "USD"),

			FetchTimeout: modules.GetDurationEnv("UPDATE_FETCH_TIMEOUT", 5*time.Second),
			StoreTimeout: modules.GetDurationEnv("UPDATE_STORE_TIMEOUT", 25*time.Second),
		},
	}
}
