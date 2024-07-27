package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *ProductRouter) InitUserRouter(Router *gin.RouterGroup) {
	//public routers
	UserRouterPublic := Router.Group("/product")
	{
		UserRouterPublic.POST("/register")
		UserRouterPublic.POST("/otp")
	}
	//private routers
	UserRouterPrivate := Router.Group("/product")
	// UserRouterPrivate.Use(Limiter())
	// UserRouterPrivate.Use(Authen())
	// UserRouterPrivate.Use(Permission())
	{
		UserRouterPrivate.GET("/get_info")
	}
}
