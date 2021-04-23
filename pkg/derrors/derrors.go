package derrors

import "errors"

var (
	UserNotFound  = errors.New("user not found")
	PasswordError = errors.New("password error")

	FileRequired = errors.New("file required")

	FileMustBeImage = errors.New("file Content-Type should be image/gif, image/jpeg, image/png")
)
