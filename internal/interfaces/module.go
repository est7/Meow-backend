package interfaces

import (
	"Meow-backend/pkg/auth"
	"github.com/gin-gonic/gin"
)

type Module interface {
	Name() string
	Init(appCtx AppContext)
	RegisterRoutes(*gin.Engine, func(auth.PermissionLevel) gin.HandlerFunc)
}

var moduleFactories []func(AppContext) Module

func RegisterModuleFactory(factory func(AppContext) Module) {
	moduleFactories = append(moduleFactories, factory)
}

func InitModules(ctx AppContext) []Module {
	modules := make([]Module, 0, len(moduleFactories))
	for _, factory := range moduleFactories {
		module := factory(ctx)
		module.Init(ctx)
		modules = append(modules, module)
	}
	return modules
}
