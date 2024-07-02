package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	c "github.com/tyri0n11/Muffin/internal/controllers"
)

func NewRouter() *gin.Engine {

	r := gin.Default()
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/user/1", c.NewUserController().GetUserById)

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
