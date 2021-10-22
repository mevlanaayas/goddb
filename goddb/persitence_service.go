package goddb

type PersistenceService interface {
	Read() (error, string)
	Write(string) error
}
