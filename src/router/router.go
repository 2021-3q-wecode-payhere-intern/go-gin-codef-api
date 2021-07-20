package router

import (
	"go-gin-codef-api/src/controller"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	depositGet := r.Group("/deposit")
	{
		depositGet.GET("", controller.GetDepositByDaily)
		depositGet.GET("/:date", controller.GetDepositByCard)
	}

	businessGet := r.Group("/business")
	{
		businessGet.GET("/status", controller.GetBusinessStatus)
	}

	return r
}
