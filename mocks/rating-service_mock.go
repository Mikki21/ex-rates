package mocks

import (
	"github.com/Mikki21/dlv-project/models"
)

type RatesServiceMock struct {
	Banks []models.CurrencyBank
	Err   error
}

func (m RatesServiceMock) GetBankRates(r models.MainRequest) (banks []models.CurrencyBank, err error) {
	return m.Banks, m.Err
}
