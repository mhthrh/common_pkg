package xErrors

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"net/http"
	"time"
)

const (
	SuccessCode = "00"
	timeFormat  = "2006-01-02 15:04:05.000"
	User        = "invalidUser"
	Validation  = "validation"
	Loader      = "configLoader"
	Token       = "invalidToken"
	Convert     = "CastError"
	Successful  = "success"
	General     = "general"
)

type Error struct {
	Code       string `json:"code"`
	ErrorType  string `json:"-"`
	Message    string `json:"message"`
	Detail     string `json:"detail"`
	Internal   *Error `json:"-"`
	baseError  error
	HttpStatus int        `json:"-"`
	GrpcStatus codes.Code `json:"-"`
	Time       time.Time  `json:"time"`
}

func GetHttpStatus(e *Error, method string) int {

	switch {
	case e.ErrorType == Validation:
		return http.StatusBadRequest
	case e.ErrorType == Successful && method == "POST":
		return http.StatusCreated
	case e.ErrorType == Successful:
		return http.StatusOK
	default:
		if result := e.HttpStatus; result != 0 {
			return result
		}
		return http.StatusNotImplemented
	}

}
func GetGrpcCode(e *Error) codes.Code {
	if e == nil {
		return codes.OK
	}
	if e.ErrorType == Validation || e.ErrorType == Convert {
		return codes.FailedPrecondition
	}
	return e.GrpcStatus
}
func StringVerbal(e *Error) string {
	return fmt.Sprintf("error code:%s, error message %s, detail: %s, internal error: %v, base error: %v, time: %s", e.Code, e.Message, e.Detail, e.Internal, e.baseError, e.Time.Format(timeFormat))
}
func String(e *Error) string {
	return fmt.Sprintf("error code:%s, error message %s, detail: %s, time: %s", e.Code, e.Message, e.Detail, e.Time.Format(timeFormat))
}

func Success() *Error {
	return &Error{
		Code:       SuccessCode,
		Message:    "operation was success",
		ErrorType:  Successful,
		Detail:     "successful",
		Internal:   nil,
		baseError:  nil,
		GrpcStatus: codes.OK,
		Time:       time.Now(),
	}
}
func NewErrNotImplemented(s string) *Error {
	return &Error{
		Code:       "20000",
		Message:    "method/route not found/implemented",
		ErrorType:  General,
		Detail:     fmt.Sprintf("method: %s not found/implemented", s),
		Internal:   nil,
		baseError:  nil,
		HttpStatus: http.StatusNotFound,
		GrpcStatus: codes.NotFound,
		Time:       time.Now(),
	}
}
