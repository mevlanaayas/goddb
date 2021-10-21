package goddb

type defaultService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return defaultService{repository: repository}
}

func (receiver defaultService) Save(request SaveValue) error {
	return receiver.repository.Save(request.Key, request.Value)
}

func (receiver defaultService) Retrieve(request RetrieveValue) (error, string) {
	return receiver.repository.Retrieve(request.Key)
}

func (receiver defaultService) Flush() error {
	return receiver.repository.Flush()
}
