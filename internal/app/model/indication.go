package model

import (
	"encoding/json"
	"time"
)

// Indication of device
type Indication struct {
	ID        uint
	DeviceID  uint
	RSSI      int
	CreatedAt time.Time
	Data      IndicationData
}

// IndicationData ...
type IndicationData struct {
	Temperature float32 `json:"temperature,omitempty"`
	Humidity    uint    `json:"humidity,omitempty"`
}

// DataToJSON returns json string
func (i *Indication) DataToJSON() (string, error) {
	b, err := json.Marshal(i.Data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
