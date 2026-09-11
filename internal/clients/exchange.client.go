package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"microservices-currency/internal/configs"
	"net/http"
	"net/url"
	"path"
)

func NewExchangeClient(appConfig *configs.AppConfig) (*ExchangeClient, error) {
	baseURL, err := url.Parse(appConfig.ExchangeClient.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to parse %s : %w", appConfig.ExchangeClient.Address, err)
	}

	return &ExchangeClient{
		baseURL: baseURL,
		token:   appConfig.ExchangeClient.Token,
		client: &http.Client{
			Timeout: appConfig.ExchangeClient.Timeout,
		},
	}, nil
}

/* --- --- --- */

func (client *ExchangeClient) GetRates(ctx context.Context, baseCurrency string) (*GetRatesResponse, error) {
	client.baseURL.Path = path.Join(client.baseURL.Path, client.token, "latest", baseCurrency)

	/* --- --- --- */

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, client.baseURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare request : %w", err)
	}

	res, err := client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request : %w", err)
	}
	defer res.Body.Close()

	/* --- --- --- */

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API response : %s", res.Status)
	}

	body := http.MaxBytesReader(nil, res.Body, 10*1024*1024)

	var data GetRatesResponse
	if err := json.NewDecoder(body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode response : %w", err)
	}

	if data.ConversionRates == nil {
		return nil, fmt.Errorf("rates missed")
	}

	return &data, nil
}
