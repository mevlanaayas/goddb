package goddb

import "fmt"

type inMemoryRepository struct {
	storage map[string]string
}

func NewInMemoryRepository() Repository {
	return &inMemoryRepository{
		storage: make(map[string]string),
	}
}

func (receiver inMemoryRepository) Put(key, value string) error {
	receiver.storage[key] = value
	return nil
}

func (receiver inMemoryRepository) Retrieve(key string) (error, string) {
	return nil, receiver.storage[key]
}

func (receiver *inMemoryRepository) Flush() error {
	receiver.storage = make(map[string]string)
	return nil
}

func (receiver inMemoryRepository) Save() error {
	fmt.Printf("saving...")
	// TODO: save all key:values to json file
	return nil
}

func (receiver inMemoryRepository) Load() error {
	fmt.Printf("loading...")
	// TODO: get all key:values from json file
	return nil
}
