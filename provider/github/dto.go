package github

import "time"

type workflowDTO struct {
	ID        int64     `json:"id"`
	NodeID    string    `json:"node_id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	URL       string    `json:"url"`
	HTMLURL   string    `json:"html_url"`
	BadgeURL  string    `json:"badge_url"`
}

type listWorkflowsResponse struct {
	TotalCount int           `json:"total_count"`
	Workflows  []workflowDTO `json:"workflows"`
}

type workflowDispatchRequest struct {
	Ref    string         `json:"ref"`
	Inputs map[string]any `json:"inputs,omitempty"`
}

type workflowDispatchResponse struct {
	WorkflowRunID int64  `json:"workflow_run_id"`
	RunURL        string `json:"run_url"`
	HTMLURL       string `json:"html_url"`
}

type workflowRunDTO struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	RunNumber    int64     `json:"run_number"`
	Status       string    `json:"status"`
	Conclusion   string    `json:"conclusion"`
	HeadBranch   string    `json:"head_branch"`
	HeadSHA      string    `json:"head_sha"`
	DisplayTitle string    `json:"display_title"`
	Event        string    `json:"event"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	HTMLURL      string    `json:"html_url"`
	WorkflowID   int64     `json:"workflow_id"`
}

type listWorkflowRunsResponse struct {
	TotalCount   int              `json:"total_count"`
	WorkflowRuns []workflowRunDTO `json:"workflow_runs"`
}

type jobDTO struct {
	ID          int64      `json:"id"`
	RunID       int64      `json:"run_id"`
	RunURL      string     `json:"run_url"`
	Name        string     `json:"name"`
	Status      string     `json:"status"`
	Conclusion  string     `json:"conclusion"`
	StartedAt   *time.Time `json:"started_at"`
	CompletedAt *time.Time `json:"completed_at"`
	HTMLURL     string     `json:"html_url"`
}

type listJobsResponse struct {
	TotalCount int      `json:"total_count"`
	Jobs       []jobDTO `json:"jobs"`
}

type runDTO struct {
	ID               int64  `json:"id"`
	RepositoryID     int64  `json:"repository_id"`
	HeadRepositoryID int64  `json:"head_repository_id"`
	HeadBranch       string `json:"head_branch"`
	HeadSHA          string `json:"head_sha"`
}

type artifactDTO struct {
	ID                 int64      `json:"id"`
	NodeID             string     `json:"node_id"`
	Name               string     `json:"name"`
	SizeInBytes        int64      `json:"size_in_bytes"`
	URL                string     `json:"url"`
	ArchiveDownloadURL string     `json:"archive_download_url"`
	Expired            bool       `json:"expired"`
	CreatedAt          time.Time  `json:"created_at"`
	ExpiresAt          *time.Time `json:"expires_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
	Digest             string     `json:"digest"`
	WorkflowRun        *runDTO    `json:"workflow_run"`
}

type listArtifactsResponse struct {
	TotalCount int           `json:"total_count"`
	Artifacts  []artifactDTO `json:"artifacts"`
}

type ConcurrencyGroup struct {
	GroupName      string     `json:"group_name"`
	GroupURL       string     `json:"group_url"`
	LastAcquiredAt *time.Time `json:"last_acquired_at"`
}

type listConcurencyGroups struct {
	TotalCount        int                `json:"total_count"`
	ConcurrencyGroups []ConcurrencyGroup `json:"concurrency_groups"`
}
