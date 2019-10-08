package store

import "github.com/pitshifer/valera-acceptor/internal/app/model"

// DeviceRepository ...
type DeviceRepository interface {
	Create(*model.Device) error
	FindByUUID(string) (*model.Device, error)
}
