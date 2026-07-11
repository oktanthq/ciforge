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
	ID               int64     `json:"id"`
	Name             string    `json:"name"`
	NodeID           string    `json:"node_id"`
	CheckSuiteID     int64     `json:"check_suite_id"`
	CheckSuiteNodeID string    `json:"check_suite_node_id"`
	HeadBranch       string    `json:"head_branch"`
	HeadSHA          string    `json:"head_sha"`
	Path             string    `json:"path"`
	RunNumber        int64     `json:"run_number"`
	Event            string    `json:"event"`
	DisplayTitle     string    `json:"display_title"`
	Status           string    `json:"status"`
	Conclusion       string    `json:"conclusion"`
	WorkflowID       int64     `json:"workflow_id"`
	URL              string    `json:"url"`
	HTMLURL          string    `json:"html_url"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	RunAttempt       int64     `json:"run_attempt"`
	RunStartedAt     time.Time `json:"run_started_at"`
}

type listWorkflowRunsResponse struct {
	TotalCount   int              `json:"total_count"`
	WorkflowRuns []workflowRunDTO `json:"workflow_runs"`
}

type stepsDTO struct {
	Name        string `json:"name"`
	Status      string `json:"status"`
	Conclusion  string `json:"conclusion"`
	Number      int64  `json:"number"`
	StartedAt   string `json:"started_at"`
	CompletedAt string `json:"completed_at"`
}

type jobDTO struct {
	ID              int64      `json:"id"`
	RunID           int64      `json:"run_id"`
	RunURL          string     `json:"run_url"`
	NodeID          string     `json:"node_id"`
	HeadSHA         string     `json:"head_sha"`
	URL             string     `json:"url"`
	HTMLURL         string     `json:"html_url"`
	Status          string     `json:"status"`
	Conclusion      string     `json:"conclusion"`
	StartedAt       *time.Time `json:"started_at"`
	CompletedAt     *time.Time `json:"completed_at"`
	Name            string     `json:"name"`
	Steps           []stepsDTO `json:"steps"`
	CheckRunURL     string     `json:"check_run_url"`
	Labels          []string   `json:"labels"`
	RunnerID        int64      `json:"runner_id"`
	RunnerName      string     `json:"runner_name"`
	RunnerGroupID   int64      `json:"runner_group_id"`
	RunnerGroupName string     `json:"runner_group_name"`
	WorkflowName    string     `json:"workflow_name"`
	HeadBranch      string     `json:"head_branch"`
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

type workflowApprovalDTO struct {
	State        string           `json:"state"`
	Comment      string           `json:"comment"`
	Environments []environmentDTO `json:"environments"`
	User         userDTO          `json:"user"`
}

type environmentDTO struct {
	ID        int64     `json:"id"`
	NodeID    string    `json:"node_id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	HTMLURL   string    `json:"html_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type userDTO struct {
	Login             string `json:"login"`
	ID                int64  `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}
