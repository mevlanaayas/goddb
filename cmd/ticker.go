package cmd

import (
	"fmt"
	"goddb/goddb"
	"time"
)

type Ticker struct {
	service     goddb.StorageService
	ticker      *time.Ticker
	syncChannel chan bool
}

func NewTicker(service goddb.StorageService, duration int) Ticker {
	return Ticker{
		service:     service,
		ticker:      time.NewTicker(time.Duration(duration) * time.Minute),
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
				fmt.Printf("%v scheduler called handler method \n", t)
			}
		}
	}()
}
