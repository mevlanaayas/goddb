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
		ticker:      time.NewTicker(15 * time.Second),
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
				fmt.Printf("%v scheduler calling handler method \n", t)
				err := receiver.service.Save()
				if err != nil {
					fmt.Printf("%v error while saving current state %v\n\t", t, err)
				}
			}
		}
	}()
}
