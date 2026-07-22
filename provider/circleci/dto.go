package circleci

import "time"

// Pipeline represents a CircleCI pipeline.
//
// A pipeline is a unit of work requested of CircleCI. Each pipeline is triggered
// on a given project, on a given branch, by a particular actor with a
// particular configuration.
type Pipeline struct {
	// ID is the unique identifier of the pipeline.
	ID string `json:"id"`

	// State is the current state of the pipeline.
	State string `json:"state"`

	// Number is the pipeline number (sequential within a project).
	Number int64 `json:"number"`

	// CreatedAt is the timestamp when the pipeline was created.
	CreatedAt time.Time `json:"created_at"`

	// Trigger contains information about what triggered the pipeline.
	Trigger PipelineTrigger `json:"trigger"`

	// VCS contains information about the version control system.
	VCS PipelineVCS `json:"vcs"`

	// Errors contains any errors that occurred during pipeline processing.
	Errors []PipelineError `json:"errors"`
}

// PipelineTrigger contains information about what triggered a pipeline.
type PipelineTrigger struct {
	// Type is the type of trigger (e.g., "webhook", "api", "schedule").
	Type string `json:"type"`

	// ReceivedAt is the timestamp when the trigger was received.
	ReceivedAt time.Time `json:"received_at"`

	// Actor is the user who triggered the pipeline.
	Actor PipelineTriggerActor `json:"actor"`
}

// PipelineTriggerActor represents the user who triggered a pipeline.
type PipelineTriggerActor struct {
	// Login is the user's login on the VCS.
	Login string `json:"login"`

	// AvatarURL is the URL to the user's avatar on the VCS.
	AvatarURL string `json:"avatar_url"`
}

// PipelineVCS contains version control system information for a pipeline.
type PipelineVCS struct {
	// ProviderName is the name of the VCS provider (e.g., "github", "bitbucket").
	ProviderName string `json:"provider_name"`

	// OriginRepositoryURL is the URL of the repository where the trigger originated.
	OriginRepositoryURL string `json:"origin_repository_url"`

	// TargetRepositoryURL is the URL of the target repository.
	TargetRepositoryURL string `json:"target_repository_url"`

	// Revision is the code revision the pipeline ran.
	Revision string `json:"revision"`

	// Branch is the branch where the pipeline ran.
	Branch string `json:"branch"`

	// Tag is the tag that triggered the pipeline, if applicable.
	Tag string `json:"tag,omitempty"`
}

// PipelineError represents an error that occurred during pipeline processing.
type PipelineError struct {
	// Type is the type of error.
	Type string `json:"type"`

	// Message is a human-readable error message.
	Message string `json:"message"`
}

// Workflow represents a CircleCI workflow.
//
// A workflow is a collection of jobs that are run in a defined order with
// parallelism defined by the workflow configuration.
type Workflow struct {
	// ID is the unique identifier of the workflow.
	ID string `json:"id"`

	// Name is the name of the workflow.
	Name string `json:"name"`

	// Status is the current status of the workflow.
	Status string `json:"status"`

	// CreatedAt is the timestamp when the workflow was created.
	CreatedAt time.Time `json:"created_at"`

	// StartedAt is the timestamp when the workflow started running.
	StartedAt *time.Time `json:"started_at,omitempty"`

	// StoppedAt is the timestamp when the workflow stopped running.
	StoppedAt *time.Time `json:"stopped_at,omitempty"`

	// PipelineID is the ID of the pipeline that contains this workflow.
	PipelineID string `json:"pipeline_id"`
}

// Job represents a CircleCI job.
//
// A job is a collection of steps in a workflow. Jobs are run on executors
// and can have dependencies on other jobs.
type Job struct {
	// ID is the unique identifier of the job.
	ID string `json:"id"`

	// Name is the name of the job.
	Name string `json:"name"`

	// Status is the current status of the job.
	Status string `json:"status"`

	// StartedAt is the timestamp when the job started running.
	StartedAt *time.Time `json:"started_at,omitempty"`

	// StoppedAt is the timestamp when the job stopped running.
	StoppedAt *time.Time `json:"stopped_at,omitempty"`

	// JobNumber is the job number (sequential within a project).
	JobNumber int64 `json:"job_number,omitempty"`

	// ProjectSlug is the project slug (e.g., "gh/CircleCI-Public/circleci-cli").
	ProjectSlug string `json:"project_slug,omitempty"`

	// DurationInMillis is the duration of the job in milliseconds.
	DurationInMillis int64 `json:"duration_in_millis,omitempty"`

	// ContainerCount is the number of containers used by the job.
	ContainerCount int `json:"container_count,omitempty"`

	// Executor is the executor used by the job.
	Executor *JobExecutor `json:"executor,omitempty"`

	// Source is the source of the job.
	Source string `json:"source,omitempty"`

	// Messages contains any messages associated with the job.
	Messages []JobMessage `json:"messages,omitempty"`
}

// JobExecutor represents the executor used by a job.
type JobExecutor struct {
	// Type is the type of executor (e.g., "machine", "docker", "resource_class").
	Type string `json:"type"`

	// ResourceClass is the resource class used by the executor.
	ResourceClass string `json:"resource_class,omitempty"`

	// Image is the Docker image used by the executor.
	Image string `json:"image,omitempty"`
}

// JobMessage represents a message associated with a job.
type JobMessage struct {
	// Type is the type of message.
	Type string `json:"type"`

	// Message is the message content.
	Message string `json:"message"`
}

// Artifact represents a job artifact.
type Artifact struct {
	// Path is the path to the artifact.
	Path string `json:"path"`

	// URL is the URL to download the artifact.
	URL string `json:"url"`
}

// TestResult represents a test result from a job.
type TestResult struct {
	// Name is the name of the test.
	Name string `json:"name"`

	// Classname is the test class name.
	Classname string `json:"classname"`

	// File is the file containing the test.
	File string `json:"file"`

	// Result is the test result ("success", "failure", "error", "skipped").
	Result string `json:"result"`

	// Message is the test message (e.g., error message).
	Message string `json:"message"`

	// Duration is the test duration in seconds.
	Duration float64 `json:"duration"`

	// RunTime is the test run time in seconds.
	RunTime float64 `json:"run_time"`
}

// Project represents a CircleCI project.
type Project struct {
	// Slug is the project slug (e.g., "gh/CircleCI-Public/circleci-cli").
	Slug string `json:"slug"`

	// Name is the project name.
	Name string `json:"name"`

	// OrganizationName is the organization name.
	OrganizationName string `json:"organization_name"`

	// VCSInfo contains version control system information.
	VCSInfo ProjectVCSInfo `json:"vcs_info"`

	// DefaultBranch is the default branch of the project.
	DefaultBranch string `json:"default_branch"`
}

// ProjectVCSInfo contains version control system information for a project.
type ProjectVCSInfo struct {
	// VCSURL is the URL of the repository.
	VCSURL string `json:"vcs_url"`

	// ProviderName is the name of the VCS provider.
	ProviderName string `json:"provider_name"`

	// DefaultBranch is the default branch.
	DefaultBranch string `json:"default_branch"`
}

// User represents a CircleCI user.
type User struct {
	// ID is the unique identifier of the user.
	ID string `json:"id"`

	// Login is the user's login on the VCS.
	Login string `json:"login"`

	// Name is the user's display name.
	Name string `json:"name"`

	// AvatarURL is the URL to the user's avatar.
	AvatarURL string `json:"avatar_url"`
}

// pipelineListResponse is the response from listing pipelines.
type pipelineListResponse struct {
	Items         []Pipeline `json:"items"`
	NextPageToken string     `json:"next_page_token"`
}

// workflowListResponse is the response from listing workflows.
type workflowListResponse struct {
	Items         []Workflow `json:"items"`
	NextPageToken string     `json:"next_page_token"`
}

// jobListResponse is the response from listing jobs.
type jobListResponse struct {
	Items         []Job  `json:"items"`
	NextPageToken string `json:"next_page_token"`
}

// artifactListResponse is the response from listing artifacts.
type artifactListResponse struct {
	Items []Artifact `json:"items"`
}

// testListResponse is the response from listing tests.
type testListResponse struct {
	Items []TestResult `json:"items"`
}

type ContextBody struct {
	Name  string        `json:"name"`
	Owner *ContextOwner `json:"owner,omitempty"`
}

type ContextOwner struct {
	ID   string `json:"id"`
	Type string `json:"type"` // "organization" or "account"
}

type Context struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}
