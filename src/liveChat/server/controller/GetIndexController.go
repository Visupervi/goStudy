package controller

import (
	"github.com/gin-gonic/gin"
	"liveChat/server/dao"
)

// GetIndex
// @Tags         首页
// @Accept       json
// @Produce      json
// @Success      200  {string}   success
// @Router       /index [get]
func GetIndex(c *gin.Context) {
	dao.GetUserList()
	c.JSON(200, gin.H{
		"message": "success",
	})
}
