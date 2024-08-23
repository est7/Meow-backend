package card

import (
	"Meow-backend/internal/initialize"
	"Meow-backend/internal/modules"
	"Meow-backend/pkg/auth"
	"Meow-backend/pkg/log"
	"github.com/gin-gonic/gin"
)

type CardModule struct {
	appCtx *initialize.AppInstance
}

func NewCardModule(ctx *initialize.AppInstance) modules.Module {
	return &CardModule{ctx}
}

func (u *CardModule) Name() string {
	return "Card"
}

func (u *CardModule) Init(appCtx *initialize.AppInstance) {
	u.appCtx = appCtx
	log.Info("Initializing card module")
}

func (u *CardModule) RegisterRoutes(r *gin.Engine, authMiddleware func(auth.PermissionLevel) gin.HandlerFunc) {
}
