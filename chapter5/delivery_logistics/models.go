package main

type Delivery struct {
	Sender   Sender
	Receiver Receiver
	Package  Package
	Payment  Payment
	Carrier  Carrier
}

type Sender struct {
	ID        interface{}
	FirstName string
	LastName  string
	Address   Address
	Phone     string
}

type Receiver struct {
	ID        interface{}
	FirstName string
	LastName  string
	Address   Address
	Phone     string
}

type Package struct {
	ID         interface{}
	Dimensions Dimension
	Weight     int
	IsDamaged  bool
	Status     string
}

type Payment struct {
	ID             interface{}
	InitiatedOn    string
	SuccessfulOn   string
	MerchantId     int
	PaymentDetails PaymentDetail
}

type Carrier struct {
	ID          interface{}
	Name        string
	CarrierCode int
	IsPartner   bool
}

type Shipment struct {
	ID         interface{}
	Sender     interface{}
	Receiver   interface{}
	Package    interface{}
	Payment    interface{}
	Carrier    interface{}
	PromisedOn string
}

type Address struct {
	Type    string
	Street  string
	City    string
	State   string
	PinCode int
	Country string
}

type Dimension struct {
	Width  int
	Height int
}

type PaymentDetail struct {
	TransactionToken string
}
