package bankRatingController

import (
	"time"

	"github.com/Mikki21/dlv-project/services/index-serv"

	"github.com/Mikki21/dlv-project/models"
	"github.com/Mikki21/dlv-project/services/bank-rating"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

//RatesController is a controller for comparing page
type RatesController struct {
	beego.Controller
	RatesService bankRatingService.RatesServiceInterface
	Indexserv    indexService.IndexServiceInterface
}

//New create a new RatesController
func New(service bankRatingService.RatesServiceInterface, s1 indexService.IndexServiceInterface) *RatesController {
	return &RatesController{
		RatesService: service,
		Indexserv:    s1,
	}
}

//Get function gets request gives and output data on display
func (this *RatesController) Get() {
	toolbox.StatisticsMap.AddStatistics("GET", "/comparision", "&controllers.bankRatingController.RatesController", time.Duration(13000))
	r := models.MainRequest{
		Currency: this.GetStrings("currency"),
		Option:   this.GetString("option"),
		Bank:     this.GetStrings("bank"),
	}
	ind, _ := this.Indexserv.IndexGet()
	i := 0
	if r.Currency == nil {
		this.Data["IncorrectCurrency"] = true
		i++
	}
	if r.Bank == nil {
		this.Data["IncorrectBank"] = true
		i++
	}
	if i > 0 {
		this.Data["MainNumber"] = this.Indexserv.ToFixed((ind[0].RateBuy+ind[1].RateBuy)/2, 2)
		this.Data["NBU"] = ind[0]
		this.Data["Others"] = ind[1]
		this.Layout = "main_layout.tpl"
		this.TplName = "index.tpl"
		this.LayoutSections = make(map[string]string)
		this.LayoutSections["Second"] = "today.tpl"
		return
	}
	b, err := this.RatesService.GetBankRates(r)
	if err != nil {
		beego.Error("Error:%v", err)
		return
	}

	this.Data["Banks"] = b

	this.Layout = "comparision_layout.tpl"
	this.TplName = "comparision.tpl"
}
