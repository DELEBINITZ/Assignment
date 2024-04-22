package api

import "github.com/akshay/libraryAssign/internal/library/exceptions"

type ApiResp struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type ErrorResp struct {
	ErrorCode exceptions.ServerErrorCode `json:"error_code,omitempty"`
	Msg       string                     `json:"message,omitempty"`
	Reason    string                     `json:"reason,omitempty"`
}

func ConstructAPIResp(message string, data interface{}) ApiResp {
	if data == nil {
		return ApiResp{Status: message}
	}
	return ApiResp{
		Status: message,
		Data:   data,
	}
}

func ConstructErrorAPIResp(err error) ApiResp {
	var payload ApiResp
	if serverErr, ok := err.(exceptions.ServerError); ok {
		payload = ApiResp{
			Status: "error",
			Data: ErrorResp{
				Msg:       serverErr.Msg,
				Reason:    serverErr.ActualErr.Error(),
				ErrorCode: serverErr.Code,
			},
		}
	} else {
		payload = ApiResp{
			Status: "error",
			Data: ErrorResp{
				Msg: err.Error(),
			},
		}
	}
	return payload
}
