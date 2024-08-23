package models

import (
	"Meow-backend/pkg/app"
	"github.com/gin-gonic/gin"
)

func OkResult(c *gin.Context) {
	app.SuccessResponse(c, nil)
}

func OkWithMessage(message string, c *gin.Context) {
	var messages = make(map[string]interface{})
	messages["message"] = message
	app.SuccessResponse(c, message)
}

func OkWithData(data interface{}, c *gin.Context) {
	app.SuccessResponse(c, data)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	// messages is []string
	var messages []string
	messages = append(messages, message)
	app.SuccessResponseWithDetailed(c, data, messages)
}

func FailResult(c *gin.Context, err error) {
	app.ErrorResponse(c, err)
}

func FailWithData(data interface{}, c *gin.Context) {
	app.ErrorResponseWithData(c, data)
}

func FailWithMessage(message string, c *gin.Context) {
	app.ErrorResponseWithMessage(c, message)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	var messages []string
	messages = append(messages, message)
	app.ErrorResponseWithDetailed(c, data, messages)
}
