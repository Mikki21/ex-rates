package indexServiceController

import (
	"github.com/Mikki21/dlv-project/services/index-serv"
	"github.com/astaxie/beego"
)

//RatesController is a controller for comparing page
type IndexController struct {
	beego.Controller
	IndexService indexService.IndexServiceInterface
}

//New create a new RatesController
func New(service indexService.IndexServiceInterface) *IndexController {
	return &IndexController{
		IndexService: service,
	}
}

//Get function gets request gives and output data on display
func (this *IndexController) Get() {
	b, err := this.IndexService.IndexGet()
	if err != nil {
		beego.Error("Error:%v", err)
		return
	}
	this.Data["MainNumber"] = this.IndexService.ToFixed((b[0].RateBuy+b[1].RateBuy)/2, 2)
	this.Data["NBU"] = b[0]
	this.Data["Others"] = b[1]
	this.Layout = "main_layout.tpl"
	this.TplName = "index.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Second"] = "today.tpl"
}
