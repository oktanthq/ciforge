package models

import "time"

type Artifact struct {
	ID    string
	RunID string

	Name string

	Type ArtifactType

	Size int64

	Expired bool

	CreatedAt time.Time
	ExpiresAt *time.Time

	DownloadURL string
}
