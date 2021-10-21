package goddb

type Service interface {
	Put(request SaveValue) error
	Retrieve(request RetrieveValue) (error, string)
	Flush()
	Save()
}
