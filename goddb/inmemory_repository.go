package goddb

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

func (receiver inMemoryRepository) Get() (error, map[string]string) {
	result := make(map[string]string)
	for key, value := range receiver.storage {
		result[key] = value
	}
	return nil, result
}

func (receiver *inMemoryRepository) Load(values map[string]string) error {
	for key, value := range values {
		receiver.storage[key] = value
	}
	return nil
}
