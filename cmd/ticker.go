package cmd

import (
	"fmt"
	"goddb/goddb"
	"time"
)

type Ticker struct {
	service     goddb.Service
	ticker      *time.Ticker
	syncChannel chan bool
}

func NewTicker(service goddb.Service) Ticker {
	return Ticker{
		service:     service,
		ticker:      time.NewTicker(10 * time.Minute),
		syncChannel: make(chan bool),
	}
}

func (receiver Ticker) Schedule() {
	go func() {
		for {
			select {
			case <-receiver.syncChannel:
				return
			case t := <-receiver.ticker.C:
				fmt.Println("Tick at", t)
				err := receiver.service.Save()
				if err != nil {
					fmt.Println("error while saving current state", t)
				}
			}
		}
	}()
}
