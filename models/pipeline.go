package models

import "time"

type Pipeline struct {
	ID          string
	Name        string
	Description string

	RepositoryID string

	Default  bool
	Disabled bool

	CreatedAt time.Time
	UpdatedAt time.Time

	URL string
}
