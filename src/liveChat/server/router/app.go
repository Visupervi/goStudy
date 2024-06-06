package router

import (
	"github.com/gin-gonic/gin"
	"liveChat/server/controller"
)

func App() *gin.Engine {
	r := gin.Default()
	r.GET("/index", controller.GetIndex)
	return r
}
