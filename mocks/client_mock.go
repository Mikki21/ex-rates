package mocks

import (
	"github.com/Mikki21/dlv-project/models"
)

type BankUAClientMock struct {
	Body     []byte
	Unpacked []models.CurrencyBank
	Err      error
}

func (m BankUAClientMock) Get() (body []byte, err error) {
	return m.Body, m.Err
}

func (m BankUAClientMock) GetCurrBank() (unpacked []models.CurrencyBank, err error) {
	return m.Unpacked, m.Err
}
