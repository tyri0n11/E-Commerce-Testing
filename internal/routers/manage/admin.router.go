package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (ar *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	//public routers
	AdminRouterPublic := Router.Group("/admin")
	{
		AdminRouterPublic.POST("/login")
	}
	//private routers
	AdminRouterPrivate := Router.Group("/admin/user")
	// AdminRouterPrivate.Use(Limiter())
	// AdminRouterPrivate.Use(Authen())
	// AdminRouterPrivate.Use(Permission())
	{
		AdminRouterPrivate.GET("/activate_user")
	}
}
