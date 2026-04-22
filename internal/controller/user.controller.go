package controller

import (
	"net/http"
	"simplebank/internal/common"
	"simplebank/internal/service"

	"github.com/gin-gonic/gin"
)

type UserController struct{
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	uuid := c.Query("uuid")
	c.JSON(http.StatusOK, gin.H{
		"message": "ping ....." + uc.userService.GetUserInfo(),
		"uuid":    uuid,
		"users":   []string{"cr7", "m10"},
	})
}

func (uc *UserController) CreateUser(ctx *gin.Context) {

	// 1. Define request body struct
	var req struct {
		FullName string `json:"full_name"`
        Email    string `json:"email"`
        Password string `json:"password"`
	}

	
	
	// 2. Bind JSON body
	if err := ctx.ShouldBindJSON(&req); err != nil {
        common.Error(ctx, 400, "invalid request body")
        return
    }

	// 3. Call service
    user, err := uc.userService.CreateUser(
        ctx,
        req.FullName,
        req.Email,
        req.Password,
    )

	if err != nil {
        common.Error(ctx, 500, "cannot create user")
        return
    }

	// 4. Response 
	common.Success(
        ctx,
        201,
        "user created successfully",
        user,
    )

}

