package base_response

import (
	"github.com/vfyuliawan/my-golang-sdk/core/api_constant"
	"github.com/vfyuliawan/my-golang-sdk/core/env"
)

type BaseResponse[T any] struct {
	Code         string `json:"code"`
	CodeSystem   string `json:"codeSystem"`
	Message      string `json:"message"`
	MessageError string `json:"messageError"`
	Result       T      `json:"result"`
}

func Success[T any](result T) BaseResponse[T] {
	return BaseResponse[T]{
		Code:       api_constant.GeneralSuccess.GetCode(),
		CodeSystem: env.GetEnvString("APP_NAME"),
		Result:     result,
	}
}

func SuccessWithMessage[T any](result T, message string) BaseResponse[T] {
	return BaseResponse[T]{
		Code:       api_constant.GeneralSuccess.GetCode(),
		CodeSystem: env.GetEnvString("APP_NAME"),
		Message:    message,
		Result:     result,
	}
}

func ErrorWithResult[T any](codeSystem, code, message string, result T) Error {
	return Error{
		Code:       code,
		CodeSystem: codeSystem,
		Message:    message,
		Result:     result,
	}
}

type Error struct {
	CodeSystem string
	Code       string
	Message    string
	HttpStatus int
	Result     any
}

func NewError(codeSystem string, code string, message string, httpStatus int) Error {
	return Error{
		CodeSystem: codeSystem,
		Code:       code,
		Message:    message,
		HttpStatus: httpStatus,
	}
}

func NewErrorWithResult[T any](codeSystem string, code string, message string, httpStatus int, result T) Error {
	return Error{
		CodeSystem: codeSystem,
		Code:       code,
		Message:    message,
		HttpStatus: httpStatus,
		Result:     result,
	}
}
