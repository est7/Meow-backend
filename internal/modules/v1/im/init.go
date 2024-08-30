package im

import (
	"Meow-backend/internal/interfaces"
	"Meow-backend/internal/modules"
	"Meow-backend/pkg/auth"
	"Meow-backend/pkg/log"
	"github.com/gin-gonic/gin"
)

type IMModule struct {
	appCtx *interfaces.AppContext
}

func NewIMModule(appCtx *interfaces.AppContext) modules.Module {
	return &IMModule{appCtx: appCtx}
}

func (u *IMModule) Name() string {
	return "IM"
}

func (u *IMModule) Init(appCtx *interfaces.AppContext) {
	u.appCtx = appCtx
	log.Info("Initializing im module")
}

func (u *IMModule) RegisterRoutes(r *gin.Engine, authMiddleware func(auth.PermissionLevel) gin.HandlerFunc) {
}
