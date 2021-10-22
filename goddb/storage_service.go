package goddb

type StorageService interface {
	Put(request SaveValue) error
	Retrieve(request RetrieveValue) (error, string)
	Flush() error
	Save() error
	Load() error
}
