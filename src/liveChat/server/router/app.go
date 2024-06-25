package router

import (
	"liveChat/server/controller"
	"liveChat/server/docs"
	"liveChat/server/utils"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func App() *gin.Engine {
	r := gin.Default()
	// 跨域设置
	r.Use(utils.Cors())
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/index", controller.GetIndex)
	r.POST("/user/userRegistry", controller.UserRegistry)
	r.DELETE("/user/deleteUser/:id", controller.DeleteUser)
	r.POST("/user/updateUser", controller.UpdateUser)
	r.POST("/user/login", controller.Login)
	return r
}
