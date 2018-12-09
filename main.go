package main

import (
	"os"
	"strconv"

	_ "github.com/Mikki21/dlv-project/routers"
	"github.com/Mikki21/dlv-project/sql1"
	_ "github.com/Mikki21/dlv-project/sql1"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var err error
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	beego.BConfig.Listen.HTTPPort, err = strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	go sql1.Update()
	beego.Run()
}
