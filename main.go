package main

import (
	"embed"
	_Config "project/modules/_Config"
	_Gin "project/modules/_Gin"
	_Log "project/modules/_Log"
	// _MySQL "project/modules/_MySQL"
	// _Redis "project/modules/_Redis"
)

//go:embed views/* public/*
var f embed.FS

func main() {

	//===== Load Config ./config.json
	_Config.ConfigInit()

	//===== Log
	_Log.Loginit()

	//===== Redis
	//_Redis.RedisInit()

	//===== MySQL
	//_MySQL.DatabaseInit()

	//===== 設定Gin運行模式
	_Gin.GinInit(f)
}
