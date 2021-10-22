package goddb

type PersistenceService interface {
	Read() (error, []byte)
	Write([]byte) error
}
