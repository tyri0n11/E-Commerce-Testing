package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	r := gin.Default()
	v1 := r.Group("/version1/2024")
	{
		v1.GET("/ping", Pong)
		v1.PUT("/ping", Pong)
		v1.PATCH("/ping", Pong)
		v1.DELETE("/ping", Pong)
		v1.HEAD("/ping", Pong)
		v1.OPTIONS("/ping", Pong)
	}

	v2 := r.Group("/version2/2024")
	{
		v2.GET("/ping", Pong)
		v2.PUT("/ping", Pong)
		v2.PATCH("/ping", Pong)
		v2.DELETE("/ping", Pong)
		v2.HEAD("/ping", Pong)
		v2.OPTIONS("/ping", Pong)
	}
	return r
}
func Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "Tyr1on")
	uid := c.DefaultQuery("uid", "Tyr1on")
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the server",
		"name":    name,
		"uid":     uid,
	})
}
