package clients

import "net/http"

type ExchangeClient struct {
	url string
	cl  *http.Client
}

type ExchangeResponse struct {
	Result          string             `json:"result"`
	BaseCode        string             `json:"base_code"`
	ConversionRates map[string]float64 `json:"conversion_rates"`
}
