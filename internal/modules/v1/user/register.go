package user

import "github.com/gin-gonic/gin"

// Register 注册
// @Summary 注册
// @Description 用户注册
// @Tags 用户
// @Produce  json
// @Param req body RegisterRequest true "请求参数"
// @Success 200 {object} model.UserInfo "用户信息"
// @Router /Register [post]
func Register(c *gin.Context) {

}

type RegisterRequest struct {
	Username        string `json:"username" form:"username"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type LoginEmailCredentialsRequest struct {
	Email    string `json:"email" form:"email" binding:"required" `
	Password string `json:"password" form:"password" binding:"required" `
}

type LoginUsernameCredentialsRequest struct {
	Username string `json:"username" form:"username" binding:"required" `
	Password string `json:"password" form:"password" binding:"required" `
}

type LoginPhoneCredentialsRequest struct {
	Phone      int64 `json:"phone" form:"phone" binding:"required" example:"13010002000"`
	VerifyCode int   `json:"verify_code" form:"verify_code" binding:"required" example:"120110"`
}
