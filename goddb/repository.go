package goddb

type Repository interface {
	Save(key, value string) error
	Retrieve(key string) (error, string)
	Flush() error
}
