package model

type User struct {
	ID           int    `json:"user_id"`
	UserName     string `json:"user_name"`
	Password     string `json:"password"`
	AadharNumber int    `json:"aadhar_number"`
}
