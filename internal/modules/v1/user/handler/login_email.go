package handler

import "github.com/gin-gonic/gin"

type LoginEmailCredentialsRequest struct {
	Email    string `json:"email" form:"email" binding:"required" `
	Password string `json:"password" form:"password" binding:"required" `
}

func (handler *UserHandler) EmailLoginHandler(c *gin.Context) {

}
