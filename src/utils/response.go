package utils

import (
	"github.com/gin-gonic/gin"
)

type Responses struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ErrorResponse struct {
	StatusCode int         `json:"statusCode"`
	Error      interface{} `json:"error"`
}

func WriteSuccessResponse(ctx *gin.Context, Message string, StatusCode int, Data interface{}) {

	jsonResponse := Responses{
		StatusCode: StatusCode,
		Message:    Message,
		Data:       Data,
	}

	if StatusCode >= 400 {
		ctx.JSON(StatusCode, jsonResponse)
		defer ctx.AbortWithStatus(StatusCode)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}

func WriteErrorResponse(ctx *gin.Context, StatusCode int, Error interface{}) {
	errResponse := ErrorResponse{
		StatusCode: StatusCode,
		Error:      Error,
	}

	ctx.JSON(StatusCode, errResponse)
	defer ctx.AbortWithStatus(StatusCode)
}
