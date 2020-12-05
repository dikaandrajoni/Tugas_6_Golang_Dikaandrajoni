package common

// Struct API
type Message struct {
	Code   int         `json:"code"`
	Remark string      `json:"remark"`
	Result interface{} `json:"result"`
}

// Start order struct
type Order struct {
	OrderID    string        `json:"orderID"`
	CustomerID string        `json:"customerID,omitempty"`
	EmployeeID string        `json:"employeeID,omitempty"`
	OrderDate  string        `json:"orderDate,omitempty"`
	OrderDet   []OrderDetail `json:"ordersDetail,omitempty"`
}

type OrderDetail struct {
	OrderID     string  `json:"orderID"`
	ProductID   string  `json:"ProductID"`
	ProductName string  `json:"ProductName"`
	UnitPrice   float64 `json:"UnitPrice"`
	Quantity    int     `json:"Quantity"`
}

// End order struct

// Start customer struct
type Customer struct {
	CustomerID   string `json:"customerID"`
	CompanyName  string `json:"companyName,omitempty"`
	ContactName  string `json:"contactName,omitempty"`
	ContactTitle string `json:"contactTitle,omitempty"`
	Address      string `json:"address,omitempty"`
	City         string `json:"city,omitempty"`
	Country      string `json:"country,omitempty"`
	Phone        string `json:"phone,omitempty"`
	PostalCode   string `json:"postalCode,omitempty"`
}

// End customer struct

// Start product struct
type Product struct {
	ProductID       string         `json:"productID"`
	ProductName     string         `json:"productName,omitempty"`
	QuantityPerUnit string         `json:"quantityPerUnit,omitempty"`
	UnitPrice       string         `json:"unitPrice,omitempty"`
	UnitsInStock    string         `json:"unitsInStock,omitempty"`
	UnitsOnOrder    string         `json:"unitsOnOrder,omitempty"`
	SupplierDet     SupplierDetail `json:"supplierDetail,omitempty"`
	CategoryDet     CategoryDetail `json:"categoryDetail,omitempty"`
}

type SupplierDetail struct {
	SupplierID   string `json:"supplierID"`
	CompanyName  string `json:"companyName"`
	ContactName  string `json:"contactName"`
	ContactTitle string `json:"contactTitle"`
	Address      struct {
		Street     string `json:"street"`
		City       string `json:"city"`
		PostalCode string `json:"postalCode"`
		Country    string `json:"country"`
	} `json:"address"`
}

type CategoryDetail struct {
	CategoryID   string `json:"categoryID"`
	CategoryName string `json:"categoryName"`
	Description  string `json:"description"`
}

// End product struct

// Start faspay struct

type RequestFaspay struct {
	Request       string `json:"request"`
	TransactionID string `json:"trx_id"`
	MerchantID    string `json:"merchant_id"`
	BillNumber    string `json:"bill_no"`
	Signature     string `json:"signature"`
}

type ResponseFaspay struct {
	Response          string `json:"response"`
	TransactionID     string `json:"trx_id"`
	MerchantID        string `json:"merchant_id"`
	Merchant          string `json:"merchant"`
	BillNumber        string `json:"bill_not"`
	PaymentReff       string `json:"payment_reff"`
	PaymentDate       string `json:"payment_date"`
	PaymentStatusCode string `json:"payment_status_code"`
	PaymentStatusDesc string `json:"payment_status_desc"`
	ResponseCode      string `json:"response_code"`
	ResponseDesc      string `json:"response_desc"`
}

// End faspay struct

// Start trip struct
type RequestTrip struct {
	DepartureDate1 string `json:"depature_date_1"`
	DepartureDate2 string `json:"depature_date_2"`
	Provinsi       int    `json:"provinsi"`
}

type ResponseTrip struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    []struct {
		TripID           string `json:"TripID"`
		TravelID         string `json:"TravelID"`
		TravelName       string `json:"TravelName"`
		Description      string `json:"Description"`
		Rating           string `json:"Rating"`
		Provinsi         string `json:"Provinsi"`
		CityName         string `json:"CityName"`
		LicenseNumber    string `json:"LicenseNumber"`
		DepartureDate    string `json:"DepartureDate"`
		ReturnDate       string `json:"ReturnDate"`
		Duration         string `json:"Duration"`
		OriginCity       string `json:"OriginCity"`
		AirportName      string `json:"AirportName"`
		Origin           string `json:"Origin"`
		Destination      string `json:"Destination"`
		Transit          string `json:"Transit"`
		DetailTransit    string `json:"DetailTransit"`
		HotelName        string `json:"HotelName"`
		HotelRating      string `json:"HotelRating"`
		Currency         string `json:"Currency"`
		Price            string `json:"Price"`
		PromoCode        string `json:"PromoCode"`
		PromoDescription string `json:"PromoDescription"`
		AirlineName      string `json:"AirlineName"`
		Goods            string `json:"Goods"`
		TermCondition    string `json:"TermCondition"`
		Lat              string `json:"Lat"`
		Long             string `json:"Long"`
		DoubleType       string `json:"DoubleType"`
		TripleType       string `json:"TripleType"`
		QuadType         string `json:"QuadType"`
		Logo             string `json:"Logo"`
	} `json:"data"`
}

// End trip struct

type Result struct {
	Code   int    `json:"code"`
	Remark string `json:"remark,omitempty"`
}

//End Struct API
