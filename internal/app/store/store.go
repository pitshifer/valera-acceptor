package store

// Store ...
type Store interface {
	Device() DeviceRepository
}
