package store

import "github.com/pitshifer/valera-acceptor/internal/app/model"

// DeviceRepository ...
type DeviceRepository struct {
	store *Store
}

// Create ...
func (r *DeviceRepository) Create(device *model.Device) (*model.Device, error) {
	if _, err := r.store.db.Exec("INSERT INTO devices(uuid) VALUES(?)", device.UUID); err != nil {
		return nil, err
	}

	return device, nil
}

// FindByUUID ...
func (r *DeviceRepository) FindByUUID(UUID string) (*model.Device, error) {
	return nil, nil
}
