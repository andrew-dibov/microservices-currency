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

			KeepaliveTime:    modules.GetDurationEnv("APP_KEEPALIVE_TIME", 5*time.Second),
			KeepaliveTimeout: modules.GetDurationEnv("APP_KEEPALIVE_TIMEOUT", 5*time.Second),

			ShutdownTimeout: modules.GetDurationEnv("APP_SHUTDOWN_TIMEOUT", 5*time.Second),
		},

		PostgresDatabase: PostgresDatabase{
			DSN: "postgres://" +
				modules.GetStringEnv("POSTGRES_USER", "app") +
				":" + modules.GetStringEnv("POSTGRES_PASSWORD", "1234") +
				"@" + modules.GetStringEnv("POSTGRES_HOST", "postgres-currency") +
				":" + modules.GetStringEnv("POSTGRES_PORT", "5432") +
				"/" + modules.GetStringEnv("POSTGRES_DATABASE", "currency") +
				"?" + modules.GetStringEnv("POSTGRES_PARAMETERS", "sslmode=disable"),
		},

		ExchangeClient: ExchangeClient{
			Token:   modules.GetStringEnv("EXCHANGE_TOKEN", ""),
			Address: modules.GetStringEnv("EXCHANGE_ADDRESS", "https://v6.exchangerate-api.com/v6/"),
			Timeout: modules.GetDurationEnv("EXCHANGE_TIMEOUT", 10*time.Second),
		},

		UpdaterJob: UpdaterJob{
			UpdateInterval: modules.GetDurationEnv("UPDATER_UPDATE_INTERVAL", 1*time.Hour),
			FetchTimeout:   modules.GetDurationEnv("UPDATER_FETCH_TIMEOUT", 15*time.Second),
			StoreTimeout:   modules.GetDurationEnv("UPDATER_STORE_TIMEOUT", 25*time.Second),
		},
	}
}
