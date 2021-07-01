package _Config

import (
	"github.com/spf13/viper"

	_Gin "project/modules/_Gin"
	api "project/modules/_Gin/controllers/api"
	_Log "project/modules/_Log"
	_MySQL "project/modules/_MySQL"
)

type AllConfigStruct struct {
	Mode     string                 `json:"Mode"`
	IP       string                 `json:"IP"`
	RemoteIP string                 `json:"RemoteIP"`
	CacheIP  string                 `json:"CacheIP"`
	Database _MySQL.ModuleCfgStruct `json:"Database"`
	Log      _Log.ModuleCfgStruct   `json:"Log"`
}

var AllConfig *AllConfigStruct

func ConfigInit() *AllConfigStruct {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	viper.ReadInConfig()

	viper.SetDefault("Mode", "r")
	viper.SetDefault("IP", ":80")
	viper.SetDefault("RemoteIP", ":9090")
	viper.SetDefault("CacheIP", "127.0.0.1:6379")

	viper.SetDefault("Database.SQL_IP", "127.0.0.1:3306")
	viper.SetDefault("Database.SQL_Account", "root")
	viper.SetDefault("Database.SQL_Password", "123")

	viper.SetDefault("Log.OutLevel", 1)
	viper.SetDefault("Log.Format", 1)
	viper.SetDefault("Log.Path", "./logs/")

	err := viper.Unmarshal(&AllConfig)
	if err != nil {
		panic("Config Read Error")
		// log.Fatal(err)
	}

	_Gin.ModuleCfg.Mode = AllConfig.Mode
	_Gin.ModuleCfg.IP = AllConfig.IP

	// _Redis.Redis_IP = AllConfig.CacheIP

	// _MySQL.ModuleCfg = &AllConfig.Database

	_Log.ModuleCfg = &AllConfig.Log

	api.SysInfoInit(
		AllConfig.IP,
		AllConfig.Mode,
		AllConfig.RemoteIP,
		AllConfig.Database.SQL_IP)

	// fmt.Println(AllConfig)

	return AllConfig
}

// ===========================================================
