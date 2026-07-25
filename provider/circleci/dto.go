package circleci

import "time"

// Warnings represents a warning message from the CircleCI API.
type Warnings struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// Pipeline represents a CircleCI pipeline.
//
// A pipeline is a unit of work requested of CircleCI. Each pipeline is triggered
// on a given project, on a given branch, by a particular actor with a
// particular configuration.
type Pipeline struct {
	// ID is the unique identifier of the pipeline.
	ID string `json:"id"`

	// Errors contains any errors that occurred during pipeline processing.
	Errors []PipelineError `json:"errors"`

	// The project-slug for the pipeline.
	ProjectSlug string `json:"project_slug"`

	// The date and time the pipeline was last updated.
	UpdatedAt time.Time `json:"updated_at"`

	// Number is the pipeline number (sequential within a project).
	Number int64 `json:"number"`

	// string or integer or boolean or object
	TriggerParameters map[string]any `json:"trigger_parameters"`

	// State is the current state of the pipeline.
	State string `json:"state"`

	// CreatedAt is the timestamp when the pipeline was created.
	CreatedAt time.Time `json:"created_at"`

	// A sequence of warnings that have occurred within the pipeline.
	Warnings []Warnings `json:"warnings"`

	// Trigger contains information about what triggered the pipeline.
	Trigger PipelineTrigger `json:"trigger"`

	// VCS contains information about the version control system.
	VCS PipelineVCS `json:"vcs"`
}

// PipelineConfig contains the configuration for a pipeline, including
// source, compiled, and setup configurations.
type PipelineConfig struct {
	// The source configuration for the pipeline, before any config compilation
	// has been performed. If there is no config, then this field will be empty.
	Source string `json:"source"`

	// The compiled configuration for the pipeline, after all orb expansion has been
	// performed. If there were errors processing the pipeline's configuration, then
	// this field may be empty.
	Compiled string `json:"compiled"`

	// The setup configuration for the pipeline used for Setup Workflows. If there
	// were errors processing the pipeline's configuration or if setup workflows are
	// not enabled, then this field should not exist
	SetupConfig string `json:"setup-config"`

	// The compiled setup configuration for the pipeline, after all orb expansion has
	// been performed. If there were errors processing the pipeline's setup workflows,
	// then this field may be empty.
	CompiledSetupConfig string `json:"compiled-setup-config"`
}

// triggerPipelineResponse is the response from triggering a pipeline.
type triggerPipelineResponse struct {
	// The unique ID of the pipeline.
	ID string `json:"id"`

	// The current state of the pipeline.
	State string `json:"state"`

	// The number of the pipeline.
	Number int64 `json:"number"`

	// The date and time the pipeline was created.
	CreatedAt time.Time `json:"created_at"`
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

// Commit contains commit message information from the VCS.
type Commit struct {
	//The subject of the commit message.
	Subject *string `json:"subject"`

	// The body of the commit message.
	Body *string `json:"body"`
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
	Branch string `json:"branch,omitempty"`

	// Tag is the tag that triggered the pipeline, if applicable.
	Tag string `json:"tag,omitempty"`

	// The code review id.
	ReviewID string `json:"review_id,omitempty"`

	// The latest commit in the pipeline.
	Commit Commit `json:"commit"`

	// The code review URL
	ReviewURL string `json:"review_url,omitempty"`
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

// PipelineWorkflow represents a workflow within a pipeline, including
// auto-rerun information and detailed status tracking.
type PipelineWorkflow struct {
	// The ID of the pipeline this workflow belongs to.
	PipelineID string `json:"pipeline_id"`

	// UUID of user who canceled the workflow.
	CanceledBy string `json:"canceled_by"`

	// The unique ID of the workflow.
	ID string `json:"id"`

	// Present if this workflow was auto-rerun from a previous workflow.
	// The Nth auto-rerun workflow will have auto_rerun_number N
	AutoRerunNumber int64 `json:"auto_rerun_number"`

	// The name of the workflow.
	Name string `json:"name"`

	// The project-slug for the pipeline this workflow belongs to.
	ProjectSlug string `json:"project_slug"`

	// ErroredBy is the UUID of the user who errored the workflow.
	ErroredBy string `json:"errored_by"`

	// Tag is the tag that triggered the workflow, if applicable.
	Tag *string `json:"tag"`

	// The current status of the workflow.
	Status string `json:"status"`

	// StartedBy is the UUID of the user who started the workflow.
	StartedBy string `json:"started_by"`

	// The maximum number of auto reruns specified for the workflow.
	MaxAutoReruns int64 `json:"max_auto_reruns"`

	// The number of the pipeline this workflow belongs to.
	PipelineNumber int64 `json:"pipeline_number"`

	// The date and time the workflow was created.
	CreatedAt time.Time `json:"created_at"`

	// The date and time the workflow stopped.
	StoppedAt *time.Time `json:"stopped_at"`
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

// Set of organizations of which a user is a member or a collaborator.
type Collaboration struct {
	// The UUID of the organization
	ID string `json:"id"`

	// The VCS provider
	VCSType string `json:"vcs-type"`

	// The name of the organization
	Name string `json:"name"`

	// URL to the user's avatar on the VCS
	AvatarURL string `json:"avatar_url"`

	// The slug of the organization
	Slug string `json:"slug"`
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

// listPipelineWorkflowResponse is the response from listing pipeline workflows.
type listPipelineWorkflowResponse struct {
	Items         []PipelineWorkflow `json:"items"`
	NextPageToken string             `json:"next_page_token"`
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

// ContextBody represents the request body for creating or updating a context.
type ContextBody struct {
	// The user defined name of the context.
	Name  string        `json:"name"`
	Owner *ContextOwner `json:"owner,omitempty"`
}

// ContextOwner represents the owner of a context (organization or account).
type ContextOwner struct {
	// The unique ID of the owner of the context. This is the organization ID.
	// Specify either owner/organization ID or the owner/organization slug.
	// Find the organization ID and slug in the CircleCI web app (Organization Settings > Overview).
	// Owner/organization slug is not supported for CircleCI server.
	ID string `json:"id"`

	// The type of the owner. Defaults to "organization". Use "account" if you are using CircleCI server.
	Type string `json:"type"`
}

// Context represents a CircleCI context, which stores environment variables
// and configuration shared across projects.
type Context struct {
	// The unique ID of the context.
	ID string `json:"id"`

	// The user-defined name of the context.
	Name string `json:"name"`

	// The date and time the context was created.
	CreatedAt time.Time `json:"created_at"`
}

// listEnvironmentResponse is the response from listing environments.
type listEnvironmentResponse struct {
	Items []Environment `json:"items"`

	// The pagination token to use when fetching the next page in this result set.
	NextPageToken *string `json:"next_page_token"`
}

// Label represents a key-value label attached to a resource.
type Label struct {
	// The key of the label
	Key string `json:"key"`

	// The value of the label
	Value string `json:"value"`
}

// Environment represents a CircleCI deployment environment.
//
// Environments are used to manage deployment targets and their configurations.
type Environment struct {
	// The time when the environment was created.
	CreatedAt time.Time `json:"created_at"`

	// Short description given to the environment
	Description string `json:"description"`

	// The environment ID
	ID string `json:"id"`

	// The labels associated to the environment
	Labels Label `json:"labels"`

	// The environment name
	Name string `json:"name"`

	// The time when the environment was updated.
	UpdatedAt time.Time `json:"updated_at"`
}

// Component represents a CircleCI deployment component.
//
// Components are deployable units within a project that can be released independently.
type Component struct {
	// The component ID
	ID string `json:"id"`

	// The component name
	Name string `json:"name"`

	// The project ID of the component
	ProjectID string `json:"project_id"`

	// The labels associated with the component
	Labels []Label `json:"labels"`

	// The number of releases for this Component
	ReleaseCount int `json:"release_count"`

	// The time when the component was created
	CreatedAt time.Time `json:"created_at"`

	// The time when the component was updated
	UpdatedAt time.Time `json:"updated_at"`
}

// listComponentsResponse is the response from listing components.
type listComponentsResponse struct {
	Items []Component `json:"items"`

	// The pagination token to use when fetching the next page in this result set.
	NextPageToken *string `json:"next_page_token"`
}

// ComponentVersion represents a specific version of a deployment component.
type ComponentVersion struct {
	// The version name
	Name string `json:"name"`

	// The ID of the environment in which the version has been deployed
	EnvironmentID string `json:"environment_id"`

	// The namespace in which the version was deployed
	Namespace string `json:"namespace"`

	// Whether the version is live or not
	IsLive bool `json:"is_live"`

	// The ID of the pipeline that deployed the version
	PipelineID string `json:"pipeline_id"`

	// The ID of the CircleCI workflow that deployed the version
	WorkflowID string `json:"workflow_id"`

	// The ID of the CircleCI job that deployed the version
	JobID string `json:"job_id"`

	// The number of the CircleCI job that deployed the version
	JobNumber string `json:"job_number"`

	// The time at which the version was last deployed
	LastDeployedAt time.Time `json:"last_deployed_at"`
}

// listComponentVersionsResponse is the response from listing component versions.
type listComponentVersionsResponse struct {
	Items []ComponentVersion `json:"items"`

	// The pagination token to use when fetching the next page in this result set.
	NextPageToken string `json:"next_page_token"`
}

// listContextResponse is the response from listing contexts.
type listContextResponse struct {
	Items []Context `json:"items"`

	// A token to pass as a page-token query parameter to return the next page of results.
	NextPageToken *string `json:"next_page_token"`
}

// EnvVar represents an environment variable stored in a context.
type EnvVar struct {
	// The name of the environment variable
	Variable string `json:"variable"`

	// The date and time the environment variable was created.
	CreatedAt time.Time `json:"created_at"`

	// The date and time the environment variable was updated
	UpdatedAt time.Time `json:"updated_at"`

	// ID of the context (UUID)
	ContextID string `json:"context_id"`
}

// listEnvVarsResponse is the response from listing environment variables.
type listEnvVarsResponse struct {
	Items []EnvVar `json:"items"`

	// A token to pass as a page-token query parameter to return the next page of results.
	NextPageToken string `json:"next_page_token"`
}

// Restriction represents a context restriction that controls which projects,
// groups, or expressions can access a context.
type Restriction struct {
	// UUID of the context
	ContextID string `json:"context_id"`

	// UUID of the restriction
	ID string `json:"id"`

	// Deprecated: Use RestrictionValue instead.
	// For "project" restrictions, RestrictionValue contains the project ID.
	ProjectID string `json:"project_id"`

	// Contains a human-readable reference for the restriction. For "project" restrictions this is the name of the project.
	Name *string `json:"name"`

	// Enum: "project" "expression" "group"
	// Type of the restriction
	RestrictionType string `json:"restriction_type"`

	// Value used to evaluate the restriction
	RestrictionValue string `json:"restriction_value"`
}

// listContextRestrictionResponse is the response from listing context restrictions.
type listContextRestrictionResponse struct {
	Items []Restriction `json:"items"`

	// Token that can be used to retrieve next page of results
	NextPageToken string `json:"next_page_token"`
}

// RollbackResponse represents the response from a rollback operation.
type RollbackResponse struct {
	// The ID of the rollback pipeline or the command created to handle the rollback
	ID string `json:"id"`

	// The type of the rollback
	RollbackType string `json:"rollback_type"`
}

// Repository represents a version control system repository.
type Repository struct {
	// The fully-qualified name of the repository.
	FullName string `json:"full_name"`

	// External identifier for the repository, as defined by the respective version control provider.
	ExternalID string `json:"external_id"`
}

// AttributionActor represents the user that pipeline runs from a schedule trigger are attributed to.
type AttributionActor *struct {
	// The unique ID of the user that pipeline runs are attributed to.
	ID string `json:"id"`
}

// Schedule represents a cron-based schedule event source for a trigger.
type Schedule struct {
	// The cron expression defining when this trigger fires.
	CronExpression string `json:"cron_expression"`

	// The actor that pipeline runs from this trigger are attributed to.
	// On creation, callers supply current or system as a string; the response returns the resolved actor as an object containing their user ID.
	AttributionActor AttributionActor `json:"attribution_actor"`
}

// EventSource represents the source of events that a trigger listens to.
//
// Contains a repo, webhook, or schedule depending on the provider type.
type EventSource struct {
	// The integration provider for this resource. One of github_app, github_server, github_oauth, webhook, or schedule.
	Provider string `json:"provider"`

	// Information pertaining to the repository used as a source of events for this trigger, if applicable.
	Repo Repository `json:"repo"`

	// Information pertaining to the custom webhook used as a source of events for this trigger, if applicable.
	Webhook Webhook `json:"webhook"`

	// Information pertaining to the schedule used as a source of events for this trigger, if applicable.
	Schedule Schedule `json:"schedule"`
}

// PipelineDefinitionTrigger represents a trigger attached to a pipeline definition.
//
// Triggers define the events and sources that cause pipelines to run.
type PipelineDefinitionTrigger struct {
	// The unique ID of the trigger.
	ID string `json:"id"`

	// The name of the event that will trigger the pipeline.
	Event string `json:"event"`

	// The source of events to use for this trigger.
	// Will contain either a repo, webhook, or schedule object depending on the provider.
	// (The github_app, github_server, and github_oauth providers imply a repo, webhook implies a webhook, and schedule implies a schedule.)
	EventSource EventSource `json:"event_source"`

	// The name of the event preset to use when filtering events for this trigger.
	// When event_source.provider is github_app or github_server, all preset values are supported.
	// When event_source.provider is github_oauth, this field is required and only the following values are accepted: all-pushes, only-build-prs.
	// Not applicable when event_source.provider is webhook or schedule.
	EventPreset string `json:"event_preset"`

	// The ref to use when checking out code for pipeline runs created from this trigger. If empty, the ref provided in the trigger event is used.
	CheckoutRef string `json:"checkout_ref"`

	// The ref to use when fetching config for pipeline runs created from this trigger. If empty, the ref provided in the trigger event is used.
	ConfigRef string `json:"config_ref"`

	// Whether the trigger is disabled. Not supported for triggers where event_source.provider is github_oauth.
	Disabled bool `json:"disabled"`

	// Pipeline parameters passed when running pipelines from this trigger.
	// Currently only supported for triggers where event_source.provider is schedule.
	Parameters map[string]any `json:"parameters"`
}

// UsageExport represents a usage export job for an organization.
type UsageExport struct {
	// string that is a UUID
	UsageExportJobID string `json:"usage_export_job_id"`

	// Enum: "created" "processing" "failed" "completed"
	State string `json:"state"`

	Start time.Time `json:"start"`

	End time.Time `json:"end"`

	// A list of pre signed urls that the client can use to download the results of a Usage Export.
	DownloadURLs []string `json:"download_urls"`
}

// usageExportResponse is the response from getting a usage export job.
type usageExportResponse struct {
	// string that is a UUID
	UsageExportJobID string `json:"usage_export_job_id"`

	// Enum: "created" "processing" "failed" "completed"
	State string `json:"state"`

	// A list of pre signed urls that the client can use to download the results of a Usage Export.
	DownloadURLs []string `json:"download_urls"`

	ErrorReason string `json:"error_reason"`
}

// Scope represents the scope in which relevant events that will trigger webhooks.
type Scope struct {
	// ID of the scope being used (at the moment, only project ID is supported).
	ID string `json:"id"`

	// Type of the scope being used.
	Type string `json:"type"`
}

// Webhook represents a CircleCI webhook.
//
// Webhooks deliver notifications about events (e.g., workflow-completed, job-completed)
// to a specified URL via HTTPS.
type Webhook struct {
	// ID string jsonid
	ID string `json:"id"`

	// Name of the webhook.
	Name string `json:"name"`

	// URL to deliver the webhook to. Note: protocol must be included as well (only https is supported).
	URL string `json:"url"`

	// Whether to enforce TLS certificate verification when delivering the webhook.
	VerifyTLS bool `json:"verify_tls"`

	// Masked value of the secret used to build an HMAC hash of the payload and passed as a header in the webhook request.
	SigningSecret string `json:"signing_secret"`

	// Items Enum: "workflow-completed" "job-completed"
	// Events that will trigger the webhook.
	Events []string `json:"events"`

	// The scope in which the relevant events that will trigger webhooks.
	Scope Scope `json:"scope"`

	// The date and time the webhook was created.
	CreatedAt time.Time `json:"created_at"`

	// The date and time the webhook was last updated.
	UpdatedAt time.Time `json:"updated_at"`
}

// listWebhookResponse is the response from listing webhooks.
type listWebhookResponse struct {
	// Array of objects (Webhook)
	Items []Webhook `json:"items"`

	// A token to pass as a page-token query parameter to return the next page of results.
	NextPageToken *string `json:"next_page_token"`
}
