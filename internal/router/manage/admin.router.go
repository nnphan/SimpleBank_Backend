package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct {
}

func (ar *AdminRouter) InitAdminRouter(router *gin.RouterGroup) {

	// Public router
	adminRouterPublic := router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}

	//Private router
	adminRouterPrivate := router.Group("/admin/user")
	{
		adminRouterPrivate.GET("/profile")
		adminRouterPrivate.PUT("/update-profile")
		adminRouterPrivate.POST("/active-user")
	}

}