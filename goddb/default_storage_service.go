package goddb

import (
	"encoding/json"
	"fmt"
)

type defaultStorageService struct {
	repository         Repository
	persistenceService PersistenceService
}

func NewDefaultStorageService(repository Repository, persistenceService PersistenceService) StorageService {
	service := defaultStorageService{
		repository:         repository,
		persistenceService: persistenceService,
	}
	_ = service.Load()
	return service
}

func (receiver defaultStorageService) Put(request SaveValue) error {
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

func (receiver defaultStorageService) Retrieve(request RetrieveValue) (error, string) {
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

func (receiver defaultStorageService) Flush() error {
	err := receiver.repository.Flush()
	if err != nil {
		return NewError(fmt.Sprintf("error while flushing storage %v", err.Error()), 100500, err)
	}
	return nil
}

func (receiver defaultStorageService) Save() error {
	err, values := receiver.repository.Get()
	if err != nil {
		return NewError(fmt.Sprintf("error while getting values value by key %v", err.Error()), 100500, err)
	}
	jsonString, err := json.Marshal(values)
	if err != nil {
		return NewError(fmt.Sprintf("error while converting values into json string %v", err.Error()), 100500, err)
	}
	fmt.Println(jsonString)
	err = receiver.persistenceService.Write(jsonString)
	if err != nil {
		return NewError(fmt.Sprintf("error while persisting storage %v", err.Error()), 100500, err)
	}
	return nil
}

func (receiver defaultStorageService) Load() error {
	err, jsonString := receiver.persistenceService.Read()
	var values map[string]string

	err = json.Unmarshal(jsonString, &values)
	if err != nil {
		return NewError(fmt.Sprintf("error while converting json string into key:value map %v", err.Error()), 100500, err)
	}
	err = receiver.repository.Load(values)
	if err != nil {
		return NewError(fmt.Sprintf("error while loading values into storage %v", err.Error()), 100500, err)
	}
	return nil
}
