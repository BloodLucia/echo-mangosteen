package errors

type Error struct {
	Code int
	Msg  string
	Err  error
}

func New(code int, msg string) *Error {
	return &Error{Code: code, Msg: msg}
}

// Error implements error
func (e *Error) Error() string {
	return e.Msg
}

// WithMsg with message
func (e *Error) WithMsg(msg string) *Error {
	e.Msg = msg
	return e
}

// WithErr with error
func (e *Error) WithErr(err error) *Error {
	e.Err = err
	return e
}
