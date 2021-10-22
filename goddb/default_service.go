package goddb

import (
	"fmt"
)

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
	err := request.Validate()
	if err != nil {
		return NewError(fmt.Sprintf("error while validating retrieve request %v", err.Error()), 100400, err), ""
	}
	err, value := receiver.repository.Retrieve(request.Key)
	if err != nil {
		return NewError(fmt.Sprintf("error while retrieving value by key %s %v", request.Key, err.Error()), 100500, err), ""
	}
	return nil, value
}

func (receiver defaultService) Flush() error {
	err := receiver.repository.Flush()
	if err != nil {
		return NewError(fmt.Sprintf("error while flushing storage %v", err.Error()), 100500, err)
	}
	return nil
}

func (receiver defaultService) Save() error {
	/*
		jsonString, err := json.Marshal(receiver.)
		fmt.Println(err)

		return &InternalError{
			m: "",
			c: 0,
			t: fmt.Errorf("df"),
		}

	*/
	// return receiver.repository.Get()
	return nil
}

func (receiver defaultService) Load() error {
	return nil
	// return receiver.repository.Load()
}
