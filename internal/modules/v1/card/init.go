package card

import (
	"Meow-backend/internal/interfaces"
	"Meow-backend/internal/modules"
	"Meow-backend/pkg/auth"
	"Meow-backend/pkg/log"
	"github.com/gin-gonic/gin"
)

type CardModule struct {
	appCtx *interfaces.AppContext
}

func NewCardModule(ctx *interfaces.AppContext) modules.Module {
	return &CardModule{ctx}
}

func (u *CardModule) Name() string {
	return "Card"
}

func (u *CardModule) Init(appCtx *interfaces.AppContext) {
	u.appCtx = appCtx
	log.Info("Initializing card module")
}

func (u *CardModule) RegisterRoutes(r *gin.Engine, authMiddleware func(auth.PermissionLevel) gin.HandlerFunc) {
}
