package bankRatingController

import (
	"testing"

	"github.com/oreuta/easytrip/mocks"
	"github.com/oreuta/easytrip/models"
)

func TestGet(t *testing.T) {
	banks := []models.CurrencyBank{
		{
			BankName:  "ПриватБанк",
			CodeAlpha: "USD",
			RateBuy:   66.6,
			RateSale:  33.3,
		},
		{
			BankName:  "Піреус Банк",
			CodeAlpha: "USD",
			RateBuy:   229,
			RateSale:  22,
		},
	}

	mserv := mocks.RatesServiceMock{
		Banks: banks,
	}
	ratesController := New(mserv)

	ratesController.Get()
}
