package store

import "github.com/pitshifer/valera-acceptor/internal/app/model"

// DeviceRepository ...
type DeviceRepository interface {
	Create(*model.Device) error
	FindByID(uint) (*model.Device, error)
}

// IndicationRepository ...
type IndicationRepository interface {
	Insert(*model.Indication) error
}
