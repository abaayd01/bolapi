package models

import (
	"bolapi/internal/pkg/bolproto"
	"time"
)

type PriceEvaluation struct {
	bolproto.PriceEvaluationResponse
	CreatedTime time.Time
}

func New(priceEvaluationResponse *bolproto.PriceEvaluationResponse) (*PriceEvaluation, error) {
	priceEvaluation := PriceEvaluation{
		CreatedTime: time.Now(),
	}

	priceEvaluation.PriceEvaluationResponse = *priceEvaluationResponse

	return &priceEvaluation, nil
}

type PriceEvaluations []PriceEvaluation
