package e

import "errors"

var (
	OK                 = errors.New("ok")
	ErrUserExisted     = errors.New("user existed")
	ErrNameOrPassWrong = errors.New("username or password wrong")
	ErrTokenEmpty      = errors.New("token is empty")
	ErrTokenInvalid    = errors.New("token is invalid")
	ErrTokenTimeout    = errors.New("token is timeout")
)
