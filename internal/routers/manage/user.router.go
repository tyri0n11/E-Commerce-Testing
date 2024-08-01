package manage

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//public routers
	// UserRouterPublic := Router.Group("/admin/user")
	// {
	// 	UserRouterPublic.POST("/register")
	// 	UserRouterPublic.POST("/otp")
	// }
	//private routers
	UserRouterPrivate := Router.Group("admin/user")
	// UserRouterPrivate.Use(Limiter())
	// UserRouterPrivate.Use(Authen())
	// UserRouterPrivate.Use(Permission())
	{
		UserRouterPrivate.GET("/active_user")
	}
}
