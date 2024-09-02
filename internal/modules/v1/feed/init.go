package feed

import (
	"Meow-backend/internal/interfaces"
	"Meow-backend/pkg/auth"
	"Meow-backend/pkg/log"
	"github.com/gin-gonic/gin"
)

type FeedModule struct {
	appCtx interfaces.AppContext
}

func NewFeedModule(appCtx interfaces.AppContext) interfaces.Module {
	return &FeedModule{appCtx: appCtx}
}

func (u *FeedModule) Name() string {
	return "Feed"
}

func (u *FeedModule) Init(appCtx interfaces.AppContext) {
	u.appCtx = appCtx
	log.Info("Initializing feed module")
}

func (u *FeedModule) RegisterRoutes(r *gin.Engine, authMiddleware func(auth.PermissionLevel) gin.HandlerFunc) {
}
