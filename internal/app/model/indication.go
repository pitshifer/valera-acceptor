package model

import (
	"encoding/json"
	"time"
)

// Indication of device
type Indication struct {
	ID        uint
	DeviceID  uint
	CreatedAt time.Time
	Data      IndicationData
}

// IndicationData ...
type IndicationData struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
	Date  time.Time   `json:"date"`
}

// DataToJSON returns json string
func (i *Indication) DataToJSON() (string, error) {
	b, err := json.Marshal(i.Data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
