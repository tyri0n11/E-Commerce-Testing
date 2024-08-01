package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/tyri0n11/Muffin/global"
	"github.com/tyri0n11/Muffin/internal/routers"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// r.Use() //logger
	// r.Use() //cross
	// r.Use() //limiter
	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	mainGroup := r.Group("/v1/2024")
	{
		mainGroup.GET("/checkStatus") //tracking monitor
	}
	{
		userRouter.InitUserRouter(mainGroup)
		userRouter.InitProductRouter(mainGroup)
	}
	{
		managerRouter.InitAdminRouter(mainGroup)
		managerRouter.InitUserRouter(mainGroup)
	}
	return r
}
