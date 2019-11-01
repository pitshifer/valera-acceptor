package sqlstore_test

// import (
// 	"testing"
// 	"time"

// 	"github.com/pitshifer/valera-acceptor/internal/app/model"
// 	"github.com/pitshifer/valera-acceptor/internal/app/store"
// 	"github.com/pitshifer/valera-acceptor/internal/app/store/sqlstore"
// 	"github.com/stretchr/testify/assert"
// )

// func TestDeviceRepository_Create(t *testing.T) {
// 	db, teardown := sqlstore.TestDB(t, databaseURL)
// 	defer teardown("devices")

// 	s := sqlstore.New(db)
// 	err := s.Device().Create(&model.Device{
// 		MacAddress: "00:AB:CD:EF:01:30",
// 		RegAt:      time.Now(),
// 	})

// 	assert.NoError(t, err)
// }

// func TestDeviceRepository_FindByID(t *testing.T) {
// 	db, teardown := sqlstore.TestDB(t, databaseURL)
// 	defer teardown("devices")

// 	s := sqlstore.New(db)
// 	ID := uint(1)
// 	_, err := s.Device().FindByID(ID)
// 	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

// 	device := &model.Device{
// 		MacAddress: "00:AB:CD:EF:01:30",
// 		RegAt:      time.Now(),
// 	}
// 	s.Device().Create(device)
// 	d, err := s.Device().FindByID(device.ID)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, d)
// }

// func TestDeviceRepository_FindByMacAddress(t *testing.T) {
// 	db, teardown := sqlstore.TestDB(t, databaseURL)
// 	defer teardown("devices")

// 	s := sqlstore.New(db)
// 	macAddress := "00:AB:CD:EF:01:30"
// 	_, err := s.Device().FindByMacAddress(macAddress)
// 	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

// 	device := &model.Device{
// 		MacAddress: "00:AB:CD:EF:01:30",
// 		RegAt:      time.Now(),
// 	}
// 	s.Device().Create(device)
// 	d, err := s.Device().FindByMacAddress(device.MacAddress)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, d)
// }
