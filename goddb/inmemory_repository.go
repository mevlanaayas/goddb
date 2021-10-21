package goddb

import "goddb/errors"

type inMemoryRepository struct {
}

func NewInMemoryRepository() Repository {
	return nil
}

func (receiver inMemoryRepository) Save(key, value string) error {
	return &errors.InternalError{}
}

func (receiver inMemoryRepository) Retrieve(key string) (error, string) {
	return &errors.InternalError{}, ""
}

func (receiver inMemoryRepository) Flush() error {
	return &errors.InternalError{}
}
