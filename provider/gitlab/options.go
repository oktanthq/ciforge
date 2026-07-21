package gitlab

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// SortDirection specifies the ordering of list results supported by the GitLab API.
type SortDirection string

const (
	SortAscending  SortDirection = "asc"
	SortDescending SortDirection = "desc"
)

// JobScope represents a valid GitLab job status value for filtering.
type JobScope string

const (
	JobScopeCanceled           JobScope = "canceled"
	JobScopeCanceling          JobScope = "canceling"
	JobScopeCreated            JobScope = "created"
	JobScopeFailed             JobScope = "failed"
	JobScopeManual             JobScope = "manual"
	JobScopePending            JobScope = "pending"
	JobScopePreparing          JobScope = "preparing"
	JobScopeRunning            JobScope = "running"
	JobScopeScheduled          JobScope = "scheduled"
	JobScopeSkipped            JobScope = "skipped"
	JobScopeSuccess            JobScope = "success"
	JobScopeWaitingForCallback JobScope = "waiting_for_callback"
	JobScopeWaitingForResource JobScope = "waiting_for_resource"
)

// appendQueryParams builds a query string from the struct fields of opts
// that are tagged with "query". Only non-zero-value fields are included.
// Use the "query-default" tag to supply a value when the field is zero-valued.
func appendQueryParams(path string, opts any) string {
	if opts == nil {
		return path
	}

	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return path
	}

	q := url.Values{}
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		ft := t.Field(i)

		tag := ft.Tag.Get("query")
		if tag == "" {
			continue
		}

		defaultVal := ft.Tag.Get("query-default")

		switch field.Kind() {
		case reflect.String:
			s := field.String()
			if s == "" && defaultVal != "" {
				s = defaultVal
			}
			if s != "" {
				q.Set(tag, s)
			}
		case reflect.Int, reflect.Int64:
			n := field.Int()
			if n == 0 && defaultVal != "" {
				n, _ = strconv.ParseInt(defaultVal, 10, 64)
			}
			if n > 0 {
				q.Set(tag, strconv.FormatInt(n, 10))
			}
		case reflect.Bool:
			if field.Bool() {
				q.Set(tag, "true")
			}
		}
	}

	if len(q) == 0 {
		return path
	}

	return path + "?" + q.Encode()
}

// appendScopeParams appends scope query parameters to the path.
// GitLab expects repeated scope[] parameters for multi-value filtering,
// e.g. scope[]=pending&scope[]=running.
func appendScopeParams(path string, scopes []JobScope) string {
	if len(scopes) == 0 {
		return path
	}

	var b strings.Builder
	b.WriteString(path)
	b.WriteByte('?')
	for i, s := range scopes {
		if i > 0 {
			b.WriteByte('&')
		}
		b.WriteString("scope[]=")
		b.WriteString(string(s))
	}
	return b.String()
}

// ListJobsForProjectOptions configures filtering when listing jobs for a project.
type ListJobsForProjectOptions struct {
	// Scope filters jobs by one or more status values.
	// When multiple scopes are provided, only jobs matching any of them are returned.
	Scope []JobScope

	// OrderBy orders jobs by the given field. Currently only "id" is supported.
	OrderBy string `query:"order_by"`

	// Sort specifies the sort direction: "asc" or "desc". Default is "desc".
	Sort SortDirection `query:"sort" query-default:"desc"`
}

// ListJobsForPipelineOptions configures filtering when listing jobs for a pipeline.
type ListJobsForPipelineOptions struct {
	// Scope filters jobs by one or more status values.
	// When multiple scopes are provided, only jobs matching any of them are returned.
	Scope []JobScope

	// IncludeRetried includes retried jobs in the response. Default is false.
	IncludeRetried bool `query:"include_retried"`
}

// ListTriggerJobsForPipelineOptions configures filtering when listing trigger
// jobs for a pipeline.
type ListTriggerJobsForPipelineOptions struct {
	// Scope filters jobs by one or more status values.
	// When multiple scopes are provided, only jobs matching any of them are returned.
	Scope []JobScope
}

// CancelJobOptions configures the cancel job request.
type CancelJobOptions struct {
	// Force forces cancellation of a job in canceling state.
	Force bool `query:"force"`
}

// PipelineOrderBy specifies the field to order pipeline results by.
type PipelineOrderBy string

const (
	PipelineOrderByID        PipelineOrderBy = "id"
	PipelineOrderByStatus    PipelineOrderBy = "status"
	PipelineOrderByRef       PipelineOrderBy = "ref"
	PipelineOrderByUpdatedAt PipelineOrderBy = "updated_at"
	PipelineOrderByUserID    PipelineOrderBy = "user_id"
)

// PipelineScope filters pipelines by their lifecycle scope.
type PipelineScope string

const (
	PipelineScopeRunning  PipelineScope = "running"
	PipelineScopePending  PipelineScope = "pending"
	PipelineScopeFinished PipelineScope = "finished"
	PipelineScopeBranches PipelineScope = "branches"
	PipelineScopeTags     PipelineScope = "tags"
)

// PipelineStatus filters pipelines by their current status.
type PipelineStatus string

const (
	PipelineStatusCreated            PipelineStatus = "created"
	PipelineStatusWaitingForResource PipelineStatus = "waiting_for_resource"
	PipelineStatusPreparing          PipelineStatus = "preparing"
	PipelineStatusWaitingForCallback PipelineStatus = "waiting_for_callback"
	PipelineStatusPending            PipelineStatus = "pending"
	PipelineStatusRunning            PipelineStatus = "running"
	PipelineStatusSuccess            PipelineStatus = "success"
	PipelineStatusFailed             PipelineStatus = "failed"
	PipelineStatusCanceling          PipelineStatus = "canceling"
	PipelineStatusCanceled           PipelineStatus = "canceled"
	PipelineStatusSkipped            PipelineStatus = "skipped"
	PipelineStatusManual             PipelineStatus = "manual"
	PipelineStatusScheduled          PipelineStatus = "scheduled"
)

// PipelineSource filters pipelines by their triggering source.
type PipelineSource string

const (
	PipelineSourcePush                    PipelineSource = "push"
	PipelineSourceWeb                     PipelineSource = "web"
	PipelineSourceTrigger                 PipelineSource = "trigger"
	PipelineSourceSchedule                PipelineSource = "schedule"
	PipelineSourceAPI                     PipelineSource = "api"
	PipelineSourceExternal                PipelineSource = "external"
	PipelineSourcePipeline                PipelineSource = "pipeline"
	PipelineSourceChat                    PipelineSource = "chat"
	PipelineSourceWebIDE                  PipelineSource = "webide_merge_request_train"
	PipelineSourceMergeRequestEvent       PipelineSource = "merge_request_event"
	PipelineSourceParentPipeline          PipelineSource = "parent_pipeline"
	PipelineSourceExternalPullRequestEvent PipelineSource = "external_pull_request_event"
	PipelineSourceCodecov                 PipelineSource = "codecov"
)

// ListPipelinesOptions configures filtering and pagination when listing
// project pipelines.
//
// Zero-value fields are omitted from the request, allowing GitLab's default
// behavior to be used.
type ListPipelinesOptions struct {
	// Name filters the returned pipelines by name.
	Name string `query:"name"`

	// OrderBy is the field to order pipelines by.
	//
	// Valid values are: "id", "status", "ref", "updated_at", "user_id".
	//
	// Default: "id".
	OrderBy PipelineOrderBy `query:"order_by"`

	// Ref filters pipelines for the specified branch or tag.
	Ref string `query:"ref"`

	// Scope filters pipelines by their lifecycle scope.
	//
	// Valid values are: "running", "pending", "finished", "branches", "tags".
	//
	// When set to "branches" or "tags", only the latest pipeline for each
	// branch or tag ref is returned.
	Scope PipelineScope `query:"scope"`

	// SHA filters pipelines to the specified commit SHA.
	SHA string `query:"sha"`

	// Sort specifies the sort order for pipelines.
	//
	// Valid values are: "asc", "desc".
	//
	// Default: "desc".
	Sort SortDirection `query:"sort"`

	// Source filters pipelines with the specified source.
	Source PipelineSource `query:"source"`

	// Status filters pipelines by their current status.
	Status PipelineStatus `query:"status"`

	// UpdatedAfter filters pipelines updated after the specified date.
	//
	// Expected in ISO 8601 format (2019-03-15T08:00:00Z).
	UpdatedAfter string `query:"updated_after"`

	// UpdatedBefore filters pipelines updated before the specified date.
	//
	// Expected in ISO 8601 format (2019-03-15T08:00:00Z).
	UpdatedBefore string `query:"updated_before"`

	// CreatedAfter filters pipelines created after the specified date.
	//
	// Expected in ISO 8601 format (2019-03-15T08:00:00Z).
	CreatedAfter string `query:"created_after"`

	// CreatedBefore filters pipelines created before the specified date.
	//
	// Expected in ISO 8601 format (2019-03-15T08:00:00Z).
	CreatedBefore string `query:"created_before"`

	// Username filters pipelines triggered by the specified username.
	Username string `query:"username"`

	// YAMLErrors, when true, returns only pipelines with invalid configurations.
	YAMLErrors bool `query:"yaml_errors"`

	// Page is the page number of results to retrieve.
	//
	// Page numbering starts at 1.
	//
	// Default: 1.
	Page int `query:"page"`

	// PerPage is the maximum number of pipelines to return per page.
	//
	// Default: 20.
	PerPage int `query:"per_page"`
}

// GetLatestPipelineOptions configures the query for retrieving the latest
// pipeline for a specific ref.
type GetLatestPipelineOptions struct {
	// Ref filters pipelines for the specified branch or tag.
	//
	// Defaults to the project's default branch when not specified.
	Ref string `query:"ref"`
}

// PipelineScheduleScope filters pipeline schedules by their active state.
type PipelineScheduleScope string

const (
	PipelineScheduleScopeActive   PipelineScheduleScope = "active"
	PipelineScheduleScopeInactive PipelineScheduleScope = "inactive"
)

// ListPipelineSchedulesOptions configures filtering when listing pipeline schedules.
type ListPipelineSchedulesOptions struct {
	// Scope filters pipeline schedules by their active state.
	//
	// Valid values are: "active", "inactive".
	Scope PipelineScheduleScope `query:"scope"`
}

// ListPipelinesByScheduleOptions configures filtering when listing pipelines
// triggered by a pipeline schedule.
type ListPipelinesByScheduleOptions struct {
	// Scope filters pipelines by their lifecycle scope.
	//
	// Valid values are: "running", "pending", "finished", "branches", "tags".
	Scope PipelineScope `query:"scope"`

	// Sort specifies the sort order for pipelines.
	//
	// Valid values are: "asc", "desc".
	//
	// Default: "asc".
	Sort SortDirection `query:"sort"`

	// Status filters pipelines by their current status.
	Status PipelineStatus `query:"status"`

	// UpdatedAfter filters pipelines updated after the specified date.
	//
	// Expected in ISO 8601 format (2019-03-15T08:00:00Z).
	UpdatedAfter string `query:"updated_after"`

	// UpdatedBefore filters pipelines updated before the specified date.
	//
	// Expected in ISO 8601 format (2019-03-15T08:00:00Z).
	UpdatedBefore string `query:"updated_before"`

	// CreatedAfter filters pipelines created after the specified date.
	//
	// Expected in ISO 8601 format (2019-03-15T08:00:00Z).
	CreatedAfter string `query:"created_after"`

	// CreatedBefore filters pipelines created before the specified date.
	//
	// Expected in ISO 8601 format (2019-03-15T08:00:00Z).
	CreatedBefore string `query:"created_before"`
}

// DownloadJobArtifactsByRefOptions configures the download of job artifacts by reference name.
type DownloadJobArtifactsByRefOptions struct {
	// JobToken is the CI/CD job token for multi-project pipelines.
	// Premium and Ultimate only.
	JobToken string `query:"job_token"`

	// SearchRecentSuccessfulPipelines searches across recent successful
	// pipelines instead of just the latest one.
	SearchRecentSuccessfulPipelines bool `query:"search_recent_successful_pipelines"`
}

// DownloadSingleArtifactFileByRefOptions configures the download of a single
// artifact file by reference name.
type DownloadSingleArtifactFileByRefOptions struct {
	// JobToken is the CI/CD job token for multi-project pipelines.
	// Premium and Ultimate only.
	JobToken string `query:"job_token"`

	// SearchRecentSuccessfulPipelines searches across recent successful
	// pipelines instead of just the latest one.
	SearchRecentSuccessfulPipelines bool `query:"search_recent_successful_pipelines"`
}

// ListArtifactTreeOptions configures the listing of files in an artifacts archive.
type ListArtifactTreeOptions struct {
	// Path is the path to browse in the artifacts archive.
	// Defaults to root directory.
	Path string `query:"path"`

	// Recursive returns all entries recursively when true.
	Recursive bool `query:"recursive"`

	// JobToken is the CI/CD job token for multi-project pipelines.
	// Premium and Ultimate only.
	JobToken string `query:"job_token"`
}

// ListProjectVariablesOptions configures filtering when listing project variables.
type ListProjectVariablesOptions struct {
	// Filter filters results when multiple variables share the same key.
	// Possible values: [environment_scope].
	Filter map[string]string `query:"filter"`

	// Page is the page number of results to retrieve.
	//
	// Page numbering starts at 1.
	//
	// Default: 1.
	Page int `query:"page"`

	// PerPage is the maximum number of variables to return per page.
	//
	// Default: 20.
	PerPage int `query:"per_page"`
}

// GetProjectVariableOptions configures filtering when retrieving a single project variable.
type GetProjectVariableOptions struct {
	// Filter filters results when multiple variables share the same key.
	// Possible values: [environment_scope].
	Filter map[string]string `query:"filter"`
}

// DeleteProjectVariableOptions configures filtering when deleting a project variable.
type DeleteProjectVariableOptions struct {
	// Filter filters results when multiple variables share the same key.
	// Possible values: [environment_scope].
	Filter map[string]string `query:"filter"`
}
