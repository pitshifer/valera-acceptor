package teststore

import (
	"github.com/pitshifer/valera-acceptor/internal/app/model"
	"github.com/pitshifer/valera-acceptor/internal/app/store"
)

// Store ...
type Store struct {
	deviceRepository *DeviceRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// Device ...
func (s *Store) Device() store.DeviceRepository {
	if s.deviceRepository != nil {
		return s.deviceRepository
	}
	s.deviceRepository = &DeviceRepository{
		store:   s,
		devices: make(map[string]*model.Device),
	}
	return s.deviceRepository
}
