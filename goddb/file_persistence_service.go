package goddb

import (
	"fmt"
	"os"
	"time"
)

type filePersistenceService struct {
	config string
}

func NewFilePersistenceService(config string) ReadWriter {
	return filePersistenceService{
		config: config,
	}
}

func (receiver filePersistenceService) Read() (error, []byte) {
	filename := "/tmp/goddb/latest-data.json"
	bytes, err := os.ReadFile(filename)
	return err, bytes
}

func (receiver filePersistenceService) Write(value []byte) error {
	filename := fmt.Sprintf("/tmp/goddb/%s-data.json", time.Now().Format("20060102150405"))
	_ = os.WriteFile(filename, value, 0644)
	filename = "/tmp/goddb/latest-data.json"
	_ = os.WriteFile(filename, value, 0644)
	return nil
}
