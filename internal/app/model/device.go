package model

import "time"

// Device ...
type Device struct {
	UUID         string     `json:"uuid"`
	LastActivity *time.Time `json:"last_activity,omitempty"`
}

// Sanitize ...
func (d *Device) Sanitize() {
	d.LastActivity = nil
}
