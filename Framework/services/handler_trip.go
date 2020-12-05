package services

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	cm "Tugas_6_Golang_Dikaandrajoni/Framework/common"

	_ "github.com/go-sql-driver/mysql"
)

func (PaymentService) TripHandler(ctx context.Context, req cm.RequestTrip) (res cm.ResponseTrip) {
	defer panicRecovery()

	// request to server
	requestTrip := &cm.RequestTrip{
		DepartureDate1: req.DepartureDate1,
		DepartureDate2: req.DepartureDate2,
		Provinsi:       req.Provinsi,
	}

	reqBody, err := json.Marshal(requestTrip)
	if err != nil {
		panic(err.Error())
	}

	resp, err := http.Post("http://35.186.147.192/travel/GetTripsSample.php", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	var response cm.ResponseTrip
	json.Unmarshal(resBody, &response)

	res.Status = response.Status
	res.Message = response.Message
	res.Data = response.Data

	// Insert into database
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

	for _, current := range response.Data {
		travelID := current.TravelID
		travelName := current.TravelName
		airportName := current.AirportName
		hotelName := current.HotelName
		description := current.Description

		sql := "INSERT INTO trip (travel_id, travel_name, airport_name, hotel_name, description) VALUES(?,?,?,?,?)"
		stmt, err := db.Prepare(sql)
		if err != nil {
			panic(err.Error())
		}

		_, err = stmt.Exec(travelID, travelName, airportName, hotelName, description)
	}
	return
}
