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
	err := c.AddFunc("@every 10s", func() { takeSnapshot(cW.CryptoCompareClient) })

	c.Start()

	if err != nil {
		return err
	}

	return nil
}

func takeSnapshot(client *crypto_compare.CryptoCompareClient) {
	log.Println("taking snapshot")
	result, _ := client.GetCurrentPrice("ETH", "USD")
	log.Println(*result)
}
