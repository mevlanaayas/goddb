package goddb

type Service interface {
	Save(request SaveValue) error
	Retrieve(request RetrieveValue) (error, string)
	Flush() error
}
