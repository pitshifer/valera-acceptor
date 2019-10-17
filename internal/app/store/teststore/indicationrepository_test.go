package teststore_test

import (
	"testing"

	"github.com/pitshifer/valera-acceptor/internal/app/model"
	"github.com/pitshifer/valera-acceptor/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	s := teststore.New()
	err := s.Indication().Insert(&model.Indication{
		ID: 5,
	})

	assert.NoError(t, err)
}
