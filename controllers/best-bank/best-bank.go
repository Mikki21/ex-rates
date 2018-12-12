package bestBankController

import (
	"time"

	"github.com/Mikki21/dlv-project/models"
	"github.com/Mikki21/dlv-project/services/best-bank"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

type bestBankController struct {
	beego.Controller
	BestService bestBankService.BestBankServiceInterface
}

func New(s bestBankService.BestBankServiceInterface) *bestBankController {
	return &bestBankController{BestService: s}
}

func (r *bestBankController) Get() {
	toolbox.StatisticsMap.AddStatistics("GET", "/best", "&controllers.bestBankController.bestBankController", time.Duration(15000))
	inpData := models.MainRequest{
		Currency: r.GetStrings("currency"),
		Option:   r.GetString("option"),
		Bank:     r.GetStrings("bank"),
	}

	{
		i := 0
		if inpData.Currency == nil {
			r.Data["IncorrectCurrency"] = true
			i++
		}
		if inpData.Bank == nil {
			r.Data["IncorrectBank"] = true
			i++
		}
		if i > 0 {

			r.Layout = "main_layout.tpl"
			r.TplName = "index.tpl"
			return
		}
	}

	sale, buy, err := r.BestService.GetBestBanks(inpData)
	if err != nil {
		beego.Error("GetBestBanks func in BestService: %v", err)
		return
	}
	r.Layout = "bestBank_layout.tpl"
	r.TplName = "bestBank.tpl"
	r.Data["Buy"] = buy
	r.Data["TitleBuy"] = "No Banks for buying"
	r.Data["Sale"] = sale
	r.Data["TitleSale"] = "No Banks for saleing"
	if buy != nil {
		r.Data["TitleBuy"] = "Best Buy"
	}
	if sale != nil {
		r.Data["TitleSale"] = "Best Sale"
	}
}
