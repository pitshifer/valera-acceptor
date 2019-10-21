package sqlstore

import (
	"database/sql"
	"log"
	"time"

	"github.com/pitshifer/valera-acceptor/internal/app/model"
	"github.com/pitshifer/valera-acceptor/internal/app/store"
	"github.com/sirupsen/logrus"
)

// DeviceRepository ...
type DeviceRepository struct {
	store *Store
	cache map[string]*model.Device
}

func (r DeviceRepository) setUp() error {
	rows, err := r.store.db.Query("SELECT id, mac_address, reg_at FROM devices")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		model := &model.Device{}
		if err := rows.Scan(&model.ID, &model.MacAddress, &model.RegAt); err != nil {
			log.Fatal(err)
		}
		r.cache[model.MacAddress] = model
		logrus.Infof("model id = %d and mac address = %s in cache", model.ID, model.MacAddress)
	}

	return nil
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

// FindByMacAddress ...
func (r *DeviceRepository) FindByMacAddress(macAddress string) (*model.Device, error) {
	if device, ok := r.cache[macAddress]; ok {
		return device, nil
	}

	device := &model.Device{}
	if err := r.store.db.QueryRow("SELECT id, mac_address, reg_at FROM devices WHERE mac_address = $1", macAddress).Scan(&device.ID, &device.MacAddress, &device.RegAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	r.cache[device.MacAddress] = device

	return device, nil
}
