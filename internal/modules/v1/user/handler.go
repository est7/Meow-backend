package user

type Handler struct {
	service *service.Service
}

func NewHandler() *Handler {
	return &Handler{
		service: service.NewService(),
	}
}
