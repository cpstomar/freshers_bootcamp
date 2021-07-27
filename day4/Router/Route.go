package Router

import (
	"github.com/gin-gonic/gin"
	"store-api/Controllers"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("")
	{
		grp1.GET("product", Controllers.GetAllProductController)
		grp1.POST("product", Controllers.AddProductController)
		grp1.GET("product/:id", Controllers.GetProductbyIDController)
		grp1.PATCH("product/:id", Controllers.UpdateProductController)

		grp1.GET("order/:id", Controllers.GetOrderByIDController)
		grp1.POST("order", Controllers.PlaceOrderController)

		grp1.POST("user", Controllers.CreateUserController)

		grp1.GET("transaction", Controllers.TransactionController)
	}
	return r
}
