package models

import "time"

// Artifact represents a GitHub Actions artifact produced by a workflow run.
//
// Artifacts are files generated during a workflow execution and stored by
// GitHub for later download. Examples include build outputs, test reports,
// logs, and release packages.
type Artifact struct {
	// ID uniquely identifies the artifact.
	ID int64

	// RunID identifies the workflow run that produced the artifact.
	RunID int64

	// Name is the artifact's display name.
	Name string

	// Type identifies the artifact type.
	Type ArtifactType

	// Size is the size of the artifact in bytes.
	Size int64

	// Expired indicates whether the artifact is no longer available
	// for download.
	Expired bool

	// CreatedAt is when the artifact was created.
	CreatedAt time.Time

	// ExpiresAt is when the artifact will expire and be automatically
	// deleted. It is nil if no expiration is set.
	ExpiresAt *time.Time

	// DownloadURL is the URL used to download the artifact archive.
	DownloadURL string
}
