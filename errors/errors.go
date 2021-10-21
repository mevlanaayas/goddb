package errors

type Error interface {
	error
	Code() int
	Message() string
}

func New(message string, code int, trace error) Error {
	return &InternalError{
		m: message,
		c: code,
		t: trace,
	}
}

type InternalError struct {
	m string
	c int
	t error
}

func (e *InternalError) Error() string {
	return e.t.Error()
}

func (e *InternalError) Code() int {
	return e.c
}

func (e *InternalError) Message() string {
	return e.m
}
