package models

import "time"

type Repository struct {
	ID string

	Owner string
	Name  string

	DefaultBranch string

	Private bool
	Fork    bool

	CreatedAt time.Time
	UpdatedAt time.Time

	URL string
}
