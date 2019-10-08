package model

import "time"

// Device ...
type Device struct {
	UUID         string
	LastActivity *time.Time
}
