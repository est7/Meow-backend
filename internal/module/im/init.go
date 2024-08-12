package im

import "github.com/gin-gonic/gin"

type ModuleIm struct{}

func (u *ModuleIm) GetName() string {
	return "Im"
}

func (u *ModuleIm) Init() {}

func (u *ModuleIm) InitRouter(rgPublic *gin.RouterGroup, rgPrivate *gin.RouterGroup) {

}
