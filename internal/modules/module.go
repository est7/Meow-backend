package modules

import (
	"Meow-backend/internal/interfaces"
	"Meow-backend/pkg/auth"
	"github.com/gin-gonic/gin"
)

type Module interface {
	Name() string
	Init(appCtx *interfaces.AppContext)
	RegisterRoutes(*gin.Engine, func(auth.PermissionLevel) gin.HandlerFunc)
}

var moduleFactories []func(*interfaces.AppContext) Module

func RegisterModuleFactory(factory func(*interfaces.AppContext) Module) {
	moduleFactories = append(moduleFactories, factory)
}

func InitModules(ctx *interfaces.AppContext) []Module {
	modules := make([]Module, 0, len(moduleFactories))
	for _, factory := range moduleFactories {
		module := factory(ctx)
		module.Init(ctx)
		modules = append(modules, module)
	}
	return modules
}
