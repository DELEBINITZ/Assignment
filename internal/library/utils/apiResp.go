package utils

import (
	"fmt"
	"net/http"

	"github.com/akshay/libraryAssign/internal/library/exceptions"
	"github.com/gin-gonic/gin"
)

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
		fmt.Println()
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

func RespondJSON(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, payload)
}

// RespondError makes the error response with payload as json format
func RespondError(c *gin.Context, err error, payload interface{}) {
	//send the error to track
	statusCode := http.StatusInternalServerError
	if serverErr, ok := err.(exceptions.ServerError); ok {
		statusCode = serverErr.HttpCode
	}
	RespondJSON(c, statusCode, payload)
}
