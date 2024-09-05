package card

import (
	"Meow-backend/internal/interfaces"
	"Meow-backend/pkg/auth"
	"Meow-backend/pkg/log"
	"github.com/gin-gonic/gin"
)

type CardModule struct {
	appCtx interfaces.AppContext
}

func NewCardModule(ctx interfaces.AppContext) interfaces.Module {
	//repo := interfaces.NewRepository(ctx.GetGormDB())
	//serviceFactory := factoryforservice.NewCardServiceFactory()
	//service := serviceFactory.CreateService(repo, ctx.GetRedisClient())
	//cardHandler := handler.NewCardHandler(service)
	//return &CardModule{
	//	appCtx:  ctx,
	//	handler: cardHandler,
	//}
	return nil
}
func (u *CardModule) Name() string {
	return "Card"
}

func (u *CardModule) Init(appCtx interfaces.AppContext) {
	u.appCtx = appCtx
	log.Info("Initializing card module")
}

func (u *CardModule) RegisterRoutes(r *gin.Engine, authMiddleware func(auth.PermissionLevel) gin.HandlerFunc) {
}
