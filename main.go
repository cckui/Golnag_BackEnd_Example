package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	_Config "project/modules/_Config"
	_Gin "project/modules/_Gin"
	_Log "project/modules/_Log"

	// _MySQL "project/modules/_MySQL"
	// _Redis "project/modules/_Redis"
	"database/sql"

	_ "github.com/sijms/go-ora/v2"
)

//go:embed views/* public/*
var f embed.FS
var (
	username   = flag.String("uname", "scott", "oracle username")
	password   = flag.String("password", "tiger", "oracle password")
	oraclehost = flag.String("oraclehost", "dbhost", "oracle database host")
	oracleport = flag.Int("oracleport", 1521, "oracle database port")
	dbname     = flag.String("dbname", "orclpdb1", "oracle database name")
)

func main() {
	flag.Parse()

	//===== Load Config ./config.json
	_Config.ConfigInit()

	//===== Log
	_Log.Loginit()

	//===== Redis
	//_Redis.RedisInit()

	//===== MySQL
	//_MySQL.DatabaseInit()

	//===== 設定Gin運行模式

	if false {
		oracleTest()
	}
	_Gin.GinInit(f)
}

func oracleTest() {

	osqlInfo := fmt.Sprintf("oracle://%s:%s@%s:%d/%s", *username, *password, *oraclehost, *oracleport, *dbname)

	db, err := sql.Open("oracle", osqlInfo)

	if err != nil {
		log.Fatalf("connect oracle db error: %s:", err.Error())
	}

	rows, err := db.Query("select to_char(sysdate,'yyyy-mm-dd hh24:mi:ss') AS name from dual")

	if err != nil {
		fmt.Println("exec query error:", err.Error())
	}

	for rows.Next() {
		var name string
		rows.Scan(&name)
		fmt.Println("fetch item:")
		fmt.Println(name)
	}
}
