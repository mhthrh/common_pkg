package xErrors

import "time"

func NewErrMobilePhone(err error, Err *Error) *Error {
	return &Error{
		Code:      "101001",
		ErrorType: Validation,
		Message:   "mobile phone not valid",
		Detail:    "mobile phone not valid",
		Internal:  Err,
		baseError: err,
		Time:      time.Now(),
	}
}

func NewErrName(err error, Err *Error) *Error {
	return &Error{
		Code:      "101002",
		ErrorType: Validation,
		Message:   "name not valid",
		Detail:    "name not valid",
		Internal:  Err,
		baseError: err,
		Time:      time.Now(),
	}
}
func NewErrPasswordValidation(err error, Err *Error) *Error {
	return &Error{
		Code:      "101003",
		ErrorType: Validation,
		Message:   "password not valid",
		Detail:    "password not valid",
		Internal:  Err,
		baseError: err,
		Time:      time.Now(),
	}
}
func NewErrEmailValidation(err error, Err *Error) *Error {
	return &Error{
		Code:      "101004",
		ErrorType: Validation,
		Message:   "email not valid",
		Detail:    "email not valid",
		Internal:  Err,
		baseError: err,
		Time:      time.Now(),
	}
}
