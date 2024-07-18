package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *ProductRouter) InitUserRouter(Router *gin.RouterGroup) {
	//public routers
	UserRouterPublic := Router.Group("/product")
	{
		UserRouterPublic.POST("/register")
		UserRouterPublic.GET("/otp")
	}
	//private routers
	UserRouterPrivate := Router.Group("/product")
	{
		UserRouterPrivate.POST("/register")
		UserRouterPrivate.GET("/otp")
	}
}
