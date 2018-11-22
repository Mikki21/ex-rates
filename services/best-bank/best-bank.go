package bestBankService

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/oreuta/easytrip/clients"
	"github.com/oreuta/easytrip/models"
)

//BestBankServiceInterface ..
/*
type BestBankServiceInterface interface {
	GetBestBanksSale(r models.MainRequest) (banks []models.CurrencyBank, err error)
	GetBestBanksBuy(r models.MainRequest) (banks []models.CurrencyBank, err error)
}
*/
type BestBankServiceInterface interface {
	GetBestBanks(r models.MainRequest) (bBSale []models.CurrencyBank, bBBuy []models.CurrencyBank, err error)
}

//BestBankService ..
type BestBankService struct {
	Client clients.BankUAClient
}

//New ..
func New(newClient clients.BankUAClient) BestBankServiceInterface {
	return &BestBankService{Client: newClient}
}

func (b BestBankService) GetBestBanks(data models.MainRequest) (bBSale, bBBuy []models.CurrencyBank, err error) {
	jsn, err := b.Client.Get()
	if err != nil {
		return bBSale, bBBuy, fmt.Errorf("Method Get in Client BankUACient: %v", err)
	}
	banks := []models.CurrencyBank{}
	err = json.Unmarshal(jsn, &banks)
	if err != nil {
		return bBSale, bBBuy, fmt.Errorf("json.Unmarshal %v:", err)
	}
	banks = FilterCurrency(data, FilterBank(data, banks))
	if data.Option != "buy" {
		bBSale = BestSale(banks)
	}
	if data.Option != "sale" {
		bBBuy = BestBuy(banks)
	}

	return bBSale, bBBuy, nil
}

func FilterBank(data models.MainRequest, inpBanks []models.CurrencyBank) (OutpBanks []models.CurrencyBank) {
	if len(data.Bank) == 0 {
		return inpBanks
	}
	s := strings.Join(data.Bank, "")
	if strings.Contains(s, "privat") {
		for _, value := range inpBanks {
			if value.BankName == "\u041f\u0440\u0438\u0432\u0430\u0442\u0411\u0430\u043d\u043a" {
				OutpBanks = append(OutpBanks, value)
			}
		}
	}
	if strings.Contains(s, "pireus") {
		for _, value := range inpBanks {
			if value.BankName == "\u041f\u0456\u0440\u0435\u0443\u0441 \u0411\u0430\u043d\u043a" {
				OutpBanks = append(OutpBanks, value)
			}
		}
	}
	if strings.Contains(s, "otp") {
		for _, value := range inpBanks {
			if value.BankName == "\u041e\u0422\u041f \u0411\u0430\u043d\u043a" {
				OutpBanks = append(OutpBanks, value)
			}
		}
	}
	return OutpBanks
}
func FilterCurrency(data models.MainRequest, inpBanks []models.CurrencyBank) (OutpBanks []models.CurrencyBank) {
	s := strings.Join(data.Currency, "")

	if strings.Contains(s, "usd") || len(data.Currency) == 0 {
		for _, value := range inpBanks {
			if value.CodeAlpha == "USD" {
				OutpBanks = append(OutpBanks, value)
			}
		}
	}

	if strings.Contains(s, "eur") || len(data.Currency) == 0 {
		for _, value := range inpBanks {
			if value.CodeAlpha == "EUR" {
				OutpBanks = append(OutpBanks, value)
			}
		}
	}
	return
}

//min
func BestSale(inpBanks []models.CurrencyBank) (OutpBanks []models.CurrencyBank) {
	banks := append([]models.CurrencyBank(nil), inpBanks...)
	sort.Slice(banks, func(i, j int) bool {
		return banks[i].RateSale < banks[j].RateSale
	})
	var eur float64
	var usd float64
	for _, value := range banks {
		if value.CodeAlpha == "EUR" {
			eur = value.RateSale
			break
		}
	}
	for _, value := range banks {
		if value.CodeAlpha == "USD" {
			usd = value.RateSale
			break
		}
	}
	for _, value := range banks {
		if (value.CodeAlpha == "EUR" && value.RateSale == eur) || (value.CodeAlpha == "USD" && value.RateSale == usd) {
			OutpBanks = append(OutpBanks, value)
		}
	}
	return OutpBanks
}

//max
func BestBuy(inpBanks []models.CurrencyBank) (OutpBanks []models.CurrencyBank) {
	banks := append([]models.CurrencyBank(nil), inpBanks...)
	sort.Slice(banks, func(i, j int) bool {
		return banks[i].RateBuy > banks[j].RateBuy
	})
	var eur float64
	var usd float64
	for _, value := range banks {
		if value.CodeAlpha == "EUR" {
			eur = value.RateBuy
			break
		}
	}
	for _, value := range banks {
		if value.CodeAlpha == "USD" {
			usd = value.RateBuy
			break
		}
	}
	for _, value := range banks {
		if (value.CodeAlpha == "EUR" && value.RateBuy == eur) || (value.CodeAlpha == "USD" && value.RateBuy == usd) {
			OutpBanks = append(OutpBanks, value)
		}
	}
	return OutpBanks
}
