package sql1

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Mikki21/dlv-project/clients"
	"github.com/Mikki21/dlv-project/models"
)

var db *sql.DB

func createConnect(connection string) (Db *sql.DB) {
	var err error
	Db, err = sql.Open("mysql", connection)
	if err != nil {
		fmt.Errorf("Connection open error: %v", err)
	}
	return
}

//Update updates information from BankUAclient to database
func Update() {
	db = createConnect(models.Base())
	for {

		a := clients.New()
		var err error
		res, err := a.GetCurrBank()
		if err != nil {
			fmt.Printf("err for update sql is : %v", err)
		}
		for i := range res {
			_, err := db.Query("update BanksList set RateBuy=?,RateSale=? where BankName=? and CodeAlpha=?", res[i].RateBuy, res[i].RateSale, res[i].BankName, res[i].CodeAlpha)
			if err != nil {
				log.Println("err of query")
			}
		}
		// db.Close()
		time.Sleep(100 * time.Second)

	}
}

//JsnChanger creates list of banks from database
func JsnChanger() (res []models.CurrencyBank, err error) {

	db = createConnect(models.Base())
	rows, err := db.Query("select * from BanksList")
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
	db.Close()
	return
}
