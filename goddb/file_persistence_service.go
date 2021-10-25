package goddb

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
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
	fmt.Printf("%v filePersistenceService.Read is called \n", time.Now())
	filename, err := filepath.Abs(fmt.Sprintf("%s/latest-data.json", receiver.path))
	if err != nil {
		return NewError(fmt.Sprintf("error while generating filename %v\n\t", err.Error()), 100500, err), nil
	}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, []byte("{}")
		}
		return NewError(fmt.Sprintf("error while reading file %v\n\t", err.Error()), 100500, err), nil
	}
	return nil, bytes
}

func (receiver filePersistenceService) Write(value []byte) error {
	fmt.Printf("%v filePersistenceService.Write is called \n", time.Now())
	filename, err := filepath.Abs(fmt.Sprintf("%s/%s-data.json", receiver.path, time.Now().Format("20060102150405")))
	if err != nil {
		return NewError(fmt.Sprintf("error while generating filename with ts %v\n\t", err.Error()), 100500, err)
	}
	err = ioutil.WriteFile(filename, value, 0644)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Printf("File does not exist \n\t%v", err)
			return nil
		}
		return NewError(fmt.Sprintf("error while writing to file %v\n\t", err.Error()), 100500, err)
	}
	filename, err = filepath.Abs(fmt.Sprintf("%s/latest-data.json", receiver.path))
	if err != nil {
		return NewError(fmt.Sprintf("error while generating latest file filename %v\n\t", err.Error()), 100500, err)
	}
	err = ioutil.WriteFile(filename, value, 0644)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Printf("File does not exist \n\t%v", err)
			return nil
		}
		return NewError(fmt.Sprintf("error while writing to file %v\n\t", err.Error()), 100500, err)
	}
	return nil
}
