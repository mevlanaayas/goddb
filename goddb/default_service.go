package goddb

import "fmt"

type defaultService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return defaultService{repository: repository}
}

func (receiver defaultService) Put(request SaveValue) error {
	err := request.Validate()
	if err != nil {
		return NewError(fmt.Sprintf("error while validating save request %v", err.Error()), 100400, err)
	}
	err = receiver.repository.Put(request.Key, request.Value)
	if err != nil {
		return NewError(fmt.Sprintf("error while saving key:value %s:%s %v", request.Key, request.Value, err.Error()), 100500, err)
	}
	return nil
}

func (receiver defaultService) Retrieve(request RetrieveValue) (error, string) {
	return receiver.repository.Retrieve(request.Key)
}

func (receiver defaultService) Flush() error {
	return receiver.repository.Flush()
}

func (receiver defaultService) Save() error {
	return receiver.repository.Save()
}

func (receiver defaultService) Load() error {
	return receiver.repository.Load()
}
