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
		return nil, fmt.Errorf("ExchangeClient failed to parse %s : %w", appConfig.ExchangeClient.Address, err)
	}

	return &ExchangeClient{
		token:   appConfig.ExchangeClient.Token,
		baseURL: baseURL,
		client: &http.Client{
			Timeout: appConfig.ExchangeClient.Timeout,
		},
	}, nil
}

/* --- --- --- */

func (client *ExchangeClient) Rates(ctx context.Context, baseCurrency string) (*GetRatesResponse, error) {
	currentURL := *client.baseURL
	currentURL.Path = path.Join(currentURL.Path, client.token, "latest", baseCurrency)

	/* --- --- --- */

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, currentURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Rates request preparation failed : %w", err)
	}

	res, err := client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Rates request performance failed : %w", err)
	}
	defer res.Body.Close()

	/* --- --- --- */

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Rates API responded : %s", res.Status)
	}

	body := http.MaxBytesReader(nil, res.Body, 10*1024*1024)

	var data GetRatesResponse
	if err := json.NewDecoder(body).Decode(&data); err != nil {
		return nil, fmt.Errorf("Rates response decoding failed : %w", err)
	}

	if data.ConversionRates == nil {
		return nil, fmt.Errorf("Rates response rates missing")
	}

	return &data, nil
}
