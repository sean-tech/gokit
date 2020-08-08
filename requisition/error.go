package requisition

import (
	"fmt"
)

type IError interface {
	error
	Code() int
	Msg() string
	SetLang(lang string)
	GetLang() string
}

type Error struct {
	error
	code 	int
	msg  	string
	lang 	string
}

func (this *Error) Code() int {
	return this.code
}

func (this *Error) Msg() string {
	msg := Msg(this.lang, this.code)
	if msg != "" {
		this.msg = msg
	}
	return this.msg
}


func (this *Error) SetLang(lang string) {
	this.lang = lang
}


func (this *Error) GetLang() string {
	return this.lang
}

func NewError(err error, code int) *Error {
	msg := Msg(LanguageZh, code)
	if msg == "" {
		msg = Msg(LanguageEn, code)
	}
	if msg == "" {
		msg = "msg not register"
	}
	if err == nil {
		err = fmt.Errorf("code:%d,msg:%s", code, msg)
	}
	return &Error{
		error: err,
		code:  code,
		msg:  msg,
		lang: LanguageZh,
	}
}