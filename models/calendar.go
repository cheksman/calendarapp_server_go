package models

type Address struct {
	Country string
	State string
	City string
	Street string
	Zipcode string
}

type Calendar struct {
	Date string
	Start string
	End string
	Description string
	Link string
	Address Address
}
