package goddb

type Repository interface {
	Put(key, value string) error
	Retrieve(key string) (error, string)
	Flush() error
	Save() error
	Load() error
}
