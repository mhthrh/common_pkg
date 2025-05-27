package xErrors

import (
	"google.golang.org/grpc/codes"
	"net/http"
	"time"
)

func NewErrConnectionFailed(err error, Err *Error) *Error {
	return &Error{
		Code:       "104002",
		ErrorType:  Connection,
		Message:    "connection failed",
		Detail:     "connection failed",
		Internal:   Err,
		baseError:  err,
		GrpcStatus: codes.Unavailable,
		HttpStatus: http.StatusNotFound,
		Time:       time.Now(),
	}
}
