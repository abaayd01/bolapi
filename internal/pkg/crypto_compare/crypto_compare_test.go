package crypto_compare

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCurrentPrice(t *testing.T) {
	stubServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		mockPriceResponse := PriceResponse{Value: 100.0}
		json.NewEncoder(w).Encode(mockPriceResponse)
	}))
	defer stubServer.Close()

	client := CryptoCompareClient{BaseURL: stubServer.URL}

	result, err := client.GetCurrentPrice("ETH", "USD")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, *result, 100.0)
}
