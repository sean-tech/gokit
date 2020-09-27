package foundation

import (
	"fmt"
)

type Error interface {
	error
	Code() int
	Msg() string
}

type cerr struct {
	error
	code 	int
	msg  	string
}

func (this *cerr) Code() int {
	return this.code
}

func (this *cerr) Msg() string {
	return this.msg
}

func NewError(err error, code int, msg string) Error {
	if err == nil {
		err = fmt.Errorf("code:%d,msg:%s", code, msg)
	}
	return &cerr{
		error: err,
		code:  code,
		msg:  msg,
	}
}
