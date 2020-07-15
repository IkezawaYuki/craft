package server

import (
	"IkezawaYuki/craft/infrastructure/middlewares"
	"IkezawaYuki/craft/interfaces/controllers"
	"IkezawaYuki/craft/registry"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

var Router *gin.Engine

func Init() {
	r := gin.Default()

	ctn, err := registry.NewContainer()
	if err != nil {
		panic(err)
	}

	bitflyerCtr := ctn.Resolve("bitflyer-controller").(controllers.BitlyerController)

	r.LoadHTMLGlob("infrastructure/views/*.html")

	go bitflyerCtr.StreamIngestionData()

	r.Use(middlewares.Firebase())
	r.Use(cors.New(cors.Config{
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		AllowOrigins: []string{
			"https://localhost:8080",
		},
		MaxAge: 24 * time.Hour,
	}))

	r.POST("/api/candles/", func(c *gin.Context) {
		bitflyerCtr.ApiCandleHandler(c)
	})

	r.GET("/chart", func(c *gin.Context) {
		bitflyerCtr.ViewChart(c)
	})

	Router = r
	if err := Router.Run(":8080"); err != nil {
		panic(err)
	}
}
