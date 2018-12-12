package bestBankController

import (
	"time"

	"github.com/Mikki21/dlv-project/models"
	"github.com/Mikki21/dlv-project/services/best-bank"
	"github.com/Mikki21/dlv-project/services/index-serv"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

type bestBankController struct {
	beego.Controller
	BestService bestBankService.BestBankServiceInterface
	Indexserv   indexService.IndexServiceInterface
}

func New(s bestBankService.BestBankServiceInterface, s1 indexService.IndexServiceInterface) *bestBankController {
	return &bestBankController{
		BestService: s,
		Indexserv:   s1,
	}
}

func (r *bestBankController) Get() {
	toolbox.StatisticsMap.AddStatistics("GET", "/best", "&controllers.bestBankController.bestBankController", time.Duration(15000))
	inpData := models.MainRequest{
		Currency: r.GetStrings("currency"),
		Option:   r.GetString("option"),
		Bank:     r.GetStrings("bank"),
	}
	ind, _ := r.Indexserv.IndexGet()
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
			r.Data["MainNumber"] = r.Indexserv.ToFixed((ind[0].RateBuy+ind[1].RateBuy)/2, 2)
			r.Data["NBU"] = ind[0]
			r.Data["Others"] = ind[1]
			r.Layout = "main_layout.tpl"
			r.TplName = "index.tpl"
			r.LayoutSections = make(map[string]string)
			r.LayoutSections["Second"] = "today.tpl"
			r.LayoutSections["Scroll"] = "scrl.tpl"
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
