package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	c "github.com/tyri0n11/Muffin/internal/controllers"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before ---- AA")
		c.Next()
		fmt.Println("After ---- AA")
	}

}
func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Beffore ---- BB")
		c.Next()
		fmt.Println("After ---- BB")
	}
}
func CC(c *gin.Context) {
	fmt.Println("Beffore ---- CC")
	c.Next()
	fmt.Println("After ---- CC")
}

func NewRouter() *gin.Engine {

	r := gin.Default()
	//use middleware
	r.Use(AA(), BB(), CC)
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
