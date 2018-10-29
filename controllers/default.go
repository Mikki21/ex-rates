package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type GetCurrency struct {
	BankName  string
	CodeAlpha string
	RateBuy   string
	RateSale  string
}

type GC []GetCurrency

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (main *MainController) WelcomePage() {
	main.Data["MarketSite"] = "Poloniex.com"
	res, err := http.Get("http://bank-ua.com/export/exchange_rate_cash.json")
	if err != nil {
		log.Printf("Error(http.get) : %v", err)
		return
	}
	temp, err1 := ioutil.ReadAll(res.Body)
	if err1 != nil {
		log.Printf("Error(ReadAll) : %v", err1)
		return
	}

	var TakeStruct GC

	errJson := json.Unmarshal(temp, &TakeStruct)
	if errJson != nil {
		log.Printf("Unmarshal error: %v", errJson)
	}

	var structNum, structNum1 int
	var resultBuy, resultSale float64 = 1000, 0
	for i := 0; i < len(TakeStruct); i++ {
		if TakeStruct[i].CodeAlpha == "USD" {
			curBuyValue, errcheck := strconv.ParseFloat(TakeStruct[i].RateBuy, 64)
			if errcheck != nil {
				log.Printf("eror in parsing float %v", errcheck)
				return
			}
			fmt.Printf("cur buy value : %v\n", curBuyValue)

			curSaleValue, errcheck := strconv.ParseFloat(TakeStruct[i].RateSale, 64)
			if errcheck != nil {
				log.Printf("eror in parsing float %v", errcheck)
				return
			}
			fmt.Printf("cur sale value : %v\n", curSaleValue)

			if curBuyValue < resultBuy {
				resultBuy, structNum = curBuyValue, i
			}

			if resultSale < curSaleValue {
				resultSale, structNum1 = curSaleValue, i
			}
		}
	}

	fmt.Printf("FINAL buy value : %v\n", resultBuy)
	fmt.Printf("FINAL sale value : %v\n", resultSale)

	main.Data["USD_BUY"] = resultBuy
	main.Data["USD_SALE"] = resultSale
	main.Data["Bank_info"] = TakeStruct[structNum].BankName
	main.Data["Bank_info1"] = TakeStruct[structNum1].BankName

	main.TplName = "main.tpl"
}
