package user

type Handler struct {
	service *service.Service
}

func NewHandler() *Handler {
	return &Handler{
		service: service.NewService(),
	}
}

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}
