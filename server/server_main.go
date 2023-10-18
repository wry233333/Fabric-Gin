package server

import (
	"github.com/gin-gonic/gin"
	"fabric-sdk-go/sdkInit"
	"github.com/gin-contrib/cors"
)

var App sdkInit.Application

func ServerInit(app sdkInit.Application) {
	App = app
	r := gin.Default()
	r.Use(cors.Default())
	assetRouter := r.Group("/asset")
	// 资产管理
	{
		assetRouter.GET("/query/:ID", QueryAssets)
		assetRouter.POST("/add", AddAssets)
	}
	
	r.Run(":3000")
}