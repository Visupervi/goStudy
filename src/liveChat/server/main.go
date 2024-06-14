package main

import (
	"liveChat/server/dao"
	"liveChat/server/router"
	"liveChat/server/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	dao.CreateTable()
	r := router.App()
	r.Run(":9966") // 监听并在 0.0.0.0:8080 上启动服务 10.10.211.68
}
