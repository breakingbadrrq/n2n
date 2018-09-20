package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"lightningdog-app.com/customer"
)

func main() {
	fmt.Println("begin to move customer data to new table")
	conn, err := sql.Open("mysql", customer.CommonDBPath)

	rows, err := conn.Query("select user_id, is_visible  from customer")
	if err != nil {
		fmt.Println("fail to connect mysql")
		return
	}

	defer rows.Close()
	defer conn.Close()

	for rows.Next() {
		var userId int64
		var checkState int8

		rows.Scan(&userId, &checkState)
		if checkState == 1 {
			invoice, err1 := load(conn, userId)
			if err1 != nil {
				fmt.Println("error1=%s", err1.Error())
				continue
			}

			invoiceStr, err2 := json.Marshal(invoice)
			if err2 != nil {
				fmt.Println("error2=%s", err2.Error())
				continue
			}

			fmt.Printf("userId=%d, checkState=%d, invoice=%s", userId, checkState, invoiceStr)
		}
	}
}

func load(conn *sql.DB, userId int64) (customer.Invoice, error) {
	return customer.CreateInvoice(conn, userId)
}
