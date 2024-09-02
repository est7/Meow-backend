package interfaces

type Handler interface {
	GetService() Service
}

type BaseHandler struct {
	Service Service
}

func (h *BaseHandler) GetService() Service {
	return h.Service
}

func NewHandler(service Service) Handler {
	return &BaseHandler{Service: service}
}

// 不写到 Handle 中了
//func (h *BaseHandler) RegisterRoutes(*gin.Engine, func(auth.PermissionLevel) gin.HandlerFunc) {
//}
