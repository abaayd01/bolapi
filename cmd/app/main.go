package main

import (
	"bolapi/internal/app/cron"
	"bolapi/internal/app/server"
	"bolapi/internal/pkg/crypto_compare"
	"log"
)

func main() {
	cryptoCompareClient := crypto_compare.CryptoCompareClient{BaseURL: crypto_compare.APIStub.URL}

	cronWorker := cron.CronWorker{
		CryptoCompareClient: &cryptoCompareClient,
	}

	err := cronWorker.Start()

	if err != nil {
		log.Fatal(err)
		return
	}

	server.Start()
}
