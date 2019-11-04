package sqlstore

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pitshifer/valera-acceptor/internal/app/model"
	"github.com/pitshifer/valera-acceptor/internal/app/store"
	"github.com/sirupsen/logrus"
)

// Store ...
type Store struct {
	db                   *sql.DB
	deviceRepository     *DeviceRepository
	indicationRepository *IndicationRepository
	indicationCh         chan *model.Indication
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db:           db,
		indicationCh: make(chan *model.Indication, 1000),
	}
}

// Device ...
func (s *Store) Device() store.DeviceRepository {
	if s.deviceRepository != nil {
		return s.deviceRepository
	}
	s.deviceRepository = &DeviceRepository{
		store: s,
		cache: make(map[string]*model.Device, 30),
	}
	s.deviceRepository.setUp()
	return s.deviceRepository
}

// Indication ...
func (s *Store) Indication() store.IndicationRepository {
	if s.indicationRepository != nil {
		return s.indicationRepository
	}
	s.indicationRepository = &IndicationRepository{
		store: s,
	}
	return s.indicationRepository
}

// Run ...
func (s *Store) Run() {
	s.Device()

	for {
		select {
		case indication := <-s.indicationCh:
			err := s.Indication().Insert(indication)
			if err != nil {
				logrus.Errorf("Fail on insert a new indication - %v: %s", indication, err)
				break
			}
			logrus.Infof("New indication was saved - %v", indication)
		}
	}
}

func (s *Store) HandleNewIndication(i *model.Indication) {
	s.indicationCh <- i
}
