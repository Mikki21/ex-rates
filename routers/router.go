package routers

import (
	"github.com/Mikki21/dlv-project/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.Common{})
	beego.Router("/bestrates", &controllers.UsdController{}, "get:RequestWork")
}
