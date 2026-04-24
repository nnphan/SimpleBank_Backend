package user

import "github.com/gin-gonic/gin"

type ProductRouter struct {
}

func (pr *ProductRouter) InitProductRouter(router *gin.RouterGroup)  {

	//public router 
	productRouterPublic := router.Group("/product")
	{
		productRouterPublic.GET("/list")
		productRouterPublic.GET("/search")
		productRouterPublic.GET("/detail/:id")
	}

	//private router
	productRouterPrivate := router.Group("/product")	
	{
		productRouterPrivate.POST("/create")
		productRouterPrivate.PUT("/update/:id")
		productRouterPrivate.DELETE("/delete/:id")
	}
}