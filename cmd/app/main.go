package main

import (
	"bolapi/internal/app/cron"
	"bolapi/internal/app/server"
	"log"
)

func main() {
	err := cron.Start()

	if err != nil {
		log.Fatal(err)
		return
	}

	server.Start()
}
