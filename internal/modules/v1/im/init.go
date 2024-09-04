package im

import (
	"Meow-backend/internal/interfaces"
	factory "Meow-backend/internal/interfaces/servicefactory"
	"Meow-backend/internal/modules/v1/user/handler"
	"Meow-backend/pkg/auth"
	"Meow-backend/pkg/log"
	"github.com/gin-gonic/gin"
)

type IMModule struct {
	appCtx  interfaces.AppContext
	handler *handler.IMHander
}

func NewIMModule(ctx interfaces.AppContext) interfaces.Module {
	repo := interfaces.NewRepository(ctx.GetGormDB())
	userServiceFactory := factory.NewUserServiceFactory(interfaces.NewServiceFactory())
	userService := userServiceFactory.CreateService(repo, ctx.GetRedisClient())

	serviceFactory := factory.NewIMServiceFactory(interfaces.NewServiceFactory())
	imHandler := handler.NewUserHandler(interfaces.NewHandler(service))
	return &IMModule{appCtx: ctx,
		handler: imHander,
	}
}

func (u *IMModule) Name() string {
	return "IM"
}

func (u *IMModule) Init(appCtx interfaces.AppContext) {
	u.appCtx = appCtx
	log.Info("Initializing im module")
}

func (u *IMModule) RegisterRoutes(r *gin.Engine, authMiddleware func(auth.PermissionLevel) gin.HandlerFunc) {
}
