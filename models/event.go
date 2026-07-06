package models

import "time"

type Event struct {
	ID string

	Type TriggerType

	Actor string

	Timestamp time.Time

	Metadata map[string]any
}
