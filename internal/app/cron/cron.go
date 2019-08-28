package cron

import (
	"bolapi/internal/pkg/crypto_compare"
	"github.com/robfig/cron"
	"log"
)

type CronWorker struct {
	CryptoCompareClient *crypto_compare.CryptoCompareClient
}

func (cW *CronWorker) Start() error {
	c := cron.New()
	err := c.AddFunc("@every 10s", func() { cW.takePriceSnapshot() })

	c.Start()

	if err != nil {
		return err
	}

	return nil
}

func (cW *CronWorker) takePriceSnapshot() {
	log.Println("taking snapshot")
	result, _ := cW.CryptoCompareClient.GetCurrentPrice("ETH", "USD")
	log.Println(*result)
}
