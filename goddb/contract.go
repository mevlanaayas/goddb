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
	if len(receiver.Key) < 1 {
		return &InternalError{
			m: "key length is not valid",
			c: 100100,
			t: fmt.Errorf("key length is not valid"),
		}
	}
	if len(receiver.Value) < 1 {
		return &InternalError{
			m: "value length is not valid",
			c: 100101,
			t: fmt.Errorf("value length is not valid"),
		}
	}
	return nil
}

func (receiver RetrieveValue) Validate() error {
	if len(receiver.Key) < 1 {
		return &InternalError{
			m: "key length is not valid",
			c: 100100,
			t: fmt.Errorf("key length is not valid"),
		}
	}
	return nil
}
