package clients

import (
	"net/http"
	"net/url"
)

type ExchangeClient struct {
	token   string
	baseURL *url.URL
	client  *http.Client
}

type GetRatesResponse struct {
	Result          string             `json:"result"`
	BaseCode        string             `json:"base_code"`
	ConversionRates map[string]float64 `json:"conversion_rates"`
}
