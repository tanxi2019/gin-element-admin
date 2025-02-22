package main

import (
	"server/initialize"
)

// @title Swagger API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	initialize.InitConfig()         // 加载配置文件到全局配置结构体
	initialize.InitLogger()         // 初始化日志
	initialize.InitRedis()          // 初始化数据库redis
	initialize.InitMysql()          // 初始化数据库mysql
	initialize.InitCasbinEnforcer() // 初始化casbin策略
	initialize.InitValidate()       // 初始化Validator数据校验
	initialize.InitRun()            // 初始化服务器
}
