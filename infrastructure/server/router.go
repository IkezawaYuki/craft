package server

import (
	"IkezawaYuki/craft/interfaces/controllers"
	"IkezawaYuki/craft/registry"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func Init() {
	router := gin.Default()

	ctn, err := registry.NewContainer()
	if err != nil {
		panic(err)
	}

	bitflyerCtr := ctn.Resolve("bitflyer-controller").(controllers.BitlyerController)

	router.LoadHTMLGlob("infrastructure/views/*.html")

	go bitflyerCtr.StreamIngestionData()

	router.POST("/api/candle/", func(c *gin.Context) {
		bitflyerCtr.ApiCandleHandler(c)
	})

	router.GET("/chart", func(c *gin.Context) {
		bitflyerCtr.ViewChart(c)
	})

	Router = router
	if err := Router.Run(":8080"); err != nil {
		panic(err)
	}
}
