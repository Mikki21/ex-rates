package bestBankService

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/oreuta/easytrip/models"
)

type BankUAClientMock interface {
	Get() (body []byte, err error)
	GetCurrBank() (unpacked []models.CurrencyBank, err error)
}

type BankUAClientImplMock struct {
	baseURL    string
	httpClient *http.Client
}

func (bankClient BankUAClientImplMock) Get() (body []byte, err error) {
	res, err := bankClient.httpClient.Get(bankClient.baseURL)
	if err != nil {
		return nil, fmt.Errorf("Get url error: %v", err)
	}

	defer func() {
		if res.Body != nil {
			res.Body.Close()
		}
	}()

	return ioutil.ReadAll(res.Body)
}
