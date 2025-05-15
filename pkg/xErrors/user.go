package xErrors

import (
	"net/http"
	"time"
)

func NewErrUsrExist(err error, Err *Error) *Error {
	return &Error{
		Code:       100001,
		ErrorType:  User,
		Message:    "already user exist",
		Detail:     "already user exist",
		internal:   Err,
		baseError:  err,
		httpStatus: http.StatusConflict,
		Time:       time.Now(),
	}
}
func NewErrUsrNotExist(err error, Err *Error) *Error {
	return &Error{
		Code:       100002,
		ErrorType:  User,
		Message:    "user doesnt exist",
		Detail:     "user doesnt exist",
		internal:   Err,
		baseError:  err,
		httpStatus: http.StatusNotFound,
		Time:       time.Now(),
	}
}
