package models

import "time"

type Run struct {
	ID         string
	PipelineID string

	Number int64

	Status  RunStatus
	Result  RunResult
	Trigger TriggerType

	Branch    string
	Tag       string
	CommitSHA string
	Message   string

	StartedAt  *time.Time
	FinishedAt *time.Time
	Duration   time.Duration

	URL string
}
