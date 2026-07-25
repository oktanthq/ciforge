package circleci

import (
	"time"

	"github.com/abdulahshoaib/ciforge/provider/internal/httpclient"
)

// SortDirection specifies the ordering of list results supported by the CircleCI API.
type SortDirection = httpclient.SortDirection

// Supported sort directions for list results.
const (
	SortAscending  SortDirection = httpclient.SortAscending
	SortDescending SortDirection = httpclient.SortDescending
)

// appendQueryParams builds a query string from the struct fields of opts
// that are tagged with "query". Only non-zero-value fields are included.
// Use the "query-default" tag to supply a value when the field is zero-valued.
func appendQueryParams(path string, opts any) string {
	return httpclient.AppendQueryParams(path, opts)
}

// PipelineStatus filters pipelines by their current status.
type PipelineStatus string

// Pipeline status values for filtering pipelines.
const (
	PipelineStatusCreated      PipelineStatus = "created"
	PipelineStatusErrored      PipelineStatus = "error"
	PipelineStatusSetupPending PipelineStatus = "setup-pending"
	PipelineStatusSetup        PipelineStatus = "setup"
	PipelineStatusPending      PipelineStatus = "pending"
	PipelineStatusRunning      PipelineStatus = "running"
	PipelineStatusSuccess      PipelineStatus = "success"
	PipelineStatusFailed       PipelineStatus = "failed"
	PipelineStatusCanceled     PipelineStatus = "canceled"
	PipelineStatusUnauthorized PipelineStatus = "unauthorized"
)

// WorkflowStatus filters workflows by their current status.
type WorkflowStatus string

// Workflow status values for filtering workflows.
const (
	WorkflowStatusQueued       WorkflowStatus = "queued"
	WorkflowStatusSetup        WorkflowStatus = "setup"
	WorkflowStatusPending      WorkflowStatus = "pending"
	WorkflowStatusRunning      WorkflowStatus = "running"
	WorkflowStatusSuccess      WorkflowStatus = "success"
	WorkflowStatusFailed       WorkflowStatus = "failed"
	WorkflowStatusCanceled     WorkflowStatus = "canceled"
	WorkflowStatusUnauthorzied WorkflowStatus = "unauthorized"
)

// ListPipelinesOptions configures filtering and pagination when listing
// pipelines for an organization.
//
// Zero-value fields are omitted from the request, allowing CircleCI's default
// behavior to be used.
type ListPipelinesOptions struct {
	// OrgSlug is the organization slug: "{provider}/{organization}"
	// (e.g., "gh/CircleCI-Public"). Required for the organization-level
	// pipeline listing endpoint.
	OrgSlug string `query:"org-slug"`

	// PageToken is a token for pagination. Use the next_page_token from
	// a previous response to get the next page.
	PageToken string `query:"page-token"`

	// Only include entries created by your user.
	Mine bool `query:"mine"`
}

// GetYourPipelineOptions configures pagination when listing your own pipelines.
type GetYourPipelineOptions struct {
	// PageToken is a token for pagination
	PageToken string `query:"page-token"`
}

// GetPipelineOptions configures filtering when retrieving a pipeline by number.
type GetPipelineOptions struct {
	// Branch filters pipelines by branch name.
	Branch string `query:"branch"`

	// PageToken is a token for pagination
	PageToken string `query:"page-token"`
}

// ListPipelineWorkflowsOptions configures pagination when listing workflows
// for a pipeline.
type ListPipelineWorkflowsOptions struct {
	// PageToken is a token for pagination.
	PageToken string `query:"page-token"`
}

// ListWorkflowJobsOptions configures pagination when listing jobs
// for a workflow.
type ListWorkflowJobsOptions struct {
	// PageToken is a token for pagination.
	PageToken string `query:"page-token"`
}

// ContinuePipelineBody represents the request body for continuing a pipeline.
type ContinuePipelineBody struct {
	// A pipeline continuation key.
	ContinuationKey string `json:"continuation-key"`

	// A configuration string for the pipeline.
	Configuration string `json:"configuration"`

	// An object containing pipeline parameters and their values. Pipeline parameters
	// have the following size limits: 100 max entries, 128 maximum key length, 512 maximum value length.
	Parameters map[string]any `json:"parameters,omitempty"`
}

// Config specifies the branch or tag to use when fetching the pipeline configuration.
type Config struct {
	// The branch that should be used to fetch the config file. Note that branch and tag are mutually exclusive.
	// To trigger a pipeline for a PR by number use pull//head for the PR ref or pull//merge for the merge ref (GitHub only)
	Branch string `json:"branch"`

	// The tag that should be used to fetch the config file. The commit that this tag points to is used for the pipeline.
	// Note that branch and tag are mutually exclusive.
	Tag string `json:"tag"`
}

// Checkout specifies the branch or tag to use when checking out code.
type Checkout struct {
	// The branch that should be used to check out code on a checkout step. Note that branch and tag are mutually exclusive.
	// To trigger a pipeline for a PR by number use pull//head for the PR ref or pull//merge for the merge ref (GitHub only)
	Branch string `json:"branch"`

	// The tag that should be used to check out code on a checkout step. The commit that this tag points to is used for the pipeline.
	// Note that branch and tag are mutually exclusive.
	Tag string `json:"tag"`
}

// TriggerProjPipelineBody represents the request body for triggering a project pipeline.
type TriggerProjPipelineBody struct {
	// The unique id for the pipeline definition. This can be found in the page Project Settings > Pipelines.
	DefinitionID string `json:"definition_id"`

	Config Config `json:"config"`

	Checkout Checkout `json:"checkout"`

	Parameters map[string]any `json:"parameters"`
}

// ListProjectPipelinesOptions configures filtering and pagination when listing
// pipelines for a project.
type ListProjectPipelinesOptions struct {
	// PageToken is a token for pagination.
	PageToken string `query:"page-token"`

	// Branch filters pipelines by branch name.
	Branch string `query:"branch"`
}

// TriggerPipelineBody represents the request body for triggering a pipeline.
//
// The project slug format is "{provider}/{organization}/{project}" where:
//   - provider: "github"/"gh", "bitbucket"/"bb", or "circleci"
//   - organization: The org name or ID (e.g., "CircleCI-Public")
//   - project: The project name or ID (e.g., "api-preview-docs")
type TriggerPipelineBody struct {
	// Branch is the branch to run the pipeline on.
	// Optional if a tag is provided.
	Branch string `json:"branch,omitempty"`

	// Tag is the tag to run the pipeline on.
	// Optional if a branch is provided.
	Tag string `json:"tag,omitempty"`

	// Parameters are the pipeline parameters to use for the pipeline run.
	Parameters map[string]any `json:"parameters,omitempty"`
}

// GetJobArtifactsOptions configures pagination when listing job artifacts.
type GetJobArtifactsOptions struct {
	// PageToken is a token for pagination.
	PageToken string `query:"page-token"`
}

// GetJobTestsOptions configures pagination when listing job test results.
type GetJobTestsOptions struct {
	// PageToken is a token for pagination.
	PageToken string `query:"page-token"`
}

// ListWorkflowMetricsOptions configures filtering when listing workflow metrics.
type ListWorkflowMetricsOptions struct {
	// Branch filters metrics by branch name.
	Branch string `query:"branch"`

	// DateRange filters metrics by date range.
	// Valid values: "7d", "30d", "90d".
	DateRange string `query:"date-range"`
}

// ListWorkflowRunsOptions configures filtering when listing workflow run history.
type ListWorkflowRunsOptions struct {
	// Branch filters runs by branch name.
	Branch string `query:"branch"`

	// PageToken is a token for pagination.
	PageToken string `query:"page-token"`

	// DateRange filters runs by date range.
	// Valid values: "7d", "30d", "90d".
	DateRange string `query:"date-range"`
}

// ListJobMetricsOptions configures filtering when listing job metrics.
type ListJobMetricsOptions struct {
	// Branch filters metrics by branch name.
	Branch string `query:"branch"`

	// DateRange filters metrics by date range.
	// Valid values: "7d", "30d", "90d".
	DateRange string `query:"date-range"`
}

// ListEnvironmentsOptions configures filtering and pagination when listing
// deployment environments.
type ListEnvironmentsOptions struct {
	// An opaque identifier of an organization used in query parameters.
	OrgID string `query:"org-id"`

	// The number of results per page.
	PageSize int `query:"page-size"`

	// A token to specify which page of results to fetch.
	PageToken string `query:"page-token"`
}

// ListComponentsOptions configures filtering and pagination when listing
// deployment components.
type ListComponentsOptions struct {
	// An opaque identifier of an organization used in query parameters.
	OrgID string `query:"org-id"`

	// An opaque identifier of an project used in query parameters.
	ProjectID string `query:"project-id"`

	// The number of results per page.
	PageSize int `query:"page-size"`

	// A token to specify which page of results to fetch.
	PageToken string `query:"page-token"`
}

// ListContextOptions configures filtering and pagination when listing contexts.
type ListContextOptions struct {
	// The unique ID of the owner of the context. This is the organization ID.
	// Specify either owner/organization ID or the owner/organization slug.
	// Find the organization ID and slug in the CircleCI web app (Organization Settings > Overview).
	// Owner/organization slug is not supported for CircleCI server.
	OwnerID string `json:"owner-id"`

	// A string that represents an organization. This is the organization slug.
	// Specify either this or organization/owner ID. Find the organization ID and slug in the CircleCI web app (Organization Settings > Overview).
	// Owner/organization slug is not supported for CircleCI server.
	OwnerSlug string `json:"owner-slug"`

	// The type of the owner. Defaults to "organization". Use "account" if you are using CircleCI server.
	OwnerType string `json:"owner-type"`

	// A token to specify which page of results to fetch.
	PageToken string `json:"page-token"`
}

// ListEnvVarOptions configures pagination when listing environment variables.
type ListEnvVarOptions struct {
	// A token to specify which page of results to fetch.
	PageToken string `json:"page-token"`
}

// UpsertEnvVarBody represents the request body for creating or updating an environment variable.
type UpsertEnvVarBody struct {
	// The value of the environment variable
	Value string `json:"value"`
}

// CreateContextRestrictionsBody represents the request body for creating a context restriction.
type CreateContextRestrictionsBody struct {
	// Enum: "project" "expression" "group"
	// Type of the restriction.
	RestrictionType string `json:"restriction_type"`

	// Value used to evaluate the restriction.
	// If the restriction_type is project, this will be the project UUID.
	// If the restriction_type is expression, this will be the expression rule.
	RestrictionValue string `json:"restriction_value"`
}

// RollbackBody represents the request body for rolling back a component deployment.
type RollbackBody struct {
	// The component name
	ComponentName string `json:"component_name"`

	// The current version
	CurrentVersion string `json:"current_version"`

	// The environment name
	EnvironmentName string `json:"environment_name"`

	// The namespace
	Namespace string `json:"namespace"`

	// The extra parameters for the rollback pipeline
	Parameters map[string]any `json:"parameters"`

	// The reason for the rollback
	Reason string `json:"reason"`

	// The target version
	TargetVersion string `json:"target_version"`
}

// CreateTriggerBody represents the request body for creating or updating a trigger.
type CreateTriggerBody struct {
	// 	The source of events to use for this trigger.
	// The repo object must be specified when provider is github_app, github_server,
	// or github_oauth, the webhook object may only be specified when provider is webhook,
	// and the schedule object must be specified when provider is schedule.
	EventSource EventSource `json:"event_source"`

	// The event preset to use for this trigger.
	EventPreset string `json:"event_preset"`

	// The ref to use when checking out code for pipeline runs created from this trigger. Always required when event_source.provider
	// is webhook or schedule. When event_source.provider is github_app or github_server, only expected if the event source repository
	// (identified by event_source.repo.external_id) is different to the checkout source repository of the associated Pipeline Definition.
	// Otherwise, must be omitted.
	CheckoutRef string `json:"checkout_ref"`

	// The ref to use when fetching config for pipeline runs created from this trigger.
	// Always required when event_source.provider is webhook or schedule.
	// When event_source.provider is github_app or github_server, only expected if the event source
	// repository (identified by event_source.repo.external_id) is different to the config source repository of the associated Pipeline Definition.
	// Otherwise, must be omitted.
	ConfigRef string `json:"config_ref"`

	// The name of the triggering event. Required when provider is webhook or schedule. Should not be set for other providers.
	EventName string `json:"event_name"`

	// Whether the trigger should be disabled upon creation. Not supported for triggers where event_source.provider is github_oauth.
	Disabled bool `json:"disabled"`

	// Pipeline parameters to pass when running pipelines from this trigger.
	// Currently only supported for triggers where event_source.provider is schedule.
	Parameters map[string]any `json:"parameters"`
}

// UsageExportBody represents the request body for creating a usage export job.
type UsageExportBody struct {
	// The start date & time (inclusive) of the range from which data will be pulled. Must be no more than one year ago.
	Start time.Time `json:"start"`

	// The end date & time (inclusive) of the range from which data will be pulled. Must be no more than 31 days after start.
	End time.Time `json:"end"`

	// array of UUIDs
	SharedOrgIDs []string `json:"shared_org_ids"`
}

// ListWebhooksOptions configures filtering when listing webhooks.
type ListWebhooksOptions struct {
	// ID of the scope being used (at the moment, only project ID is supported).
	ScopeID string `query:"scope-id"`

	// Type of the scope being used (at the moment, only project is supported).
	ScopeType string `query:"scope-type"`
}

// CreateWebhooksBody represents the request body for creating a webhook.
type CreateWebhooksBody struct {
	// Name of the webhook.
	Name string `json:"name"`

	// Events that will trigger the webhook.
	Events []string `json:"events"`

	// URL to deliver the webhook to. Note: protocol must be included as well (only https is supported).
	URL string `json:"url"`

	// Whether to enforce TLS certificate verification when delivering the webhook.
	VerifyTLS bool `json:"verify-tls"`

	// Secret used to build an HMAC hash of the payload and passed as a header in the webhook request.
	SigningSecret string `json:"signing-secret"`

	// The scope in which the relevant events that will trigger webhooks.
	Scope Scope `json:"scope"`
}

// UpdateWebhookBody represents the request body for updating a webhook.
type UpdateWebhookBody struct {
	// Name of the webhook.
	Name string `json:"name"`

	// Events that will trigger the webhook.
	Event []string `json:"event"`

	// URL to deliver the webhook to. Note: protocol must be included as well (only https is supported).
	URL string `json:"url"`

	// Secret used to build an HMAC hash of the payload and passed as a header in the webhook request.
	SigningSecret string `json:"signing-secret"`

	// Whether to enforce TLS certificate verification when delivering the webhook.
	VerifyTLS bool `json:"verify-tls"`
}

// CreateProjectBody represents the request body for creating a project.
type CreateProjectBody struct {
	// The name of the project
	Name string `json:"name"`
}

// CheckoutKeyType specifies the type of checkout key.
type CheckoutKeyType string

// Supported checkout key types.
const (
	CheckoutKeyTypeUserKey   CheckoutKeyType = "user-key"
	CheckoutKeyTypeDeployKey CheckoutKeyType = "deploy-key"
)

// CreateCheckoutKeyBody represents the request body for creating a checkout key.
type CreateCheckoutKeyBody struct {
	Type CheckoutKeyType `json:"type"`
}

// CreateProjectEnvVarBody represents the request body for creating a project environment variable.
type CreateProjectEnvVarBody struct {
	// The name of the environment variable.
	Name string `json:"name"`

	// The value of the environment variable.
	Value string `json:"value"`
}

// UpdateProjectSettingsBody represents the request body for updating project settings.
type UpdateProjectSettingsBody struct {
	Advanced ProjectSettings `json:"advanced"`
}

// RerunWorkflowBody represents the request body for rerunning a workflow.
type RerunWorkflowBody struct {
	// Whether to enable SSH access for the triggering user on the newly-rerun job.
	// Requires the jobs parameter to be used and so is mutually exclusive with the from_failed parameter.
	EnableSSH bool `json:"enable_ssh"`

	// Whether to rerun the workflow from the failed job. Mutually exclusive with the jobs parameter.
	FromFailed bool `json:"from_failed"`

	// A list of job IDs to rerun.
	Jobs []string `json:"jobs"`

	// Completes rerun using sparse trees logic, an optimization for workflows that have disconnected subgraphs.
	// Requires jobs parameter and so is mutually exclusive with the from_failed parameter.
	SparseTree bool `json:"sparse_tree"`
}
