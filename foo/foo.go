package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "b44aea942d06ef:ca7587cd@tcp(us-cdbr-iron-east-01.cleardb.net:3306)/heroku_c1b06006d977072")
	err = db.Ping()
	if err != nil {
		fmt.Errorf("error %v", err)
	}
	defer db.Close()

	insert, err := db.Query("insert into Price_info values ('USD','28.01','Best buy','Приват Банк')")

	if err != nil {
		fmt.Errorf("error %v", err)
	}
	defer insert.Close()

}
