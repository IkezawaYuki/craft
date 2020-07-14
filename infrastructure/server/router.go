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
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			"https://localhost:8080",
		},
		// 自分で許可するしないの処理を書きたい場合は、以下のように書くこともできる
		// AllowOriginFunc: func(origin string) bool {
		//  return origin == "https://www.example.com:8080"
		// },
		// preflight requestで許可した後の接続可能時間
		// https://godoc.org/github.com/gin-contrib/cors#Config の中のコメントに詳細あり
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
