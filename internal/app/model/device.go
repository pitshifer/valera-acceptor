package model

import "time"

// Device ...
type Device struct {
	UUID         string
	lastActivity *time.Time
}
