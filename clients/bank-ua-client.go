package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/oreuta/easytrip/models"
)

// BankUAClient represents a common client to interact with a remote Bank Service
type BankUAClient interface {
	Get() (body []byte, err error)
	GetCurrBank() (unpacked []models.CurrencyBank, err error)
}

// Get returns a remote Bank Service response
func (bankClient BankUAClientImpl) Get() (body []byte, err error) {
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

// GetCurrBank returns array of structures CurrencyBank after unmarshalling
func (bankClient BankUAClientImpl) GetCurrBank() (unpacked []models.CurrencyBank, err error) {
	body, err := bankClient.Get()
	if err != nil {
		return nil, fmt.Errorf("Get url error: %v", err)
	}
	err = json.Unmarshal(body, &unpacked)
	if err != nil {
		log.Printf("Unmarshal error: %v", err)
		return nil, fmt.Errorf("Error BankUAClient(GetCurrBank):%v", err)
	}
	return unpacked, err
}

// BankUAClientImpl implements BankUAClient interface
type BankUAClientImpl struct {
	baseURL    string
	httpClient *http.Client
}

// New creates a new BankUAClient instance
//*realize throu interface
func New() BankUAClientImpl {
	return BankUAClientImpl{
		baseURL:    "http://bank-ua.com/export/exchange_rate_cash.json",
		httpClient: &http.Client{},
	}
}
