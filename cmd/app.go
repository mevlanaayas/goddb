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
		return fmt.Errorf("error while initializing config \n\t%v", err)
	}

	inMemoryRepository := goddb.NewInMemoryRepository()
	defaultPersistenceService := goddb.NewFilePersistenceService(cfg.App.Path)
	service := goddb.NewStorageService(inMemoryRepository, defaultPersistenceService)

	ticker := NewTicker(service, cfg.App.SyncInMin)
	ticker.Schedule()

	handler := api.NewHandler(service)

	server := api.NewApi(cfg, handler)
	if err := server.Start(); err != nil {
		return fmt.Errorf("error while starting server \n\t%v", err)
	}
	return nil
}
