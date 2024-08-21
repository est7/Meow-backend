package modules

import (
	"Meow-backend/internal/modules/v1/card"
	"Meow-backend/internal/modules/v1/feed"
	"Meow-backend/internal/modules/v1/im"
	"Meow-backend/internal/modules/v1/user"
	"github.com/gin-gonic/gin"
)

// Module : handler v1
type Module interface {
	GetName() string
	Init()
	InitRouter(rgPublic *gin.RouterGroup, rgPrivate *gin.RouterGroup)
}

var Modules []Module

func registerModule(m []Module) {
	Modules = append(Modules, m...)
}

func init() {
	// Register module here
	registerModule([]Module{
		&user.ModuleUser{},
		&feed.ModuleFeed{},
		&card.ModuleCard{},
		&im.ModuleIm{},
	})
}
