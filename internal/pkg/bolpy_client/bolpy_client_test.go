package bolpy_client

import (
	"bolapi/internal/pkg/bolproto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvaluatePriceViaGrpc(t *testing.T) {
	bolpyCtx, _ := NewBolpyContext("localhost:50051")
	bolpyClient := BolpyClient{Context: bolpyCtx}

	r, _ := bolpyClient.EvaluatePrice(&bolproto.PriceEvaluationRequest{
		CurrentPrice:     nil,
		HistoricalPrices: nil,
	})

	assert.Equal(t, r.Action, "BUY")
}
