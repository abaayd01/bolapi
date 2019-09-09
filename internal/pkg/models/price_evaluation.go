package models

import (
	"bolapi/internal/pkg/bolproto"
	"time"
)

type PriceEvaluation struct {
	bolproto.PriceEvaluationResponse
	CreatedTime     time.Time
	PriceSnapshotId int
}

func NewPriceEvaluation(priceEvaluationResponse *bolproto.PriceEvaluationResponse, priceSnapshot *PriceSnapshot) (*PriceEvaluation, error) {
	priceEvaluation := PriceEvaluation{
		CreatedTime:     time.Now(),
		PriceSnapshotId: priceSnapshot.Id,
	}

	priceEvaluation.PriceEvaluationResponse = *priceEvaluationResponse

	return &priceEvaluation, nil
}

type PriceEvaluations []PriceEvaluation
