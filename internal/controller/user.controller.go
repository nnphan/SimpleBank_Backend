package controller

import (
	"simplebank/internal/service"
	"simplebank/response"

	"github.com/gin-gonic/gin"
)

type UserController struct{
	userService service.IUserService
}

func NewUserController(userService service.IUserService ) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("123", "abc", "rty")
	response.SuccessResponse(c, result, nil)
}
