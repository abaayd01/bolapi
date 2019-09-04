package main

import (
	"bolapi/internal/app/cron"
	"bolapi/internal/app/server"
	"bolapi/internal/pkg/bolpy_client"
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
	// can group all this dependency setup into a cronWorker context
	cryptoCompareClient := crypto_compare.CryptoCompareClient{BaseURL: crypto_compare.APIStub.URL}

	bolpyCtx, _ := bolpy_client.NewBolpyContext("localhost:50051")
	bolpyClient := bolpy_client.BolpyClient{Context: bolpyCtx}

	cronWorker := cron.Worker{
		CryptoCompareClient: &cryptoCompareClient,
		BolpyClient:         &bolpyClient,
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
