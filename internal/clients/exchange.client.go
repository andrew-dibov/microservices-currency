package clients

import (
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

// func (c *ExchangeClient) Rates(ctx context.Context, baseCurrency string) (map[string]float64, error) {
// 	url := fmt.Sprintf("%s%s", c.url, baseCurrency)

// 	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create req %v", err)
// 	}

// 	res, err := c.http.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get rates : %v", err)
// 	}

// 	defer res.Body.Close()
// }
