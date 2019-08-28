package cron

import (
	"bolapi/internal/pkg/crypto_compare"
	"bolapi/internal/pkg/models"
	"bolapi/internal/pkg/repositories"
	"github.com/jmoiron/sqlx"
	"github.com/robfig/cron"
	"log"
	"time"
)

type CronWorker struct {
	CryptoCompareClient *crypto_compare.CryptoCompareClient
	DB                  *sqlx.DB
}

func (cW *CronWorker) Start() error {
	c := cron.New()
	err := c.AddFunc("@every 2s", func() {
		err := cW.takePriceSnapshot()

		if err != nil {
			log.Println(err)
		}
	})

	c.Start()

	if err != nil {
		return err
	}

	return nil
}

func (cW *CronWorker) takePriceSnapshot() error {
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
		return err
	}

	log.Println("successfully took price snapshot!")
	return nil
}
