package customer

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type company struct {
	Id                 int64
	SaleAdminId        int64
	Address            string
	ServicePhone       string
	City               string
	Remark             string
	Shortname          string
	FinanceScope       string
	FinanceMoney       float64
	Number             string
	BusinessLicenses   string
	Province           string
	Phone              string
	IsQuoted           int8
	Name               string
	Position           string
	Contactor          string
	BusinessLicensesQS []map[string]string
}

func CreateCompany(conn *sql.DB, userId int64) (Company, error) {
	sqlStr := `SELECT
            c.company_fname,
            c.registration_num,
            c.financing_money,
            c.financing_trun,
            c.business_licence_pic,
            ci.head_office_addr,
            ci.company_sname,
            ci.customer_service_phone,
            ci.contract_name,
            ci.contract_phone,
            ci.contract_position,
            ci.province,
            ci.remark
        FROM
            company c,
            company_info ci
        WHERE
            c.id = ci.company_id
            AND ci.user_id = ?`

	row := conn.QueryRow(sqlStr, userId)
	var company Company
	var licenseNames string
	err := row.Scan(&company.Name, &company.Number, &company.FinanceMoney, &company.FinanceScope,
		&licenseNames, &company.Address, &company.Shortname, &company.ServicePhone, &company.Contractor,
		&company.Phone, &company.Position, &company.Province, &company.Remark)
	if err != nil {
		fmt.Printf("fail to get companyInfo, userId=%d", userId)
		return nil, err
	}

	company.BusinessLicensesQS = produceLicensesMap(licenseNames)
	return company, err
}

func produceLicensesMap(licenseNamesStr) []map[string]string {
	var resultMaps = make([]map[string]string)
	licenseNames := strings.Split(licenseNamesStr, ',')
	for index, value := range licenseNames {
		licenseMap := make(map[string]string)
		licenseMap["name"] = value
		licenseMap["url"] = LicenseMapUrl + name
		append(resultMaps, licenseMap)
	}

	return resultMaps
}

func (company Company) Store(conn *sql.DB) error {
	sqlStr := `INSERT INTO company (id, sale_admin_id, name, number, is_quoted, finance_scope, 
		bussiness_licenses, address, service_phone, contractor, phone, position, city, shortname, 
		remark) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	licenses := getLicenses(company.BusinessLicensesQS)
	result, err := conn.Exec(sqlStr, company.Id, company.SaleAdminId, company.Name, company.Number,
		company.IsQuoted, company.financeScope, licenses, company.Address, company.ServicePhone,
		company.Contractor, company.Phone, company.Position, company.City, company.ShortName, company.Remark)
	if err != nil {
		panic("fail to store compay")
	}
	return err
}

func getLicenses(licenses []map[String]String) string {
	licensesStr := ""

	if len(licenses) == 0 {
		return licensesStr
	}

	for index, objMap := range licenses {
		for key, value := range objMap {
			if key == "name" {
				licensesStr += value + ","
			}
		}
	}

	return licensesStr[0 : len(licensesStr)-1]
}
