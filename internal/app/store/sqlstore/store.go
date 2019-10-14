package sqlstore

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pitshifer/valera-acceptor/internal/app/store"
)

// Store ...
type Store struct {
	db               *sql.DB
	deviceRepository *DeviceRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Device ...
func (s *Store) Device() store.DeviceRepository {
	if s.deviceRepository != nil {
		return s.deviceRepository
	}
	s.deviceRepository = &DeviceRepository{
		store: s,
	}
	return s.deviceRepository
}