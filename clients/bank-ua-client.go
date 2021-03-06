package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"

	"github.com/Mikki21/dlv-project/models"
)

var Bm cache.Cache

func init() {
	var err error
	Bm, err = cache.NewCache("memory", `{"interval":1000}`)
	if err != nil {
		fmt.Printf("cache init failed: %v", err)
	}
	_ = Bm
}

// BankUAClient represents a common client to interact with a remote Bank Service
type BankUAClient interface {
	GetCurrBank() (unpacked []models.CurrencyBank, err error)
	GetForIndex() (nbu []models.CurrencyBank, unpacked []models.CurrencyBank, err error)
}

func (bankclient BankUAClientImpl) GetForIndex() (nbu []models.CurrencyBank, unpacked []models.CurrencyBank, err error) {
	set1, err := bankclient.get()
	var set2 []models.IndexClient
	res, err := bankclient.httpClient.Get("https://bank.gov.ua/NBUStatService/v1/statdirectory/exchange?json")
	if err != nil {
		return nil, nil, fmt.Errorf("Get url error: %v", err)
	}

	defer func() {
		if res.Body != nil {
			res.Body.Close()
		}
	}()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("Get url error: %v", err)
	}

	err = json.Unmarshal(body, &set2)
	if err != nil {
		log.Printf("Unmarshal error: %v", err)
		return nil, nil, fmt.Errorf("Error BankUAClient(GetCurrBank):%v", err)
	}
	var a models.CurrencyBank
	for i := range set2 {
		if set2[i].Cc == "USD" {
			a.BankName = "NBU"
			a.CodeAlpha = set2[i].Cc
			a.RateBuy = set2[i].Rate
			a.RateSale = set2[i].Rate
			nbu = append(nbu, a)
		}
	}
	return nbu, set1, nil
}

// Get returns a remote Bank Service response
func (bankClient BankUAClientImpl) get() (banks []models.CurrencyBank, err error) {
	res, err := bankClient.httpClient.Get(bankClient.baseURL)
	if err != nil {
		return nil, fmt.Errorf("Get url error: %v", err)
	}

	defer func() {
		if res.Body != nil {
			res.Body.Close()
		}
	}()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Get url error: %v", err)
	}

	err = json.Unmarshal(body, &banks)
	if err != nil {
		log.Printf("Unmarshal error: %v", err)
		return nil, fmt.Errorf("Error BankUAClient(GetCurrBank):%v", err)
	}
	return
}

// GetCurrBank returns array of structures CurrencyBank after unmarshalling
func (bankClient BankUAClientImpl) GetCurrBank() (unpacked []models.CurrencyBank, err error) {
	if Bm.IsExist("easytrip") {
		return Bm.Get("easytrip").([]models.CurrencyBank), err
	}
	unpacked, err = bankClient.get()
	if err != nil {
		err = fmt.Errorf("GetCurBank get() err: %v", err)
	}
	Bm.Put("easytrip", unpacked, 1000)
	return
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
		baseURL:    beego.AppConfig.String("urlBank"),
		httpClient: &http.Client{},
	}
}
