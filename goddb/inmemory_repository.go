package goddb

type inMemoryRepository struct {
	storage map[string]string
}

func NewInMemoryRepository() Repository {
	return inMemoryRepository{
		storage: make(map[string]string),
	}
}

func (receiver inMemoryRepository) Initialize() {
	receiver.Flush()
}

func (receiver inMemoryRepository) Put(key, value string) error {
	receiver.storage[key] = value
	return nil
}

func (receiver inMemoryRepository) Retrieve(key string) (error, string) {
	return nil, receiver.storage[key]
}

func (receiver inMemoryRepository) Flush() {
	receiver.storage = make(map[string]string)
}
