package goddb

type Service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return Service{repository: repository}
}

func (receiver Service) Save(request SaveValue) error {
	return receiver.repository.Save(request.Key, request.Value)
}

func (receiver Service) Retrieve(request RetrieveValue) (error, string) {
	return receiver.repository.Retrieve(request.Key)
}

func (receiver Service) Flush() error {
	return receiver.repository.Flush()
}
