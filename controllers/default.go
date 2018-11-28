package controllers

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	toolbox.StatisticsMap.AddStatistics("GET", "/", "&controllers.MainController", time.Duration(13000))
	c.Layout = "main_layout.tpl"
	c.TplName = "index.tpl"
}
