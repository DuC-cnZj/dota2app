package derrors

import "errors"

var (
	UserNotFound  = errors.New("user not found")
	PasswordError = errors.New("password error")
)
