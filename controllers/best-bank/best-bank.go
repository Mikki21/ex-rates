package bestBankController

import (
	"github.com/astaxie/beego"
	"github.com/oreuta/easytrip/models"
	"github.com/oreuta/easytrip/services/best-bank"
)

type bestBankController struct {
	beego.Controller
	BestService bestBankService.BestBankServiceInterface
}

func New(s bestBankService.BestBankServiceInterface) *bestBankController {
	return &bestBankController{BestService: s}
}

func (r *bestBankController) Get() {
	inpData := models.MainRequest{
		Currency: r.GetStrings("currency"),
		Option:   r.GetString("option"),
		Bank:     r.GetStrings("bank"),
	}

	Sale, Buy, err := r.BestService.GetBestBanks(inpData)
	if err != nil {
		beego.Error("GetBestBanks func in BestService: %v", err)
		return
	}

	if inpData.Currency == nil || inpData.Bank == nil {
		r.Data["error"] = "Select currency and banks"
		r.TplName = "error.tpl"
		return
	}
	if inpData.Option == "sale" {
		r.Data["Sale"] = Sale
		r.TplName = "bestSale.tpl"
	}
	if inpData.Option == "buy" {
		r.Data["Buy"] = Buy
		r.TplName = "bestBuy.tpl"
	}
	if inpData.Option == "both" {
		r.Data["Buy"] = Buy
		r.Data["Sale"] = Sale
		r.TplName = "both.tpl"
	}
}
