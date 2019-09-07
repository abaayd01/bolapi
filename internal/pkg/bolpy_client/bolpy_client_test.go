package bolpy_client

import (
	"bolapi/internal/pkg/bolproto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvaluatePriceViaGrpc(t *testing.T) {
	bolpyCtx, _ := NewBolpyContext("localhost:50051")
	bolpyClient := BolpyClient{Context: bolpyCtx}

	var currentPrice float32 = 20.0

	var historicalPrices []float32

	for i := 0; i <= 100; i++ {
		historicalPrices = append(historicalPrices, float32(i))
	}

	r, _ := bolpyClient.EvaluatePrice(&bolproto.PriceEvaluationRequest{
		CurrentPrice:     currentPrice,
		HistoricalPrices: historicalPrices,
	})

	assert.Equal(t, "BUY", r.Action)
}
