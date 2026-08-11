package configs

import (
	"microservices-currency/internal/tools"
	"time"
)

func NewAppConfig() AppConfig {
	return AppConfig{
		Port: tools.GetStrEnv("PORT", "8080"),
		Prod: tools.GetBoolEnv("PROD", false),

		Api:   tools.GetStrEnv("EX_URL", "https://v6.exchangerate-api.com/v6/"),
		Token: tools.GetStrEnv("EX_TOKEN", ""),

		Services: Services{
			Hist: tools.GetStrEnv("HIST_ADDR", "localhost:50051"),
			Curr: tools.GetStrEnv("CURR_ADDR", "localhost:50052"),
			Conv: tools.GetStrEnv("CONV_ADDR", "localhost:50053"),
		},

		Infrastructure: Infrastructure{
			Psql: tools.GetStrEnv("PSQL_ADDR", "postgres://currency:1234@psql-currency:5432/currency"),
		},

		Timeouts: Timeouts{
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
