package user

import (
	"Meow-backend/internal/interfaces"
	factrory "Meow-backend/internal/interfaces/factoryforservice"
	"Meow-backend/internal/modules/v1/user/handler"
	"Meow-backend/pkg/auth"
	"Meow-backend/pkg/log"
	"github.com/gin-gonic/gin"
)

type UserModule struct {
	appCtx  interfaces.AppContext
	handler *handler.UserHandler
}

func NewUserModule(ctx interfaces.AppContext) interfaces.Module {
	repo := interfaces.NewRepository(ctx.GetGormDB())
	serviceFactory := factrory.NewUserServiceFactory()
	service := serviceFactory.CreateService(repo, ctx.GetRedisClient())
	userHandler := handler.NewUserHandler(service)

	return &UserModule{
		appCtx:  ctx,
		handler: userHandler,
	}
}

func (u *UserModule) Name() string {
	return "User"
}

func (u *UserModule) Init(appCtx interfaces.AppContext) {
	u.appCtx = appCtx
	log.Info("Initializing user module")
}

func (u *UserModule) RegisterRoutes(r *gin.Engine, authMiddleware func(auth.PermissionLevel) gin.HandlerFunc) {
	// 公开路由
	//
	//1. 注册 (Register)
	//2. 登录 (Login)
	//3. 找回密码 (Forgot Password)
	//4. 重置密码 (Reset Password) - 通常通过一次性令牌访问
	//5. 刷新令牌 (Refresh Token) - 使用 refresh token 获取新的 access token
	//6. 获取 sms 验证码 (Get SMS Code)
	public := r.Group("/api/v1/user", authMiddleware(auth.Public))
	{
		public.POST("/register/username", u.handler.RegisterWithUsernameHandler)
		public.POST("/register/email", u.handler.RegisterWithEmailHandler)
		public.POST("/register/phone", u.handler.RegisterWithPhoneHandler)

		public.POST("/login-username", u.handler.UsernameLoginHandler)
		public.POST("/login-email", u.handler.EmailLoginHandler)
		public.POST("/login-phone", u.handler.PhoneLoginHandler)

		public.POST("/forgot-password", u.handler.ForgotPasswordHandler)
		public.POST("/reset-password", u.handler.ResetPasswordHandler)

		public.POST("/refresh-token", u.handler.RefreshTokenHandler)

		public.GET("/vcode", u.handler.VCode)     // 验证码
		public.GET("/smscode", u.handler.SmsCode) // 短信验证码
	}

	// 需要认证的路由
	//
	//1. 修改密码 (Change Password)
	//2. 绑定手机号 (Bind Phone Number)
	//3. 绑定邮箱 (Bind Email)
	//4. 删除账号 (Delete Account)
	//5. 获取用户信息 (Get User Profile)
	//6. 更新用户信息 (Update User Profile)
	//7. 登出 (Logout) - 虽然客户端可以自行删除 token，服务端的登出接口可用于撤销 token
	// 认证相关路由
	authenticated := r.Group("/api/v1/user", authMiddleware(auth.Authenticated))
	{
		authenticated.POST("/change-password", u.handler.ChangePasswordHandler)
		//authenticated.POST("/bind-phone", handler.BindPhoneHandler)
		//authenticated.POST("/bind-email", handler.BindEmailHandler)
		//authenticated.DELETE("/account", handler.DeleteAccountHandler)
		//authenticated.GET("/profile", handler.GetUserProfileHandler)
		//authenticated.PUT("/profile", handler.UpdateUserProfileHandler)
		//authenticated.POST("/logout", handler.LogoutHandler)
	}

	// 如果需要管理员路由，可以添加如下代码
	// admin := r.Group("/api/v1/admin/user", authMiddleware(auth.Admin))
	// {
	//     admin.GET("/users", handler.GetAllUsersHandler)
	//     admin.DELETE("/users/:id", handler.DeleteUserByAdminHandler)
	// }
}
