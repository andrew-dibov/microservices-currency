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
			Port: modules.GetStringEnv("APP_PORT", "8080"),

			ReadTimeout:     modules.GetDurationEnv("APP_READ_TIMEOUT", 5*time.Second),
			IdleTimeout:     modules.GetDurationEnv("APP_IDLE_TIMEOUT", 5*time.Second),
			WriteTimeout:    modules.GetDurationEnv("APP_WRITE_TIMEOUT", 5*time.Second),
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
	}
}
