package cron

import (
	"bolapi/internal/pkg/bolproto"
	"bolapi/internal/pkg/bolpy_client"
	"bolapi/internal/pkg/crypto_compare"
	"bolapi/internal/pkg/database"
	"bolapi/internal/pkg/models"
	"bolapi/internal/pkg/repositories"
	"github.com/robfig/cron"
	"log"
	"time"
)

type Worker struct {
	CryptoCompareClient *crypto_compare.CryptoCompareClient
	BolpyClient         *bolpy_client.BolpyClient
	DB                  database.DBInterface
}

func (cW *Worker) Start() error {
	c := cron.New()
	err := c.AddFunc("@every 2s", func() {
		priceSnapshot, _ := cW.takePriceSnapshot()
		priceEvaluationResponse, _ := cW.evaluatePriceSnapshot(priceSnapshot)
		_ = cW.savePriceEvaluationResponse(priceEvaluationResponse, priceSnapshot)
		log.Println(priceEvaluationResponse)
	})

	if err != nil {
		return err
	}

	c.Start()
	return nil
}

func (cW *Worker) takePriceSnapshot() (*models.PriceSnapshot, error) {
	log.Println("taking price snapshot...")
	currentPrice, _ := cW.CryptoCompareClient.GetCurrentPrice("ETH", "USD")

	priceSnapshot := models.PriceSnapshot{
		CreatedTime: time.Now(),
		Price:       *currentPrice,
	}

	priceSnapshotRepo := repositories.NewPriceSnapshotRepository(cW.DB)

	_, err := priceSnapshotRepo.Insert(&priceSnapshot)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("successfully took price snapshot!")
	return &priceSnapshot, nil
}

func (cW *Worker) evaluatePriceSnapshot(priceSnapshot *models.PriceSnapshot) (*bolproto.PriceEvaluationResponse, error) {
	priceSnapshotRepo := repositories.NewPriceSnapshotRepository(cW.DB)
	historicalPriceSnapshots, err := priceSnapshotRepo.GetLatest(15)

	if err != nil {
		return nil, err
	}

	return cW.BolpyClient.EvaluatePrice(&bolproto.PriceEvaluationRequest{
		CurrentPrice:     float32(priceSnapshot.Price),
		HistoricalPrices: historicalPriceSnapshots.TransformToFloatSlice(),
	})
}

func (cW *Worker) savePriceEvaluationResponse(priceEvaluationResponse *bolproto.PriceEvaluationResponse, priceSnapshot *models.PriceSnapshot) error {
	priceEvaluation, err := models.NewPriceEvaluation(priceEvaluationResponse, priceSnapshot)

	if err != nil {
		log.Println(err)
	}

	priceEvaluationRepo := repositories.NewPriceEvaluationRepository(cW.DB)

	_, _ = priceEvaluationRepo.Insert(priceEvaluation)

	return nil
}
