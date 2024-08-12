package module

import (
	"Meow-backend/internal/module/card"
	"Meow-backend/internal/module/feed"
	"Meow-backend/internal/module/im"
	"Meow-backend/internal/module/user"
	"github.com/gin-gonic/gin"
)

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
