package models

import "time"

type Job struct {
	ID    string
	RunID string

	Name  string
	Stage string

	Status RunStatus
	Result RunResult

	Runner string

	StartedAt  *time.Time
	FinishedAt *time.Time
	Duration   time.Duration

	URL string
}
