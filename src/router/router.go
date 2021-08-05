package router

import (
	"go-gin-codef-api/src/controller"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "go-gin-codef-api/docs"
)

func Router() *gin.Engine {
	r := gin.Default()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

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
