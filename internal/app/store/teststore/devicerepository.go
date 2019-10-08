package teststore

import (
	"github.com/pitshifer/valera-acceptor/internal/app/model"
	"github.com/pitshifer/valera-acceptor/internal/app/store"
)

// DeviceRepository ...
type DeviceRepository struct {
	store   *Store
	devices map[string]*model.Device
}

// Create ...
func (repo *DeviceRepository) Create(device *model.Device) error {
	repo.devices[device.UUID] = device
	return nil
}

// FindByUUID ...
func (repo *DeviceRepository) FindByUUID(uuid string) (*model.Device, error) {
	if d, ok := repo.devices[uuid]; ok {
		return d, nil
	}

	return nil, store.ErrRecordNotFound
}
