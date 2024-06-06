package controller

import (
	"github.com/gin-gonic/gin"
	"liveChat/server/dao"
)

func GetIndex(c *gin.Context) {
	dao.GetUserList()
	c.JSON(200, gin.H{
		"message": "success",
	})
}
