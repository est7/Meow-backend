package interfaces

type Handler interface {
	GetService() BaseService
}

type BaseHandler struct {
	Service BaseService
}

func (h *BaseHandler) GetService() BaseService {
	return h.Service
}

func NewHandler(service BaseService) Handler {
	return &BaseHandler{Service: service}
}

// 不写到 Handle 中了
//func (h *BaseHandler) RegisterRoutes(*gin.Engine, func(auth.PermissionLevel) gin.HandlerFunc) {
//}
