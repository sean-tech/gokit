package foundation

type IError interface {
	Code() int
	Msg() string
}

type Error struct {
	error
	code 	int
	msg  	string
}

func (this *Error) Code() int {
	return this.code
}

func (this *Error) Msg() string {
	return this.msg
}

func NewError(err error, code int, msg string) *Error {
	return &Error{
		error: err,
		code:  code,
		msg:   msg,
	}
}
