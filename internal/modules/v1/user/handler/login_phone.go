package handler

import "github.com/gin-gonic/gin"

type LoginPhoneCredentialsRequest struct {
	Phone      int64 `json:"phone" form:"phone" binding:"required" example:"13010002000"`
	VerifyCode int   `json:"verify_code" form:"verify_code" binding:"required" example:"120110"`
}

func PhoneLoginHandler(c *gin.Context) {

}
