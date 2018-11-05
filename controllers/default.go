package controllers

import (
	"github.com/astaxie/beego"
)

type Common struct {
	beego.Controller
}

func (this *Common) Get() {
	this.TplName = "index.tpl"
}
