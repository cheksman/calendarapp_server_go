package models

type User struct {
	Fullname string
	Email string
	Password string
	Description string
	ProfileImage string
	Address Address
	Calendar []Calendar
}