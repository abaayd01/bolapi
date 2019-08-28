package crypto_compare

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

var APIStub *httptest.Server

func init() {
	APIStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		mockPriceResponse := PriceResponse{Value: 100.0}
		json.NewEncoder(w).Encode(mockPriceResponse)
	}))
}
