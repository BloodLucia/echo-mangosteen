package errors

import "net/http"

func InternalServer() *Error {
	return New(http.StatusInternalServerError, "服务器内部错误")
}

func IsInternalServer(err *Error) bool {
	return err.Code == 500
}

func BadRequest() *Error {
	return New(http.StatusBadRequest, "请求格式错误, 查看data中的错误字段")
}

func Unauthorized() *Error {
	return New(http.StatusUnauthorized, "Unauthorized")
}

func InvalidRequestBody() *Error {
	return New(http.StatusUnprocessableEntity, "请求格式错误, 查看data中的错误字段")
}
