package statistics

import (
	_ "github.com/Mikki21/dlv-project/controllers"
	_ "github.com/Mikki21/dlv-project/controllers/bank-rating"
	_ "github.com/Mikki21/dlv-project/controllers/best-bank"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

type StatisticController struct {
	beego.Controller
}

func (this *StatisticController) Get() {
	statistic := toolbox.StatisticsMap.GetMap()
	this.Data["Stat"] = statistic["Data"]
	this.TplName = "statistic.tpl"
}
