package models

import (
	"Meow-backend/pkg/app"
	"github.com/gin-gonic/gin"
)

func OkResult(c *gin.Context) {
	app.SuccessResponse(c, nil)
}

func OkWithMessage(c *gin.Context, message string) {
	var messages = make(map[string]interface{})
	messages["message"] = message
	app.SuccessResponse(c, message)
}

func OkWithData(c *gin.Context, data interface{}) {
	app.SuccessResponse(c, data)
}

func OkWithDetailed(c *gin.Context, data interface{}, message string) {
	// messages is []string
	var messages []string
	messages = append(messages, message)
	app.SuccessResponseWithDetailed(c, data, messages)
}

func FailResult(c *gin.Context, err error) {
	app.ErrorResponse(c, err)
}

func FailWithData(c *gin.Context, data interface{}) {
	app.ErrorResponseWithData(c, data)
}

func FailWithMessage(c *gin.Context, message string) {
	app.ErrorResponseWithMessage(c, message)
}

func FailWithDetailed(c *gin.Context, data interface{}, message string) {
	var messages []string
	messages = append(messages, message)
	app.ErrorResponseWithDetailed(c, data, messages)
}
