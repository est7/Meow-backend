package im

import (
	"Meow-backend/internal/initialize"
	"Meow-backend/internal/modules"
	"Meow-backend/pkg/auth"
	"Meow-backend/pkg/log"
	"github.com/gin-gonic/gin"
)

type IMModule struct {
	appCtx *initialize.AppInstance
}

func NewIMModule(appCtx *initialize.AppInstance) modules.Module {
	return &IMModule{appCtx: appCtx}
}

func (u *IMModule) Name() string {
	return "IM"
}

func (u *IMModule) Init(appCtx *initialize.AppInstance) {
	u.appCtx = appCtx
	log.Info("Initializing im module")
}

func (u *IMModule) RegisterRoutes(r *gin.Engine, authMiddleware func(auth.PermissionLevel) gin.HandlerFunc) {
}
