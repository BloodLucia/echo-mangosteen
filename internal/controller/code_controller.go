package controller

type codeController struct{}

type CodeController interface{}

func NewCodeController() CodeController {
	return &codeController{}
}
