package main

import (
	"os"
	"strconv"

	_ "github.com/Mikki21/dlv-project/routers"
	"github.com/astaxie/beego"
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

	beego.Run()
}
