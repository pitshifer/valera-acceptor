package teststore

import (
	"github.com/pitshifer/valera-acceptor/internal/app/model"
)

// IndicationRepository ...
type IndicationRepository struct {
	store       *Store
	indications map[uint]*model.Indication
}

// Insert ...
func (r *IndicationRepository) Insert(model *model.Indication) error {
	r.indications[model.ID] = model
	return nil
}
