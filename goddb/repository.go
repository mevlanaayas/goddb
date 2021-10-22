package goddb

type Repository interface {
	Put(key, value string) error
	Retrieve(key string) (error, string)
	Flush() error
	Get() (error, map[string]string)
	Load(values map[string]string) error
}
