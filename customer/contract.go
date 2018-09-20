package customer

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type Contract struct {
	ArchivedDate       string
	ElectricCopy       string
	EndDate            string
	DeadlineForBill    string
	IsArchived         string
	DeadlineForNoitfy  string
	CustomerInfo       string
	Remark             string
	BeginDate          string
	CustomerEmail      string
	Id                 string
	BillCycle          string
	PercentageOfCharge string
	customerInfoQS     []map[string]string
	electricCopyQS     []map[string]string
}
