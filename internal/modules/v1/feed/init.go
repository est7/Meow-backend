package feed

import (
	"Meow-backend/internal/initialize"
	"Meow-backend/internal/modules"
	"Meow-backend/pkg/auth"
	"Meow-backend/pkg/log"
	"github.com/gin-gonic/gin"
)

type FeedModule struct {
	appCtx *initialize.AppInstance
}

func NewFeedModule(appCtx *initialize.AppInstance) modules.Module {
	return &FeedModule{appCtx: appCtx}
}

func (u *FeedModule) Name() string {
	return "Feed"
}

func (u *FeedModule) Init(appCtx *initialize.AppInstance) {
	u.appCtx = appCtx
	log.Info("Initializing feed module")
}

func (u *FeedModule) RegisterRoutes(r *gin.Engine, authMiddleware func(auth.PermissionLevel) gin.HandlerFunc) {
}
