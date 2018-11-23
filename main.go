package main

import (
	"os"
	"strconv"

	"github.com/astaxie/beego"
	_ "github.com/oreuta/easytrip/routers"
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
