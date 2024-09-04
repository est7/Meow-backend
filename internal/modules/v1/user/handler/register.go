package handler

import (
	"Meow-backend/internal/models"
	"Meow-backend/pkg/errcode"
	"Meow-backend/pkg/log"
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username        string `json:"username" form:"username"`
	Email           string `json:"email" form:"email"`
	EmailVCode      string `json:"email_vcode" form:"email_vcode"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
	Phone           string `json:"phone" form:"phone"`
	PhoneVCode      string `json:"phone_vcode" form:"phone_vcode"`
}

// RegisterWithUsernameHandler 注册
func (handler *UserHandler) RegisterWithUsernameHandler(c *gin.Context) {
	var req struct {
		Username        string `json:"username" binding:"required"`
		Password        string `json:"password" binding:"required"`
		ConfirmPassword string `json:"confirm_password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Warnf("register with username bind param err: %v", err)
		models.FailResult(c, errcode.ErrInvalidParam)
		return
	}

	if req.Password != req.ConfirmPassword {
		models.FailWithMessage(c, "Passwords do not match")
		return
	}

	err := handler.userService.RegisterWithUsername(c, req.Username, req.Password)
	if err != nil {
		log.Errorf("registration with username failed: %v", err)
		models.FailResult(c, errcode.ErrCustomError)
		return
	}

	models.OkWithMessage(c, "Registration successful")
}

func (handler *UserHandler) RegisterWithEmailHandler(c *gin.Context) {
	var req struct {
		Email      string `json:"email" binding:"required,email"`
		EmailVCode string `json:"email_vcode" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Warnf("register with email bind param err: %v", err)
		models.FailResult(c, errcode.ErrInvalidParam)
		return
	}

	err := handler.userService.RegisterWithEmail(c, req.Email, req.EmailVCode)
	if err != nil {
		log.Errorf("registration with email failed: %v", err)
		models.FailResult(c, errcode.ErrInternalServer)
		return
	}

	models.OkWithMessage(c, "Registration successful")
}

func (handler *UserHandler) RegisterWithPhoneHandler(c *gin.Context) {
	var req struct {
		Phone      string `json:"phone" binding:"required"`
		PhoneVCode string `json:"phone_vcode" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Warnf("register with phone bind param err: %v", err)
		models.FailResult(c, errcode.ErrInvalidParam)
		return
	}

	err := handler.userService.RegisterWithPhone(c, req.Phone, req.PhoneVCode)
	if err != nil {
		log.Errorf("registration with phone failed: %v", err)
		models.FailResult(c, errcode.ErrInternalServer)
		return
	}

	models.OkWithMessage(c, "Registration successful")
}
