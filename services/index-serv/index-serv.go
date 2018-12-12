package indexService

import (
	"math"

	"github.com/Mikki21/dlv-project/clients"
	"github.com/Mikki21/dlv-project/models"
)

//RatesServiceInterface represents a common service to interact with BankUAClient
type IndexServiceInterface interface {
	IndexGet() (banks []models.CurrencyBank, err error)
	ToFixed(num float64, precision int) float64
}

//GetBankRates returns list of Banks response
func (obj *IndexService) IndexGet() (banks []models.CurrencyBank, err error) {
	banks, b, err := obj.Client.GetForIndex()
	var totalBuy, totalSale float64
	for i := range b {
		totalBuy += b[i].RateBuy
		totalSale += b[i].RateSale
	}

	banks[0].RateBuy = obj.ToFixed(banks[0].RateBuy, 2)
	banks[0].RateSale = obj.ToFixed(banks[0].RateSale, 2)

	q1 := obj.ToFixed(totalSale/float64(len(b)), 2)
	q2 := obj.ToFixed(totalBuy/float64(len(b)), 2)
	i := models.CurrencyBank{
		BankName:  "Банки",
		CodeAlpha: "USD",
		RateBuy:   q2,
		RateSale:  q1,
	}
	banks = append(banks, i)
	return
}

//BankRatingService implements RatesServiceInterface interface
type IndexService struct {
	Client clients.BankUAClient //Сlient interface
}

//New creates a new RatesService instance
func New(newClient clients.BankUAClient) IndexServiceInterface {
	return &IndexService{
		Client: newClient,
	}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func (obj IndexService) ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
