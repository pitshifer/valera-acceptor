package sqlstore_test

import (
	"testing"
	"time"

	"github.com/pitshifer/valera-acceptor/internal/app/model"
	"github.com/pitshifer/valera-acceptor/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("indications")

	s := sqlstore.New(db)

	err := s.Indication().Insert(&model.Indication{
		ID:        5,
		DeviceID:  1,
		CreatedAt: time.Now(),
		Data: model.IndicationData{
			Type:  "temperature",
			Value: 25.3,
		},
	})

	assert.NoError(t, err)
}
