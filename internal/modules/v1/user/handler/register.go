package handler

import (
	"Meow-backend/internal/models"
	"Meow-backend/internal/modules/v1/user"
	"Meow-backend/pkg/errcode"
	"Meow-backend/pkg/log"
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username        string `json:"username" form:"username"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

// RegisterHandler 注册
// @Summary 注册
// @Description 用户注册
// @Tags 用户
// @Produce  json
// @Param req body RegisterRequest true "请求参数"
// @Success 200 {object} model.UserInfo "用户信息"
// @Router /Register [post]
func RegisterHandler(c *gin.Context) {
	// get request params
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Warnf("register bind param err: %v", err)
		models.FailResult(c, errcode.ErrInvalidParam)
		return
	}

	log.Infof("register req: %#v", req)
	// validate request params
	if req.Username == "" || req.Email == "" || req.Password == "" || req.ConfirmPassword == "" {
		log.Warnf("params is empty: %v", req)
		models.FailResult(c, errcode.ErrInvalidParam)
		return
	}

	// check password and confirm password is same
	if req.Password != req.ConfirmPassword {
		log.Warnf("twice password is not same")
		models.FailWithMessage("twice password is not same", c)
		return
	}

	// 检查用户名是否已被注册（如果提供了用户名）
	if req.Username != "" {
		exists, err := service.CheckUsernameExists(req.Username)
		if err != nil {
			log.Errorf("error checking username existence: %v", err)
			models.FailResult(c, errcode.ErrInternalServer)
			return
		}
		if exists {
			models.FailWithMessage("Username already exists", c)
			return
		}
	}

	// 根据提供的信息进行注册
	var err error
	if req.Email != "" {
		err = service.RegisterWithEmail(c, req.Email, req.Password)
	} else {
		err = service.RegisterWithUsername(c, req.Username, req.Password)
	}

	if err != nil {
		log.Errorf("registration failed: %v", err)
		models.FailResult(c, errcode.ErrInternalServer)
		return
	}

	models.OkResult(c, "Registration successful")
	user.UserService.Register(c, req.Username, req.Email, req.Password)

}
