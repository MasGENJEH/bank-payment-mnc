package entity

import "time"

type Customers struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Pin       int    `json:"pin"`
	BirthDate time.Time `json:"birth_date"`
	Address string `json:"address"`
	Contact string `json:"contact"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}