package Controllers

import (
	"fmt"
	"net/http"
	"store-api/Service"
	"store-api/Models"
	"github.com/gin-gonic/gin"
)

func TransactionController(c *gin.Context) {
	var orders []Models.Orders
	c.BindJSON(&orders)
	err := Service.Transaction(&orders)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, orders)
	}
}
