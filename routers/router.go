package routers

import (
	"github.com/astaxie/beego"
	"github.com/oreuta/easytrip/clients"
	"github.com/oreuta/easytrip/controllers"
	"github.com/oreuta/easytrip/controllers/bank-rating"
	"github.com/oreuta/easytrip/controllers/best-bank"
	"github.com/oreuta/easytrip/services/bank-rating"
	"github.com/oreuta/easytrip/services/best-bank"
)

func init() {
	ratesclient := clients.New()
	ratesService := bankRatingService.New(ratesclient)
	ratesController := bankRatingController.New(ratesService)
	bestService := bestBankService.New(ratesclient)
	bestController := bestBankController.New(bestService)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/comparision", ratesController)
	beego.Router("/best", bestController)

}
