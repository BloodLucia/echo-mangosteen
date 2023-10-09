package e

import "echo-mangosteen/pkg/response"

var (
	ErrInternalServer = response.NewErr(500, "服务器内部错误")

	ErrDatabase = response.NewErr(501, "服务器内部错误")

	ErrInvalidRequestBody = response.NewErr(400, "bad request")

	ErrInvalidRequestParams = response.NewErr(422, "invalid request params")

	ErrInvalidEmail = response.NewErr(422, "invalid email address")
)
