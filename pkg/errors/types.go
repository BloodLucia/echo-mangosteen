package errors

import "net/http"

func InternalServer() *Error {
	return New(http.StatusInternalServerError, "服务器内部错误")
}

func IsInternalServer(err *Error) bool {
	return err.Code == http.StatusInternalServerError
}

func BadRequest() *Error {
	return New(http.StatusBadRequest, "请求格式错误, 查看data中的错误字段")
}

func Unauthorized() *Error {
	return New(http.StatusUnauthorized, "Unauthorized")
}

func NotFound() *Error {
	return New(http.StatusNotFound, "NotFound")
}

func InvalidRequestBody() *Error {
	return New(http.StatusUnprocessableEntity, "请求格式错误, 查看data中的错误字段")
}

func UnprocessableEntity() *Error {
	return New(http.StatusUnprocessableEntity, "无法解析请求体")
}

func IsUnprocessableEntity(err *Error) bool {
	return err.Code == http.StatusUnprocessableEntity
}
