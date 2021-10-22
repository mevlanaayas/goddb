package goddb

type ReadWriter interface {
	Read() (error, []byte)
	Write([]byte) error
}
