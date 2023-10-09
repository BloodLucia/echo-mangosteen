package errors

import "net/http"

func InternalServer(reason string) *Error {
	return New(http.StatusInternalServerError, reason)
}

func IsInternalServer(err *Error) bool {
	return err.Code == 500
}
