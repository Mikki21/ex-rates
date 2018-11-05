package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/astaxie/beego"
)

type UsdController struct {
	beego.Controller
}
type GetCurrency struct {
	BankName  string
	CodeAlpha string
	RateBuy   string
	RateSale  string
}

type GC []GetCurrency

func JsonUnpack(url string) (obj GC) {
	var err error
	res, err := http.Get(url)
	if err != nil {
		log.Printf("Error(http.get) : %v", err)
		return
	}
	temp, err1 := ioutil.ReadAll(res.Body)
	if err1 != nil {
		log.Printf("Error(ReadAll) : %v", err1)
		return
	}

	err = json.Unmarshal(temp, &obj)
	if err != nil {
		log.Printf("Unmarshal error: %v", err)
	}
	return

}

func RequestCatch(a *UsdController) (qm map[string][]string) {
	// Parse the RequestURI
	u, err := url.Parse(a.Ctx.Request.URL.RequestURI())
	if err != nil {
		fmt.Println("Restricted GET: Error parsing URL") // TODO Review possible errors
		return
	}

	// If user tried to access /profile params include query "profile=n"
	//Parse queries looking for key of "profile". If found, update rPage
	qm, _ = url.ParseQuery(u.RawQuery)
	return
}

func getUsd(obj GC) (usdList GC) {
	for i := 0; i < len(obj); i++ {
		if obj[i].CodeAlpha == "USD" {
			usdList = append(usdList, obj[i])
		}
	}
	return
}

func getEur(obj GC) (eurList GC) {
	for i := 0; i < len(obj); i++ {
		if obj[i].CodeAlpha == "EUR" {
			eurList = append(eurList, obj[i])
		}
	}
	return
}

func MinValue(a []float64) (min float64) {
	if len(a) == 0 {
		return 0
	}
	min = a[0]
	for i := range a {
		if a[i] < min {
			min = a[i]
		}
	}
	return
}

func MaxValue(a []float64) (max float64) {
	if len(a) == 0 {
		return 0
	}
	max = a[0]
	for i := range a {
		if a[i] > max {
			max = a[i]
		}
	}
	return
}

func bestBuyPrice(obj GC) (bbBank string, bbRate float64) {
	buy := []float64{}
	for i := 0; i < len(obj); i++ {
		curBuyValue, err := strconv.ParseFloat(obj[i].RateBuy, 64)
		if err != nil {
			log.Printf("eror in parsing float %v", err)
			return
		}
		buy = append(buy, curBuyValue)
	}
	bbRate = MinValue(buy)
	for i := 0; i < len(obj); i++ {
		if check, _ := strconv.ParseFloat(obj[i].RateBuy, 64); check == bbRate {
			bbBank = obj[i].BankName
			break
		}
	}

	return
}

func bestSalePrice(obj GC) (bsBank string, bsRate float64) {
	sale := []float64{}
	for i := 0; i < len(obj); i++ {
		curSaleValue, err := strconv.ParseFloat(obj[i].RateSale, 64)
		if err != nil {
			log.Printf("eror in parsing float %v", err)
			return
		}
		sale = append(sale, curSaleValue)
	}
	bsRate = MaxValue(sale)
	for i := 0; i < len(obj); i++ {
		if check, _ := strconv.ParseFloat(obj[i].RateSale, 64); check == bsRate {
			bsBank = obj[i].BankName
			break
		}
	}
	return
}

func (c *UsdController) RequestWork() {
	qm := RequestCatch(c)

	for _, cur := range qm["currency"] {
		switch cur {
		case "USD":
			if qm["bank"] == nil {
				TakeStruct := JsonUnpack("http://bank-ua.com/export/exchange_rate_cash.json")
				c.Data["BankBuyUsd"], c.Data["BestBuyUsd"] = bestBuyPrice(getUsd(TakeStruct))
				c.Data["BankSaleUsd"], c.Data["BestSaleUsd"] = bestSalePrice(getUsd(TakeStruct))
				c.TplName = "BuySaleUsdOnly.tpl"
			}
		case "EUR":
			if qm["bank"] == nil {
				TakeStruct := JsonUnpack("http://bank-ua.com/export/exchange_rate_cash.json")
				c.Data["BankBuyEur"], c.Data["BestBuyEur"] = bestBuyPrice(getEur(TakeStruct))
				c.Data["BankSaleEur"], c.Data["BestSaleEur"] = bestSalePrice(getEur(TakeStruct))
				c.TplName = "BuySaleEurOnly.tpl"
			}
		}
	}
}
