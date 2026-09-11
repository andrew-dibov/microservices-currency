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
		return nil, fmt.Errorf("GetRates : reqURL %s : %w", appConfig.ExchangeClient.Address, err)
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
	// if len(baseCurrency) != 3 {
	// 	return nil, fmt.Errorf("baseCurrency is not 3 chars : %s", baseCurrency)
	// }

	// for _, ch := range baseCurrency {
	// 	if ch < 'A' || ch > 'Z' {
	// 		return nil, fmt.Errorf("baseCurrency has invalid chars : %s", baseCurrency)
	// 	}
	// }

	/* --- --- --- */

	client.baseURL.Path = path.Join(client.baseURL.Path, client.token, "latest", baseCurrency)

	/* --- --- --- */

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, client.baseURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("GetRates : baseCurrency %s : prepare request : %w", baseCurrency, err)
	}

	res, err := client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GetRates : baseCurrency %s : perform request : %w", baseCurrency, err)
	}
	defer res.Body.Close()

	/* --- --- --- */

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GetRates : baseCurrency %s : API response : %s", baseCurrency, res.Status)
	}

	body := http.MaxBytesReader(nil, res.Body, 10*1024*1024)

	var data GetRatesResponse
	if err := json.NewDecoder(body).Decode(&data); err != nil {
		return nil, fmt.Errorf("GetRates : baseCurrency %s : decode response : %w", baseCurrency, err)
	}

	if data.ConversionRates == nil {
		return nil, fmt.Errorf("GetRates : baseCurrency %s : miss rates", baseCurrency)
	}

	return &data, nil
}
