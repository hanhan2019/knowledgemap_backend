package model

import "errors"

var (
	ErrorNtFoundThirdpartToken   = errors.New("can not found third party token")
	ErrorMCodeNotValidate        = errors.New("gen mcode is not validate")
	ErrorSessionTokenNotValidate = errors.New("session token is not validate")
)
