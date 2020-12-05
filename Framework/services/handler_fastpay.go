package services

import (
	"context"
	"database/sql"
	"fmt"
	cm "Tugas_6_Golang_Dikaandrajoni/Framework/common"
	"net/http"
	"strconv"
)

func (PaymentService) FaspayHandler(ctx context.Context, req cm.RequestFaspay) (res cm.ResponseFaspay) {
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

	var transactionID = req.TransactionID
	var responseFaspay cm.ResponseFaspay

	sql := `SELECT
				TransactionID,
				IFNULL(MerchantID,'') MerchantID,
				IFNULL(Merchant,'') Merchant,
				IFNULL(BillNumber,'') BillNumber,
				IFNULL(PaymentReff,'') PaymentReff,
				IFNULL(PaymentDate,'') PaymentDate,
				IFNULL(PaymentStatusCode,'') PaymentStatusCode,
				IFNULL(PaymentStatusDesc,'') PaymentStatusDesc
			FROM transactions WHERE TransactionID = ?`

	result, err := db.Query(sql, transactionID)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	if result.Next() {
		err := result.Scan(
			&responseFaspay.TransactionID, &responseFaspay.MerchantID, &responseFaspay.Merchant,
			&responseFaspay.BillNumber, &responseFaspay.PaymentReff, &responseFaspay.PaymentDate,
			&responseFaspay.PaymentStatusCode, &responseFaspay.PaymentStatusDesc,
		)

		if err != nil {
			panic(err.Error())
		}
	}

	if responseFaspay.TransactionID != "" {
		res = responseFaspay
		res.Response = req.Request
		res.ResponseCode = strconv.Itoa(http.StatusOK)
		res.ResponseDesc = "Sukses ambil data"
	} else {
		res.ResponseCode = strconv.Itoa(http.StatusNotFound)
		res.ResponseDesc = "Gagal ambil data"
	}

	return
}
