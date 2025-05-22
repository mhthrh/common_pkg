package xErrors

import (
	"fmt"
	"net/http"
	"time"
)

func NewErrConvertData(s error) *Error {
	return &Error{
		Code:       "10301",
		Message:    "cannot cast Body to struct",
		ErrorType:  Convert,
		Detail:     fmt.Sprintf("cannot cast Body to struct: => %s", s.Error()),
		Internal:   nil,
		baseError:  s,
		HttpStatus: http.StatusBadRequest,
		Time:       time.Now(),
	}
}
func NewErrKeyNotExist(s string) *Error {
	return &Error{
		Code:       "10302",
		ErrorType:  Convert,
		Message:    "key not exist in GET method",
		Detail:     fmt.Sprintf("cannot find %s in request", s),
		Internal:   nil,
		baseError:  nil,
		HttpStatus: http.StatusNotFound,
		Time:       time.Now(),
	}
}
