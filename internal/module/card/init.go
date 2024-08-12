package card

import "github.com/gin-gonic/gin"

type ModuleCard struct{}

func (u *ModuleCard) GetName() string {
	return "Card"
}

func (u *ModuleCard) Init() {}

func (u *ModuleCard) InitRouter(rgPublic *gin.RouterGroup, rgPrivate *gin.RouterGroup) {

}
