package circleci

import "github.com/abdulahshoaib/ciforge/provider/internal/httpclient"

// SortDirection specifies the ordering of list results supported by the CircleCI API.
type SortDirection = httpclient.SortDirection

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
	// OrgSlug is the organization slug (e.g., "gh/CircleCI-Public").
	// Required for the organization-level pipeline listing endpoint.
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

// ListProjectPipelinesOptions configures filtering and pagination when listing
// pipelines for a project.
type ListProjectPipelinesOptions struct {
	// PageToken is a token for pagination.
	PageToken string `query:"page-token"`

	// Branch filters pipelines by branch name.
	Branch string `query:"branch"`
}

// TriggerPipelineBody represents the request body for triggering a pipeline.
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
