package goddb

type defaultPersistenceService struct {
	config string
}

func NewDefaultPersistenceService(config string) PersistenceService {
	return defaultPersistenceService{
		config: config,
	}
}

func (receiver defaultPersistenceService) Read() (error, string) {
	return nil, ""
}

func (receiver defaultPersistenceService) Write(value string) error {
	return nil
}
