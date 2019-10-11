package teststore_test

import (
	"testing"

	"github.com/pitshifer/valera-acceptor/internal/app/model"
	"github.com/pitshifer/valera-acceptor/internal/app/store"
	"github.com/pitshifer/valera-acceptor/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestDeviceRepository_Create(t *testing.T) {
	s := teststore.New()
	err := s.Device().Create(&model.Device{
		ID: 1,
	})

	assert.NoError(t, err)
}

func TestDeviceRepository_FindByID(t *testing.T) {
	s := teststore.New()
	ID := uint(1)
	_, err := s.Device().FindByID(ID)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.Device().Create(&model.Device{
		ID: ID,
	})
	d, err := s.Device().FindByID(ID)
	assert.NoError(t, err)
	assert.NotNil(t, d)
}
