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
		UUID: "c509b714-54f2-4ff3-846f-b152f3f669c9",
	})

	assert.NoError(t, err)
}

func TestDeviceRepository_FindByUUID(t *testing.T) {
	s := teststore.New()
	uuid := "cdcd235f-ef7e-4755-9bb7-3df13ee444cd"
	_, err := s.Device().FindByUUID(uuid)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.Device().Create(&model.Device{
		UUID: uuid,
	})
	d, err := s.Device().FindByUUID(uuid)
	assert.NoError(t, err)
	assert.NotNil(t, d)
}
