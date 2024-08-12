package feed

import "github.com/gin-gonic/gin"

type ModuleFeed struct{}

func (u *ModuleFeed) GetName() string {
	return "Feed"
}

func (u *ModuleFeed) Init() {}

func (u *ModuleFeed) InitRouter(rgPublic *gin.RouterGroup, rgPrivate *gin.RouterGroup) {

}
