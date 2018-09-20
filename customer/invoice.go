package customer

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Invoice struct {
	TaxRate            float64
	InvoiceContent     string
	IsInvoice          string
	TaxpayerIdentifNum string
	CompanyName        string
	Remark             string
	Account            string
	OpeningBank        string
}

func CreateInvoice(conn *sql.DB, userId int64) (Invoice, error) {
	sqlStr := `select tax_rate, invoice_content, 
		need_invoice, taxpayer_identif_num, company_name, remark, account, 
		opening_bank from invoice where user_id=?`
	row := conn.QueryRow(sqlStr, userId)

	var invoice Invoice
	err := row.Scan(&invoice.TaxRate, &invoice.InvoiceContent, &invoice.IsInvoice, &invoice.TaxpayerIdentifNum,
		&invoice.CompanyName, &invoice.Remark, &invoice.Account, &invoice.OpeningBank)

	if err != nil {
		fmt.Println("fail to query invoice, userId=%d", userId)
	}

	return invoice, err
}
