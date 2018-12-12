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

type IndexClient struct {
	Cc   string
	Rate float64
}

//CurrencyBanks is an array of CurrencyBank
type CurrencyBanks []CurrencyBank

func Bank() map[string]string {
	banksMap := map[string]string{
		"privat": "ПриватБанк",
		"otp":    "ОТП Банк",
		"pireus": "Піреус Банк",
		"kredo":  "Кредобанк",
	}
	return banksMap
}

func Currency() map[string]string {
	currencyMap := map[string]string{
		"usd": "USD",
		"eur": "EUR",
	}
	return currencyMap
}

func Base() string {
	return "bbdc4d9aa08941:ca3c3019@tcp(us-cdbr-iron-east-01.cleardb.net:3306)/heroku_d1f744b3705e71b"
}
