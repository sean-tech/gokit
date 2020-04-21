package foundation

type Error struct {
	code int
	msg  string
}

func (this *Error) Error() string {
	return this.msg
}

func (this *Error) Code() int {
	return this.code
}

func (this *Error) Msg() string {
	return this.msg
}

func NewError(code int, msg string) *Error {
	return &Error{
		code: code,
		msg:  msg,
	}
}
