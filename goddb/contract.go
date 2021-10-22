package goddb

import "fmt"

type Request interface {
	Validate() error
}

type (
	SaveValue struct {
		Key   string
		Value string
	}
	RetrieveValue struct {
		Key string
	}
)

func (receiver SaveValue) Validate() error {
	return &InternalError{
		m: "",
		c: 0,
		t: fmt.Errorf(""),
	}
}

func (receiver RetrieveValue) Validate() error {
	return nil
}
