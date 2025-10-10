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
func NewErrClientReference(err error, Err *Error) *Error {
	return &Error{
		Code:      "101005",
		ErrorType: Validation,
		Message:   "clientVerificationReference not valid",
		Detail:    "A client submitted identifier.",
		Internal:  Err,
		baseError: err,
		Time:      time.Now(),
	}
}

func NewErrInitiatingPartyId(err error, Err *Error) *Error {
	return &Error{
		Code:      "101006",
		ErrorType: Validation,
		Message:   "initiatingPartyId not valid",
		Detail:    "The ID of the initiating party",
		Internal:  Err,
		baseError: err,
		Time:      time.Now(),
	}
}
func NewErrBank(err error, Err *Error) *Error {
	return &Error{
		Code:      "101007",
		ErrorType: Validation,
		Message:   "Bank not valid",
		Detail:    "Recipient's bank object",
		Internal:  Err,
		baseError: err,
		Time:      time.Now(),
	}
}

func NewErrType(err error, Err *Error) *Error {
	return &Error{
		Code:      "101008",
		ErrorType: Validation,
		Message:   "Type not valid",
		Detail:    "Identifies the type of recipient - \"I\" for Individual and \"C\" for Company.",
		Internal:  Err,
		baseError: err,
		Time:      time.Now(),
	}
}
