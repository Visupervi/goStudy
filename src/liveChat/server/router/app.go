package router

import (
	"liveChat/server/controller"
	"liveChat/server/docs"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func App() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/index", controller.GetIndex)
	r.POST("/user/userRegistry", controller.UserRegistry)
	r.POST("/user/deleteUser/:id", controller.DeleteUser)
	return r
}
