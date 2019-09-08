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

func (repo *priceEvaluationRepository) Insert(priceEvaluation *models.PriceEvaluation) error {
	query := `
	`

	_, err := repo.DB.Exec(query)

	if err != nil {
		log.Printf("Error inserting PriceEvaluation")
		return err
	}

	return nil
}
