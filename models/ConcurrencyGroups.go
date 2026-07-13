package models

import "time"

type ConcurencyGroup struct {
	GroupName      string
	GroupURL       string
	LastAcquiredAt *time.Time
}
