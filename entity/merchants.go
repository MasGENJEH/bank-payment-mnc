package entity

import "time"

type Merchants struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	MerchantCode  string `json:"merchant_code"`
	Balance float64 `json:"balance"`
	UpdatedAt       time.Time   `json:"updated_at"`
}