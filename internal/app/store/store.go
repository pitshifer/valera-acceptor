package store

import "github.com/pitshifer/valera-acceptor/internal/app/model"

// Store ...
type Store interface {
	Device() DeviceRepository
	Indication() IndicationRepository

	Run()
	HandleNewIndication(*model.Indication)
}
