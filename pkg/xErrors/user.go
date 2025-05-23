package xErrors

import (
	"google.golang.org/grpc/codes"
	"net/http"
	"time"
)

func NewErrUsrExist(err error, Err *Error) *Error {
	return &Error{
		Code:       "100001",
		ErrorType:  User,
		Message:    "already user exist",
		Detail:     "already user exist",
		Internal:   Err,
		baseError:  err,
		HttpStatus: http.StatusConflict,
		GrpcStatus: codes.AlreadyExists,
		Time:       time.Now(),
	}
}
func NewErrUsrNotExist(err error, Err *Error) *Error {
	return &Error{
		Code:       "100002",
		ErrorType:  User,
		Message:    "user doesnt exist",
		Detail:     "user doesnt exist",
		Internal:   Err,
		baseError:  err,
		GrpcStatus: codes.NotFound,
		HttpStatus: http.StatusNotFound,
		Time:       time.Now(),
	}
}
