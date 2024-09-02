package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
