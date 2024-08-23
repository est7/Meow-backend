package modules

import (
	"Meow-backend/internal/initialize"
	"Meow-backend/pkg/auth"
	"github.com/gin-gonic/gin"
)

type Module interface {
	Name() string
	Init(appCtx *initialize.AppInstance)
	RegisterRoutes(*gin.Engine, func(auth.PermissionLevel) gin.HandlerFunc)
}

var moduleFactories []func(*initialize.AppInstance) Module

func RegisterModuleFactory(factory func(*initialize.AppInstance) Module) {
	moduleFactories = append(moduleFactories, factory)
}

func InitModules(ctx *initialize.AppInstance) []Module {
	modules := make([]Module, 0, len(moduleFactories))
	for _, factory := range moduleFactories {
		module := factory(ctx)
		module.Init(ctx)
		modules = append(modules, module)
	}
	return modules
}
