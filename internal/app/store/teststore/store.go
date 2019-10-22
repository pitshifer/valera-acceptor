package teststore

import (
	"github.com/pitshifer/valera-acceptor/internal/app/model"
	"github.com/pitshifer/valera-acceptor/internal/app/store"
)

// Store ...
type Store struct {
	deviceRepository     *DeviceRepository
	indicationRepository *IndicationRepository
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
		devices: make(map[uint]*model.Device),
	}
	return s.deviceRepository
}

func (s *Store) Indication() store.IndicationRepository {
	if s.indicationRepository != nil {
		return s.indicationRepository
	}
	s.indicationRepository = &IndicationRepository{
		store:       s,
		indications: make(map[uint]*model.Indication),
	}
	return s.indicationRepository
}

func (s *Store) Run() {
	for {
	}
}

func (s *Store) HandleNewIndication(i *model.Indication) {}
