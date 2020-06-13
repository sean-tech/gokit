package foundation

import (
	"errors"
	"fmt"
)

type IError interface {
	error
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
	if err == nil {
		err = errors.New(fmt.Sprintf("code:%d,msg:%s", code, msg))
	}
	return &Error{
		error: err,
		code:  code,
		msg:   msg,
	}
}
