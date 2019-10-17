package sqlstore

import (
	"time"

	"github.com/pitshifer/valera-acceptor/internal/app/model"
)

// IndicationRepository ...
type IndicationRepository struct {
	store Store
}

// Insert ...
func (r *IndicationRepository) Insert(model *model.Indication) error {
	data, err := model.DataToJSON()
	if err != nil {
		return err
	}
	return r.store.db.QueryRow("INSERT INTO indicators (device_id, created_at, data) VALUES($1, $2, $3) RETURNING id, created_at", model.DeviceID, time.Now(), data).Scan(&model.ID, &model.CreatedAt)
}
