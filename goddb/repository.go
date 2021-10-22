package goddb

type GetPutFlusher interface {
	Get(key string) (error, string)
	GetAll() (error, map[string]string)
	Put(key, value string) error
	PutAll(values map[string]string) error
	Flush() error
}
