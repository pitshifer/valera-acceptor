package store_test

import (
	"testing"

	"github.com/pitshifer/valera-acceptor/internal/app/model"
	"github.com/pitshifer/valera-acceptor/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestDeviceRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("devices")

	d, err := s.Device().Create(&model.Device{
		UUID: "c509b714-54f2-4ff3-846f-b152f3f669c9",
	})

	assert.NoError(t, err)
	assert.NotNil(t, d)
}
