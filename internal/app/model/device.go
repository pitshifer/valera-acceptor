package model

import "time"

// Device ...
type Device struct {
	ID         uint      `json:"id"`
	MacAddress string    `json:"mac_address"`
	RegAt      time.Time `json:"reg_at"`
}
