package server

import (
	"github.com/gin-gonic/gin"
	"fabric-sdk-go/sdkInit"
)

var App sdkInit.Application

func ServerInit(app sdkInit.Application) {
	App = app
	r := gin.Default()
	r.GET("/query/:ID", QueryAssets)
	r.POST("/add", AddAssets)
	r.Run(":8080")
}