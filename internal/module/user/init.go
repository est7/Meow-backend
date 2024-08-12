package user

import "github.com/gin-gonic/gin"

type ModuleUser struct{}

func (u *ModuleUser) GetName() string {
	return "User"
}

func (u *ModuleUser) Init() {}

func (u *ModuleUser) InitRouter(rgPublic *gin.RouterGroup, rgPrivate *gin.RouterGroup) {
	// 认证相关路由
	rgPublic.POST("/register", Register)
	rgPublic.POST("/login", Login)
	rgPublic.POST("/login/phone", PhoneLogin)
	rgPublic.GET("/vcode", VCode) //验证码
	// 用户
	rgPublic.GET("/users/:id", GetUser) //获取用户信息
}

func GetUser(context *gin.Context) {

}

func VCode(context *gin.Context) {

}

func PhoneLogin(context *gin.Context) {

}

func Login(context *gin.Context) {

}
