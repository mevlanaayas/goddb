package cmd

import (
	"fmt"
	"goddb/api"
	"goddb/config"
)

func Run() error {
	cfg := config.Config{}
	err := cfg.Init()
	if err != nil {
		return fmt.Errorf("error while parsing config \n\t%v", err)
	}

	server := api.NewApi(&cfg, api.Handler{})

	if err := server.Start(); err != nil {
		return fmt.Errorf("error while starting server \n\t%v", err)
	}
	return nil
}
