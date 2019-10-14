package sqlstore

import (
	"database/sql"
	"time"

	"github.com/pitshifer/valera-acceptor/internal/app/model"
	"github.com/pitshifer/valera-acceptor/internal/app/store"
)

// DeviceRepository ...
type DeviceRepository struct {
	store *Store
}

// Create ...
func (r *DeviceRepository) Create(device *model.Device) error {
	return r.store.db.QueryRow(
		"INSERT INTO devices (mac_address, reg_at) VALUES($1, $2) RETURNING id, reg_at",
		device.MacAddress,
		time.Now(),
	).Scan(&device.ID, &device.RegAt)
}

// FindByID ...
func (r *DeviceRepository) FindByID(ID uint) (*model.Device, error) {
	device := &model.Device{}
	if err := r.store.db.QueryRow("SELECT id, mac_address, reg_at FROM devices WHERE id = $1", ID).Scan(&device.ID, &device.MacAddress, &device.RegAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return device, nil
}
