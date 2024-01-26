package entity

import "time"

type Customers struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password       int    `json:"pin"`
	BirthDate time.Time `json:"birth_date"`
	Address string `json:"address"`
}