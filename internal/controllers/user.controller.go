package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tyri0n11/Muffin/internal/service"
	"github.com/tyri0n11/Muffin/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// controller -> service -> repo -> models -> dbs
func (uc *UserController) GetUserById(c *gin.Context) {

	// response.SuccessResponse(c, 20001, []string{"User", "Tyr1on"})
	response.ErrorResponse(c, 20002)
}
