package foundation

type Error struct {
	Code int
	Msg string
}

func (this *Error) Error() string {
	return this.Msg
}

func NewError(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}
