package goddb

import (
	"fmt"
	"os"
	"time"
)

type defaultPersistenceService struct {
	config string
}

func NewDefaultPersistenceService(config string) ReadWriter {
	return defaultPersistenceService{
		config: config,
	}
}

func (receiver defaultPersistenceService) Read() (error, []byte) {
	filename := "/tmp/goddb/latest-data.json"
	bytes, err := os.ReadFile(filename)
	return err, bytes
}

func (receiver defaultPersistenceService) Write(value []byte) error {
	filename := fmt.Sprintf("/tmp/goddb/%s-data.json", time.Now().Format("20060102150405"))
	_ = os.WriteFile(filename, value, 0644)
	filename = "/tmp/goddb/latest-data.json"
	_ = os.WriteFile(filename, value, 0644)
	return nil
}
