package util

type MyError struct {
	Code int
	Msg  string
	Data []interface{}
}

func NewMyError(code int, msg string, data []interface{}) *MyError {
	return &MyError{Code: code, Msg: msg, Data: data}
}

func (e *MyError) Error() string {
	return e.Msg
}
