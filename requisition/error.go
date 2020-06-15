package requisition

import (
	"errors"
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
	if err == nil {
		err = errors.New("")
	}
	return &Error{
		error: err,
		code:  code,
		msg:   "msg not register",
		lang: LanguageZh,
	}
}