package models

import (
	"bolapi/internal/pkg/bolproto"
	"time"
)

type PriceEvaluation struct {
	PriceSnapshotId int       `db:"price_snapshot_id"`
	CreatedTime     time.Time `db:"created_time"`
	Action          string    `db:"action"`
	EvaluationPrice float32   `db:"evaluation_price"`
	TargetExitPrice float32   `db:"target_exit_price"`
	StopLossPrice   float32   `db:"stop_loss_price"`
	BolUpper        float32   `db:"bol_upper"`
	BolLower        float32   `db:"bol_lower"`
	MovingAverage   float32   `db:"moving_average"`
}

func NewPriceEvaluation(priceEvaluationResponse *bolproto.PriceEvaluationResponse, priceSnapshot *PriceSnapshot) (*PriceEvaluation, error) {
	priceEvaluation := PriceEvaluation{
		PriceSnapshotId: priceSnapshot.Id,
		CreatedTime:     time.Now(),
		Action:          priceEvaluationResponse.Action,
		EvaluationPrice: priceEvaluationResponse.EvaluationPrice,
		TargetExitPrice: priceEvaluationResponse.TargetExitPrice,
		StopLossPrice:   priceEvaluationResponse.StopLossPrice,
		BolUpper:        priceEvaluationResponse.BolUpper,
		BolLower:        priceEvaluationResponse.BolLower,
		MovingAverage:   priceEvaluationResponse.MovingAverage,
	}

	return &priceEvaluation, nil
}

type PriceEvaluations []PriceEvaluation
