package response

import (
	"errors"
	"net/http"

	myErrors "echo-mangosteen/pkg/errors"

	"github.com/labstack/echo/v4"
)

type respBody struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func buildErrResp(err *myErrors.Error) *respBody {
	if err == nil {
		return &respBody{
			Code: http.StatusInternalServerError,
			Msg:  "Unknown Error",
			Data: nil,
		}
	}
	return &respBody{
		Code: err.Code,
		Msg:  err.Msg,
		Data: nil,
	}
}

func Build(ctx echo.Context, err error, data any) error {
	if err == nil {
		return ctx.JSON(http.StatusOK, respBody{
			Code: http.StatusOK,
			Msg:  "ok",
			Data: data,
		})
	}

	var myErr *myErrors.Error

	// unknow error 未知错误
	if !errors.As(err, &myErr) {
		return ctx.JSON(http.StatusInternalServerError, buildErrResp(nil))
	}

	// internal server error 服务器内部错误
	if myErrors.IsInternalServer(myErr) {
		return ctx.JSON(http.StatusInternalServerError, buildErrResp(myErr))
	}

	resp := respBody{
		Code: myErr.Code,
		Msg:  myErr.Msg,
		Data: nil,
	}

	if data != nil {
		resp.Data = data
	}

	return ctx.JSON(myErr.Code, resp)
}
