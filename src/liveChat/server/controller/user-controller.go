package controller

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"liveChat/server/dao"
	"liveChat/server/model"
	"liveChat/server/service"
	"liveChat/server/utils"
	"math/rand"
	"strconv"
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
	uuidValue := uuid.New()
	json := make(map[string]interface{}) // 注意该结构接受的内容
	c.BindJSON(&json)
	salt := fmt.Sprintf("%06d", rand.Int31())

	// log.Printf("%v", &json)
	if json["password"] != json["repassword"] {
		res := &model.ResponseData{
			Status: 0,
			Data:   "两次密码不一致",
		}
		c.JSON(0, res)
		return
	}
	u := service.GetUserByName(json["userName"].(string))

	if u.Name != "" {
		res := &model.ResponseData{
			Status: 0,
			Data:   "用户名已存在",
		}
		c.JSON(0, res)
		return
	}

	user.Password = utils.MakeRandomNum(json["password"].(string), salt) // 类型断言
	user.Name = json["userName"].(string)                                // 类型断言
	user.Salt = salt
	user.Identity = uuidValue.String()
	service.InsertUserToTable(user)
	res := &model.ResponseData{
		Status: 200,
		Data:   "添加成功",
	}
	c.JSON(200, res)
	return

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

// UpdateUser godoc
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
	// fmt.Println("intNum", json["id"])
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

// Login godoc
// @Summary      用户登陆
// @Tags         用户模块
// @Accept       json
// @Produce      json
// @Param        user body model.UserRequestData false "用户登陆"
// @Success      200  {object}  model.ResponseData
// @Failure      400  {object}  model.ResponseData
// @Failure      404  {object}  model.ResponseData
// @Failure      500  {object}  model.ResponseData
// @Router       /user/login [post]
func Login(c *gin.Context) {
	// user := &model.UserBasic{}
	conn := utils.Pool.Get()
	defer conn.Close()
	res := &model.ResponseData{}
	result := make(map[string]interface{}) // 注意该结构接受的内容
	c.BindJSON(&result)
	user := service.GetUserByName(result["userName"].(string))

	if user.Identity == "" {
		res.Status = 0
		res.Data = "用户不存在"
	} else {
		flag := utils.ValidPassword(result["password"].(string), user.Salt, user.Password)
		if flag {
			data, _ := json.Marshal(user)
			strData := string(data)
			redis.String(conn.Do("hset", "user", user.Password, strData))
			res.Status = 200
			res.Data = "登陆成功"
		} else {
			res.Status = 0
			res.Data = "登陆失败"
		}
	}

	c.JSON(200, res)
}
