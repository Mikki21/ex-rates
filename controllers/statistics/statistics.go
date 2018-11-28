package statistics

import (
	_ "github.com/oreuta/easytrip/controllers"
	_ "github.com/oreuta/easytrip/controllers/bank-rating"
	_ "github.com/oreuta/easytrip/controllers/best-bank"

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
