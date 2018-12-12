package bankRatingController

import (
	"time"

	"github.com/Mikki21/dlv-project/models"
	"github.com/Mikki21/dlv-project/services/bank-rating"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

//RatesController is a controller for comparing page
type RatesController struct {
	beego.Controller
	RatesService bankRatingService.RatesServiceInterface
}

//New create a new RatesController
func New(service bankRatingService.RatesServiceInterface) *RatesController {
	return &RatesController{
		RatesService: service,
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
	if r.Currency == nil {
		this.Data["IncorrectCurrency"] = true

		this.Layout = "main_layout.tpl"
		this.TplName = "index.tpl"
		return
	} else if r.Bank == nil {
		this.Data["IncorrectBank"] = true
		this.Layout = "main_layout.tpl"
		this.TplName = "index.tpl"
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
