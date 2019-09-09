package repositories

import (
	"bolapi/internal/pkg/database"
	"bolapi/internal/pkg/models"
	"log"
)

type PriceEvaluationRepository interface {
	Insert(priceEvaluation *models.PriceEvaluation) error
}

type priceEvaluationRepository struct {
	DB database.DBInterface
}

func NewPriceEvaluationRepository(DB database.DBInterface) *priceEvaluationRepository {
	return &priceEvaluationRepository{DB: DB}
}

func (repo *priceEvaluationRepository) Insert(priceEvaluation *models.PriceEvaluation) (*models.PriceEvaluation, error) {
	query := `
		INSERT INTO price_evaluations (
		    price_snapshot_id,
			created_time,
			action,
			evaluation_price,
			target_exit_price,
			stop_loss_price,
			bol_upper,
			bol_lower,
			moving_average
		) VALUES (
			(SELECT price_snapshot_id FROM price_snapshots WHERE price_snapshot_id = $1),
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9
		)
	`

	_, err := repo.DB.Exec(query,
		priceEvaluation.PriceSnapshotId,
		priceEvaluation.CreatedTime,
		priceEvaluation.Action,
		priceEvaluation.EvaluationPrice,
		priceEvaluation.TargetExitPrice,
		priceEvaluation.StopLossPrice,
		priceEvaluation.BolUpper,
		priceEvaluation.BolLower,
		priceEvaluation.MovingAverage,
	)

	if err != nil {
		log.Printf("Error inserting PriceEvaluation: %s", err)
		return nil, err
	}

	return priceEvaluation, nil
}
