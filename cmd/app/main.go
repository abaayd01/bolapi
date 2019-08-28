package main

import (
	"bolapi/internal/app/cron"
	"bolapi/internal/app/server"
	"bolapi/internal/pkg/crypto_compare"
	"log"
)

func main() {
	cryptoCompareApiStub := crypto_compare.APIStub

	client := crypto_compare.CryptoCompareClient{BaseURL: cryptoCompareApiStub.URL}

	cronWorker := cron.CronWorker{
		CryptoCompareClient: &client,
	}

	err := cronWorker.Start()

	if err != nil {
		log.Fatal(err)
		return
	}

	server.Start()
}
