package cron

import (
	"github.com/robfig/cron"
	"log"
)

func Start() error {
	c := cron.New()
	err := c.AddFunc("@every 10s", func() { log.Println("every ten seconds") })

	c.Start()

	if err != nil {
		return err
	}

	return nil
}
