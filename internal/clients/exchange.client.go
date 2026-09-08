package clients

import (
	"context"
	"fmt"
	"microservices-currency/internal/configs"
	"net/http"
)

func NewExchangeClient(appConfig configs.AppConfig) *ExchangeClient {
	return &ExchangeClient{
		url: fmt.Sprintf("%s%s/latest/", appConfig.ExchangeClient.Address, appConfig.ExchangeClient.Token),
		client: &http.Client{
			Timeout: appConfig.ExchangeClient.Timeout,
		},
	}
}

/* --- --- --- */

func (client *ExchangeClient) GetRates(ctx context.Context, baseCurrency string) (map[string]float64, error) {
	//
}
