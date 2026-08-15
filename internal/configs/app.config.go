package configs

import (
	"microservices-currency/internal/tools"
	"time"
)

func NewAppConfig() AppConfig {
	return AppConfig{
		Port: tools.GetStrEnv("PORT", "8080"),
		Prod: tools.GetBoolEnv("PROD", false),

		Api:   tools.GetStrEnv("EX_API", "https://v6.exchangerate-api.com/v6/"),
		Token: tools.GetStrEnv("EX_TKN", ""),

		Cert: tools.GetStrEnv("TLS_CERT", ""),
		Key:  tools.GetStrEnv("TLS_KEY", ""),

		Services: Services{
			Hist: tools.GetStrEnv("HIST_ADDR", "localhost:50051"),
			Curr: tools.GetStrEnv("CURR_ADDR", "localhost:50052"),
			Conv: tools.GetStrEnv("CONV_ADDR", "localhost:50053"),
		},

		Infra: Infra{
			Psql: tools.GetStrEnv("PSQL_ADDR", "postgres://currency:1234@psql-currency:5432/currency"),

			PsqlMaxOpen:     tools.GetIntEnv("PSQL_MAX_OPEN", 25),
			PsqlMaxIdle:     tools.GetIntEnv("PSQL_MAX_IDLE", 25),
			PsqlMaxLifetime: tools.GetDurEnv("PSQL_MAX_LIFE", 5*time.Minute),
		},

		Timeouts: Timeouts{
			Ex: tools.GetDurEnv("EX_TOUT", 5*time.Second),

			Hist: tools.GetDurEnv("HIST_TOUT", 5*time.Second),
			Curr: tools.GetDurEnv("CURR_TOUT", 5*time.Second),
			Conv: tools.GetDurEnv("CONV_TOUT", 5*time.Second),

			Read:     tools.GetDurEnv("READ_TOUT", 10*time.Second),
			Idle:     tools.GetDurEnv("IDLE_TOUT", 15*time.Second),
			Write:    tools.GetDurEnv("WRITE_TOUT", 20*time.Second),
			Shutdown: tools.GetDurEnv("SHUTDOWN_TOUT", 25*time.Second),
		},

		Limits: Limits{},
	}
}
