package infrastructure

import (
	infrastructure "IkezawaYuki/craft/infrastructure/datastore"
	"IkezawaYuki/craft/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	bitflyerCtr := controllers.NewBitlyerController(infrastructure.Connect())

	router.POST("/api/candle/", func(c *gin.Context) { bitflyerCtr.StreamIngestionData() })

	Router = router
}
