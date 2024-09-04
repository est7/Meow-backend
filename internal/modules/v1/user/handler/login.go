package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginPhoneCredentialsRequest struct {
	Phone      int64 `json:"phone" form:"phone" binding:"required" example:"13010002000"`
	VerifyCode int   `json:"verify_code" form:"verify_code" binding:"required" example:"120110"`
}

func (handler *UserHandler) PhoneLoginHandler(c *gin.Context) {

}

type LoginUsernameCredentialsRequest struct {
	Username string `json:"username" form:"username" binding:"required" `
	Password string `json:"password" form:"password" binding:"required" `
}

func (handler *UserHandler) UsernameLoginHandler(c *gin.Context) {
	// get request params
	var req LoginUsernameCredentialsRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// validate request params
	if req.Username == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

}

type LoginEmailCredentialsRequest struct {
	Email    string `json:"email" form:"email" binding:"required" `
	Password string `json:"password" form:"password" binding:"required" `
}

func (handler *UserHandler) EmailLoginHandler(c *gin.Context) {

}
