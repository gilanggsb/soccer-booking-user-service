package constants

import "errors"

var (
	ErrUserNotFound                   = errors.New("user not found")                      // 1 usage
	ErrPasswordIncorrect              = errors.New("password incorrect")                  // 1 usage
	ErrUsernameExist                  = errors.New("username already exist")              // 1 usage
	ErrEmailExist                     = errors.New("email already exist")                 // 1 usage
	ErrConfirmPasswordShouldntBeEmpty = errors.New("confirm password shouldn't be empty") // 1 usage
	ErrPasswordDoesNotMatch           = errors.New("password does not match")             // 1 usage
)

var UserErrors = []error{ // no usages
	ErrUserNotFound,
	ErrPasswordIncorrect,
	ErrUsernameExist,
	ErrPasswordDoesNotMatch,
	ErrConfirmPasswordShouldntBeEmpty,
}
