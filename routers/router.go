package routers

import (
	"github.com/Mikki21/dlv-project/clients"
	"github.com/Mikki21/dlv-project/controllers/bank-rating"
	"github.com/Mikki21/dlv-project/controllers/best-bank"
	"github.com/Mikki21/dlv-project/controllers/index-service"
	"github.com/Mikki21/dlv-project/controllers/statistics"
	"github.com/Mikki21/dlv-project/services/bank-rating"
	"github.com/Mikki21/dlv-project/services/best-bank"
	"github.com/Mikki21/dlv-project/services/index-serv"
	"github.com/astaxie/beego"
)

func init() {
	ratesclient := clients.New()

	ratesService := bankRatingService.New(ratesclient)
	ratesController := bankRatingController.New(ratesService)

	bestService := bestBankService.New(ratesclient)
	bestController := bestBankController.New(bestService)

	indexService := indexService.New(ratesclient)
	indexController := indexServiceController.New(indexService)

	beego.Router("/", indexController)
	beego.Router("/comparision", ratesController)
	beego.Router("/best", bestController)
	beego.Router("/today", indexController)
	beego.Router("/statistics", &statistics.StatisticController{})
}
