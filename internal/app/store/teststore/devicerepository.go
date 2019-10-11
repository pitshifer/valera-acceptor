package teststore

import (
	"github.com/pitshifer/valera-acceptor/internal/app/model"
	"github.com/pitshifer/valera-acceptor/internal/app/store"
)

// DeviceRepository ...
type DeviceRepository struct {
	store   *Store
	devices map[uint]*model.Device
}

// Create ...
func (repo *DeviceRepository) Create(device *model.Device) error {
	repo.devices[device.ID] = device
	return nil
}

// FindByID ...
func (repo *DeviceRepository) FindByID(ID uint) (*model.Device, error) {
	if d, ok := repo.devices[ID]; ok {
		return d, nil
	}

	return nil, store.ErrRecordNotFound
}
