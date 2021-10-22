package goddb

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type filePersistenceService struct {
	path string
}

func NewFilePersistenceService(path string) ReadWriter {
	return filePersistenceService{
		path: path,
	}
}

func (receiver filePersistenceService) Read() (error, []byte) {
	filename, err := filepath.Abs(fmt.Sprintf("%s/latest-data.json", receiver.path))
	if err != nil {

	}
	bytes, err := os.ReadFile(filename)
	if err != nil {

	}
	return nil, bytes
}

func (receiver filePersistenceService) Write(value []byte) error {
	filename, err := filepath.Abs(fmt.Sprintf("%s/%s-data.json", receiver.path, time.Now().Format("20060102150405")))
	if err != nil {

	}
	err = os.WriteFile(filename, value, 0644)
	if err != nil {

	}
	filename, err = filepath.Abs(fmt.Sprintf("%s/latest-data.json", receiver.path))
	if err != nil {

	}
	err = os.WriteFile(filename, value, 0644)
	if err != nil {

	}
	return nil
}
