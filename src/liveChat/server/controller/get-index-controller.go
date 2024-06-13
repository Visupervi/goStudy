package controller

import (
	"fmt"
	"liveChat/server/dao"
	"liveChat/server/model"
	"liveChat/server/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetIndex
// @Summary 获取用户列表
// @Tags         用户模块
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

// UserRegistry godoc
// @Summary      用户注册
// @Tags         用户模块
// @Accept       json
// @Produce      json
// @Param        user body model.UserRequestData false "用户注册信息"
// @Success      200  {object}  model.ResponseData
// @Failure      400  {object}  model.ResponseData
// @Failure      404  {object}  model.ResponseData
// @Failure      500  {object}  model.ResponseData
// @Router       /user/userRegistry [post]
func UserRegistry(c *gin.Context) {
	user := &model.UserBasic{}
	json := make(map[string]interface{}) // 注意该结构接受的内容
	c.BindJSON(&json)
	// log.Printf("%v", &json)
	if json["password"] != json["repassword"] {
		res := &model.ResponseData{
			Status: 0,
			Data:   "两次密码不一致",
		}
		c.JSON(0, res)
	}

	user.Password = json["password"].(string) // 类型断言
	user.Name = json["userName"].(string)     // 类型断言

	service.InsertUserToTable(user)
	res := &model.ResponseData{
		Status: 200,
		Data:   "添加成功",
	}
	c.JSON(200, res)

}

// DeleteUser godoc
// @Summary      删除用户
// @Tags         用户模块
// @Accept       json
// @Produce      json
// @Param        id path int false "用户id"
// @Success      200  {object}  model.ResponseData
// @Failure      400  {object}  model.ResponseData
// @Failure      404  {object}  model.ResponseData
// @Failure      500  {object}  model.ResponseData
// @Router       /user/deleteUser/{id} [delete]
func DeleteUser(c *gin.Context) {
	// json := make(map[string]interface{}) // 注意该结构接受的内容
	//
	id := c.Param("id")
	intNum, _ := strconv.Atoi(id)
	uintNum := uint(intNum)
	service.DeleteUserById(uintNum)
	res := &model.ResponseData{
		Status: 200,
		Data:   "删除成功",
	}
	c.JSON(200, res)
}

// UserUpdate godoc
// @Summary      用户更新
// @Tags         用户模块
// @Accept       json
// @Produce      json
// @Param        user body model.UserRequestData false "用户更新信息"
// @Success      200  {object}  model.ResponseData
// @Failure      400  {object}  model.ResponseData
// @Failure      404  {object}  model.ResponseData
// @Failure      500  {object}  model.ResponseData
// @Router       /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := &model.UserBasic{}
	json := make(map[string]interface{}) // 注意该结构接受的内容
	c.BindJSON(&json)
	// intNum, _ := strconv.Atoi(json["id"].(string))
	fmt.Println("intNum", json["id"])
	// uintNum := uint(intNum)
	// fmt.Println("uintNum", uintNum)
	// uint(json["id"].(float64))
	user.ID = uint(json["id"].(float64))
	user.Name = json["userName"].(string)
	user.Password = json["password"].(string)
	service.UpdateUser(user)
	res := &model.ResponseData{
		Status: 200,
		Data:   "更新成功",
	}
	c.JSON(200, res)
}
