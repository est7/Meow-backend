package im

import (
	"Meow-backend/internal/interfaces"
	factory "Meow-backend/internal/interfaces/factoryforservice"
	"Meow-backend/internal/modules/v1/im/handler"
	"Meow-backend/pkg/auth"
	"Meow-backend/pkg/log"
	"github.com/gin-gonic/gin"
)

type IMModule struct {
	appCtx  interfaces.AppContext
	handler *handler.IMHandler
}

func NewIMModule(ctx interfaces.AppContext) interfaces.Module {
	repo := interfaces.NewRepository(ctx.GetGormDB())
	serviceFactory := factory.NewIMServiceFactory()
	imService := serviceFactory.CreateService(repo, ctx.GetRedisClient())
	imHandler := handler.NewIMHandler(imService)
	return &IMModule{
		appCtx:  ctx,
		handler: imHandler,
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
