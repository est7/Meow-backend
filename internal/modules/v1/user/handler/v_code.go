package handler

import (
	"Meow-backend/internal/models"
	"github.com/gin-gonic/gin"
)

type VCodeModel struct {
	Code string `json:"code"`
}

func (handler *UserHandler) VCode(c *gin.Context) {
	//mock send code 123456
	response := VCodeModel{
		Code: "123456",
	}
	models.OkWithData(c, response)
}
