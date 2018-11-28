package sql1

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/oreuta/easytrip/clients"
	"github.com/oreuta/easytrip/models"
)

var Db *sql.DB

func init() {
	Db = CreateConnect("b9ebde93dc372a:191521fe@tcp(us-cdbr-iron-east-01.cleardb.net:3306)/heroku_6566a8fa9264fc0")
	_ = Db
}

//CreateConnect open connection to MySQL database
func CreateConnect(connection string) (Db *sql.DB) {
	var err error
	Db, err = sql.Open("mysql", connection)
	if err != nil {
		fmt.Errorf("Connection open error: %v", err)
	}
	return
}

//Update updates information from BankUAclient to database
func Update(Db *sql.DB) error {
	a := clients.New()
	var err error
	var res []models.CurrencyBank
	res, err = a.GetCurrBank()
	if err != nil {
		fmt.Printf("err for update sql is : %v", err)
	}
	for {
		for i := range res {
			_, err := Db.Query("update BanksList set RateBuy=? RateSale=? where BankName=? and CodeAlpha=?", res[i].RateBuy, res[i].RateSale, res[i].BankName, res[i].CodeAlpha)
			if err != nil {
				return fmt.Errorf("Number[%v] cant update to database: %v", res[i], err)

			}
		}
		time.Sleep(3 * time.Second)
	}

}

//JsnChanger creates list of banks from database
func JsnChanger(Db *sql.DB) (res []models.CurrencyBank, err error) {
	rows, err := Db.Query("select * from BanksList")
	if err != nil {
		return nil, fmt.Errorf("Select query failed:%v", err)
	}
	var a models.CurrencyBank
	for rows.Next() {
		err := rows.Scan(&a.BankName, &a.CodeAlpha, &a.RateBuy, &a.RateSale)
		if err != nil {
			return nil, fmt.Errorf("Scan error:%v", err)
		}
		res = append(res, a)
	}
	return
}
