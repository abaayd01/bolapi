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
		currentPrice, _ := cW.takePriceSnapshot()
		priceEvaluationResponse, _ := cW.evaluatePrice(currentPrice)
		log.Println(priceEvaluationResponse)
	})

	if err != nil {
		return err
	}

	c.Start()
	return nil
}

func (cW *Worker) takePriceSnapshot() (*float64, error) {
	log.Println("taking price snapshot...")
	currentPrice, _ := cW.CryptoCompareClient.GetCurrentPrice("ETH", "USD")

	priceSnapshot := models.PriceSnapshot{
		CreatedTime: time.Now(),
		Price:       *currentPrice,
	}

	priceSnapshotRepo := repositories.NewPriceSnapshotRepository(cW.DB)

	err := priceSnapshotRepo.Insert(&priceSnapshot)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("successfully took price snapshot!")
	return currentPrice, nil
}

func (cW *Worker) evaluatePrice(price *float64) (*bolproto.PriceEvaluationResponse, error) {
	// should fetch the historical prices in here
	priceSnapshotRepo := repositories.NewPriceSnapshotRepository(cW.DB)
	historicalPriceSnapshots, err := priceSnapshotRepo.GetLatest(15)

	if err != nil {
		return nil, err
	}

	return cW.BolpyClient.EvaluatePrice(&bolproto.PriceEvaluationRequest{
		CurrentPrice:     float32(*price),
		HistoricalPrices: historicalPriceSnapshots.TransformToFloatSlice(),
	})
}

func (cW *Worker) savePriceEvaluationResponse(priceEvaluationResponse *bolproto.PriceEvaluationResponse) error {

	return nil
}
