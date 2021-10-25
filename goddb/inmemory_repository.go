package goddb

import (
	"fmt"
	"time"
)

type inMemoryRepository struct {
	storage map[string]string
}

func NewInMemoryRepository() GetPutFlusher {
	return &inMemoryRepository{
		storage: make(map[string]string),
	}
}

func (receiver inMemoryRepository) Put(key, value string) error {
	fmt.Printf("%v inMemoryRepository.Put is called \n", time.Now())
	receiver.storage[key] = value
	return nil
}

func (receiver inMemoryRepository) Get(key string) (error, string) {
	fmt.Printf("%v inMemoryRepository.Get is called \n", time.Now())
	return nil, receiver.storage[key]
}

func (receiver *inMemoryRepository) Flush() error {
	fmt.Printf("%v inMemoryRepository.Flush is called \n", time.Now())
	receiver.storage = make(map[string]string)
	return nil
}

func (receiver inMemoryRepository) GetAll() (error, map[string]string) {
	fmt.Printf("%v inMemoryRepository.GetAll is called \n", time.Now())
	result := make(map[string]string)
	for key, value := range receiver.storage {
		result[key] = value
	}
	return nil, result
}

func (receiver *inMemoryRepository) PutAll(values map[string]string) error {
	fmt.Printf("%v inMemoryRepository.PutlAll is called \n", time.Now())
	for key, value := range values {
		receiver.storage[key] = value
	}
	return nil
}
