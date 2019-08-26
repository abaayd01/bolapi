package crypto_compare

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	CryptoCompareBaseURL = "https://min-api.cryptocompare.com"
)

type CryptoCompareClient struct {
	BaseURL string
}

type PriceResponse struct {
	Value float64 `json:"USD"`
}

func (c *CryptoCompareClient) GetCurrentPrice(fsym string, tsyms string) (*float64, error) {
	reqUrl := fmt.Sprintf("%s/data/price?fsym=%s&tsyms=%s", c.BaseURL, fsym, tsyms)

	resp, err := http.Get(reqUrl)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	var priceResponse PriceResponse
	err = json.NewDecoder(resp.Body).Decode(&priceResponse)

	if err != nil {
		return nil, err
	}

	return &priceResponse.Value, nil
}
