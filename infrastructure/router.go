package infrastructure

import (
	infrastructure "IkezawaYuki/craft/infrastructure/datastore"
	"IkezawaYuki/craft/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	bitflyerCtr := controllers.NewBitlyerController(infrastructure.Connect())

	router.POST("/users", func(c *gin.Context) { bitflyerCtr.StreamIngestionData() })

	Router = router
}
