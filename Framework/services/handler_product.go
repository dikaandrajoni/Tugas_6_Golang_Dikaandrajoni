package services

import (
	"context"
	"database/sql"
	"fmt"

	cm "Tugas_6_Golang_Dikaandrajoni/Framework/common"

	_ "github.com/go-sql-driver/mysql"
)

func (PaymentService) ProductHandler(ctx context.Context, req cm.Product) (res cm.Message) {
	defer panicRecovery()

	host := cm.Config.Connection.Host
	port := cm.Config.Connection.Port
	user := cm.Config.Connection.User
	pass := cm.Config.Connection.Password
	data := cm.Config.Connection.Database

	var mySQL = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, pass, host, port, data)

	db, err := sql.Open("mysql", mySQL)

	if err != nil {
		panic(err.Error())
	}

	var product cm.Product
	var supplierDet cm.SupplierDetail
	var categoryDet cm.CategoryDetail

	sql := `SELECT
				ProductID,
				IFNULL(ProductName,'') ProductName,
				IFNULL(SupplierID,'') SupplierID,
				IFNULL(CategoryID,'') CategoryID,
				IFNULL(QuantityPerUnit,'') QuantityPerUnit,
				IFNULL(UnitPrice,'') UnitPrice,
				IFNULL(UnitsInStock,'') UnitsInStock,
				IFNULL(UnitsOnOrder,'') UnitsOnOrder
			FROM products WHERE ProductID = ?`
	result, err := db.Query(sql, req.ProductID)
	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err = result.Scan(
			&product.ProductID, &product.ProductName, &supplierDet.SupplierID,
			&categoryDet.CategoryID, &product.QuantityPerUnit, &product.UnitPrice,
			&product.UnitsInStock, &product.UnitsOnOrder,
		)
		if err != nil {
			panic(err.Error())
		}

		// get supplier detail
		sqlSupplier := `SELECT
							SupplierID,
							IFNULL(CompanyName,'') CompanyName,
							IFNULL(ContactName,'') ContactName,
							IFNULL(ContactTitle,'') ContactTitle,
							IFNULL(Address,'') Address,
							IFNULL(City,'') City,
							IFNULL(PostalCode,'') PostalCode,
							IFNULL(Country,'') Country
						FROM suppliers WHERE SupplierID = ?`
		resultSupplier, errSupplier := db.Query(sqlSupplier, supplierDet.SupplierID)
		defer resultSupplier.Close()

		if errSupplier != nil {
			panic(errSupplier.Error())
		}

		for resultSupplier.Next() {
			errSupplier = resultSupplier.Scan(
				&supplierDet.SupplierID, &supplierDet.CompanyName, &supplierDet.ContactName,
				&supplierDet.ContactTitle, &supplierDet.Address.Street, &supplierDet.Address.City,
				&supplierDet.Address.PostalCode, &supplierDet.Address.Country,
			)

			if errSupplier != nil {
				panic(errSupplier.Error())
			}

			// attach to product struct
			product.SupplierDet = supplierDet
		}

		// get category detail
		sqlCategory := `SELECT
							CategoryID,
							IFNULL(CategoryName, '') CategoryName,
							IFNULL(Description, '') Description
						FROM categories WHERE CategoryID = ?`
		resultCategory, errCategory := db.Query(sqlCategory, categoryDet.CategoryID)
		defer resultCategory.Close()

		if errCategory != nil {
			panic(errCategory.Error())
		}

		for resultCategory.Next() {
			errCategory = resultCategory.Scan(
				&categoryDet.CategoryID, &categoryDet.CategoryName, &categoryDet.Description,
			)

			// attacth to product struct
			product.CategoryDet = categoryDet
		}
	}

	if &product != nil {
		res.Code = 100
		res.Remark = "Success"
	}

	res.Result = product

	return
}
