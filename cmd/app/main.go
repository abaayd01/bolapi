package main

import (
	"bolapi/internal/app/cron"
	"bolapi/internal/app/server"
	"bolapi/internal/pkg/crypto_compare"
	"bolapi/internal/pkg/database"
	"fmt"
	"log"
)

func init() {
	err := database.InitDB()

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	cryptoCompareClient := crypto_compare.CryptoCompareClient{BaseURL: crypto_compare.APIStub.URL}
	cronWorker := cron.Worker{
		CryptoCompareClient: &cryptoCompareClient,
		DB:                  database.DBStub,
	}

	err := cronWorker.Start()

	if err != nil {
		log.Fatal(err)
		return
	}

	bolAPIServer := server.BolAPIServer{
		DB: database.DBStub,
	}

	bolAPIServer.Start()
}
