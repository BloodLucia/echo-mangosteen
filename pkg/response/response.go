package response

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type respBody struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Build(ctx echo.Context, err error, data any) error {
	if err == nil {
		return ctx.JSON(http.StatusOK, respBody{
			Code: http.StatusOK,
			Msg:  "ok",
			Data: data,
		})
	}

	return ctx.JSON(errCodeMap[err], respBody{
		Code: errCodeMap[err],
		Msg:  err.Error(),
		Data: nil,
	})
}

var errCodeMap = map[error]int{}

func NewErr(code int, msg string) error {
	err := errors.New(msg)
	errCodeMap[err] = code
	return err
}
