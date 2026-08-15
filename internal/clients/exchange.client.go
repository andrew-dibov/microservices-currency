package clients

import (
	"context"
	"fmt"
	"microservices-currency/internal/configs"
	"net/http"
)

func NewExchangeClient(cfg *configs.AppConfig) *ExchangeClient {
	return &ExchangeClient{
		url:  fmt.Sprintf("%s%s/latest/", cfg.Api, cfg.Token),
		http: &http.Client{Timeout: cfg.Timeouts.Ex},
	}
}

func (c *ExchangeClient) Rates(ctx context.Context, baseCurrency string) (map[string]float64, error) {
}
