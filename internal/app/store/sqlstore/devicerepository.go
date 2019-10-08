package sqlstore

import (
	"database/sql"

	"github.com/pitshifer/valera-acceptor/internal/app/model"
	"github.com/pitshifer/valera-acceptor/internal/app/store"
)

// DeviceRepository ...
type DeviceRepository struct {
	store *Store
}

// Create ...
func (r *DeviceRepository) Create(device *model.Device) error {
	if _, err := r.store.db.Exec("INSERT INTO devices(uuid) VALUES(?)", device.UUID); err != nil {
		return err
	}

	return nil
}

// FindByUUID ...
func (r *DeviceRepository) FindByUUID(UUID string) (*model.Device, error) {
	device := &model.Device{}
	if err := r.store.db.QueryRow("SELECT uuid, last_activity FROM devices").Scan(&device.UUID, &device.LastActivity); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return device, nil
}
