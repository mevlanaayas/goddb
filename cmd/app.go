package cmd

import (
	"fmt"
	"goddb/api"
	"goddb/config"
	"goddb/goddb"
)

func Run() error {
	cfg := config.Config{}
	err := cfg.Init()
	if err != nil {
		return fmt.Errorf("error while parsing config \n\t%v", err)
	}

	server := api.NewApi(&cfg, goddb.Handler{})

	if err := server.Start(); err != nil {
		return fmt.Errorf("error while starting server \n\t%v", err)
	}
	return nil
}
