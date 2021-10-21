package goddb

type defaultService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return defaultService{repository: repository}
}

func (receiver defaultService) Put(request SaveValue) error {
	return receiver.repository.Put(request.Key, request.Value)
}

func (receiver defaultService) Retrieve(request RetrieveValue) (error, string) {
	return receiver.repository.Retrieve(request.Key)
}

func (receiver defaultService) Flush() {
	receiver.repository.Flush()
}
