package github

import (
	"encoding/json"
	"time"
)

type GitHubApp struct {
	ID                 int64          `json:"id"`
	Slug               string         `json:"slug"`
	NodeID             string         `json:"node_id"`
	ClientID           string         `json:"client_id"`
	Owner              any            `json:"owner"`
	Name               string         `json:"name"`
	Description        *string        `json:"description"`
	ExternalURL        string         `json:"external_url"`
	HTMLURL            string         `json:"html_url"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	Permissions        AppPermissions `json:"permissions"`
	Events             []string       `json:"events"`
	InstallationsCount *int           `json:"installations_count,omitempty"`
}

type AppPermissions map[string]string

type WorkflowRunRef struct {
	ID         int64  `json:"id"`
	RepoID     int64  `json:"repository_id"`
	HeadRepoID int64  `json:"head_repository_id"`
	HeadBranch string `json:"head_branch"`
	HeadSHA    string `json:"head_sha"`
}

type Artifact struct {
	ID                 int64           `json:"id"`
	NodeID             string          `json:"node_id"`
	Name               string          `json:"name"`
	SizeInBytes        int64           `json:"size_in_bytes"`
	URL                string          `json:"url"`
	ArchiveDownloadURL string          `json:"archive_download_url"`
	Expired            bool            `json:"expired"`
	CreatedAt          time.Time       `json:"created_at"`
	ExpiresAt          *time.Time      `json:"expires_at"`
	UpdatedAt          *time.Time      `json:"updated_at"`
	Digest             string          `json:"digest"`
	WorkflowRun        *WorkflowRunRef `json:"workflow_run"`
}

type artifactListResponse struct {
	TotalCount int        `json:"total_count"`
	Artifacts  []Artifact `json:"artifacts"`
}

type Workflow struct {
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

type workflowListResponse struct {
	TotalCount int        `json:"total_count"`
	Workflows  []Workflow `json:"workflows"`
}

type WorkflowDispatchRequest struct {
	Ref    string         `json:"ref"`
	Inputs map[string]any `json:"inputs,omitempty"`
}

type workflowDispatchResponse struct {
	WorkflowRunID int64  `json:"workflow_run_id"`
	RunURL        string `json:"run_url"`
	HTMLURL       string `json:"html_url"`
}

type WorkflowDispatchEvent struct {
	RunID  int64
	RunURL string
	URL    string
}

type WorkflowRun struct {
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

type workflowRunListResponse struct {
	TotalCount   int           `json:"total_count"`
	WorkflowRuns []WorkflowRun `json:"workflow_runs"`
}

type Step struct {
	Name        string `json:"name"`
	Status      string `json:"status"`
	Conclusion  string `json:"conclusion"`
	Number      int64  `json:"number"`
	StartedAt   string `json:"started_at"`
	CompletedAt string `json:"completed_at"`
}

type Job struct {
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
	Steps           []Step     `json:"steps"`
	CheckRunURL     string     `json:"check_run_url"`
	Labels          []string   `json:"labels"`
	RunnerID        int64      `json:"runner_id"`
	RunnerName      string     `json:"runner_name"`
	RunnerGroupID   int64      `json:"runner_group_id"`
	RunnerGroupName string     `json:"runner_group_name"`
	WorkflowName    string     `json:"workflow_name"`
	HeadBranch      string     `json:"head_branch"`
}

type jobListResponse struct {
	TotalCount int   `json:"total_count"`
	Jobs       []Job `json:"jobs"`
}

type WorkflowApproval struct {
	State        string        `json:"state"`
	Comment      string        `json:"comment"`
	Environments []Environment `json:"environments"`
	User         User          `json:"user"`
}

type Environment struct {
	ID        int64     `json:"id"`
	NodeID    string    `json:"node_id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	HTMLURL   string    `json:"html_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
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

type ConcurrencyGroup struct {
	GroupName      string     `json:"group_name"`
	GroupURL       string     `json:"group_url"`
	LastAcquiredAt *time.Time `json:"last_acquired_at"`
}

type concurrencyGroupListResponse struct {
	TotalCount        int64              `json:"total_count"`
	ConcurrencyGroups []ConcurrencyGroup `json:"concurrency_groups"`
}

type BillableRunResponse struct {
	Billable      Billable `json:"billable"`
	RunDurationMS int64    `json:"run_duration_ms"`
}

type BillableWorkflowResponse struct {
	Billable BillableWorkflow `json:"billable"`
}

type BillableWorkflow struct {
	Ubuntu  *WorkflowUsage `json:"UBUNTU,omitempty"`
	MacOS   *WorkflowUsage `json:"MACOS,omitempty"`
	Windows *WorkflowUsage `json:"WINDOWS,omitempty"`
}

type WorkflowUsage struct {
	TotalMS int64 `json:"total_ms"`
}

type Billable struct {
	Ubuntu  *RunnerUsage `json:"UBUNTU,omitempty"`
	MacOS   *RunnerUsage `json:"MACOS,omitempty"`
	Windows *RunnerUsage `json:"WINDOWS,omitempty"`
}

type RunnerUsage struct {
	TotalMS int64      `json:"total_ms"`
	Jobs    int        `json:"jobs"`
	JobRuns []JobUsage `json:"job_runs,omitempty"`
}

type JobUsage struct {
	JobID      int64 `json:"job_id"`
	DurationMS int64 `json:"duration_ms"`
}

type Deployment struct {
	URL                   string          `json:"url"`
	ID                    int64           `json:"id"`
	NodeID                string          `json:"node_id"`
	SHA                   string          `json:"sha"`
	Ref                   string          `json:"ref"`
	Task                  string          `json:"task"`
	Payload               json.RawMessage `json:"payload"`
	OriginalEnvironment   string          `json:"original_environment"`
	Environment           string          `json:"environment"`
	Description           *string         `json:"description"`
	Creator               *User           `json:"creator"`
	CreatedAt             time.Time       `json:"created_at"`
	UpdatedAt             time.Time       `json:"updated_at"`
	StatusesURL           string          `json:"statuses_url"`
	RepositoryURL         string          `json:"repository_url"`
	TransientEnvironment  bool            `json:"transient_environment"`
	ProductionEnvironment bool            `json:"production_environment"`
	PerformedViaGitHubApp *GitHubApp      `json:"performed_via_github_app"`
}

type PendingDeployment struct {
	Environment           PendingDeploymentEnvironment `json:"environment"`
	WaitTimer             int                          `json:"wait_timer"`
	WaitTimerStartedAt    *time.Time                   `json:"wait_timer_started_at"`
	CurrentUserCanApprove bool                         `json:"current_user_can_approve"`
	Reviewers             []DeploymentReviewer         `json:"reviewers"`
}

type PendingDeploymentEnvironment struct {
	ID      int64  `json:"id"`
	NodeID  string `json:"node_id"`
	Name    string `json:"name"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
}

type DeploymentReviewer struct {
	Type     string `json:"type"`
	Reviewer User   `json:"reviewer"`
}
