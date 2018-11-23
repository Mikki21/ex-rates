package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) ErrorDb() {
	this.TplName = "err.tpl"
}
