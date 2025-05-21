package xErrors

import (
	"fmt"
	"net/http"
	"time"
)

func NotImplemented(s string) *Error {
	return &Error{
		Code:       "10001",
		Message:    "method not implemented",
		ErrorType:  Loader,
		Detail:     fmt.Sprintf("%s not Implemented", s),
		internal:   nil,
		baseError:  nil,
		httpStatus: http.StatusNotFound,
		Time:       time.Now(),
	}
}
func FailedResource(err error, E *Error) *Error {
	return &Error{
		Code:       "10002",
		Message:    "resource failed",
		ErrorType:  Loader,
		Detail:     fmt.Sprintf("canot access to resource, %v ", err),
		internal:   E,
		baseError:  err,
		httpStatus: http.StatusNotFound,
		Time:       time.Now(),
	}
}
