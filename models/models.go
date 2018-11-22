package models

//MainRequest gives fields for request from main page
type MainRequest struct {
	Currency []string
	Option   string
	Bank     []string
}

//CurrencyBank gives definition of banks
type CurrencyBank struct {
	BankName  string
	CodeAlpha string
	RateBuy   float64 `json:",string"`
	RateSale  float64 `json:",string"`
}

//CurrencyBanks is an array of CurrencyBank
type CurrencyBanks []CurrencyBank
