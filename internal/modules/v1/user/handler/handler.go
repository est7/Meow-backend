package handler

import (
	"Meow-backend/internal/interfaces"
	service "Meow-backend/internal/modules/v1/user/service"
)

type UserHandler struct {
	userService *service.UserServiceImpl
}

func NewUserHandler(base interfaces.Handler) *UserHandler {
	return &UserHandler{
		userService: service.NewUserService(base.GetService()),
	}
}
