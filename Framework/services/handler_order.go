package services

import (
	"context"
	"database/sql"
	"fmt"

	cm "Tugas_6_Golang_Dikaandrajoni/Framework/common"

	_ "github.com/go-sql-driver/mysql"
)

func (PaymentService) OrderHandler(ctx context.Context, req cm.Order) (res cm.Message) {
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

	var order cm.Order
	var orderDet cm.OrderDetail

	sql := `SELECT
				OrderID,
				IFNULL(CustomerID,'') CustomerID,
				IFNULL(EmployeeID,'') EmployeeID,
				IFNULL(OrderDate,'') OrderDate
			FROM orders WHERE OrderID = ?`

	result, err := db.Query(sql, req.OrderID)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&order.OrderID, &order.CustomerID, &order.EmployeeID, &order.OrderDate)

		if err != nil {
			panic(err.Error())
		}

		sqlDetail := `SELECT
						order_details.OrderID
						, products.ProductID
						, products.ProductName
						, order_details.UnitPrice
						, order_details.Quantity
					FROM
						order_details
						INNER JOIN products
							ON (order_details.ProductID = products.ProductID)
					WHERE order_details.OrderID	= ?`

		orderID := &order.OrderID
		resultDetail, errDet := db.Query(sqlDetail, *orderID)

		defer resultDetail.Close()

		if errDet != nil {
			panic(err.Error())
		}

		for resultDetail.Next() {

			err := resultDetail.Scan(&orderDet.OrderID, &orderDet.ProductID, &orderDet.ProductName, &orderDet.UnitPrice, &orderDet.Quantity)

			if err != nil {
				panic(err.Error())
			}

			order.OrderDet = append(order.OrderDet, orderDet)

		}

	}
	if &order != nil {
		res.Code = 100
		res.Remark = "Success"
	}

	res.Result = order

	return
}
