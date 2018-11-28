package main

import (
	"os"
	"strconv"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/oreuta/easytrip/routers"
	"github.com/oreuta/easytrip/sql1"
	_ "github.com/oreuta/easytrip/sql1"
)

func main() {
	var err error
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	beego.BConfig.Listen.HTTPPort, err = strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	sql1.Db = sql1.CreateConnect(beego.AppConfig.String("connect"))
	go sql1.Update(sql1.Db)
	beego.Run()
}
