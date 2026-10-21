package github

import "github.com/oktanthq/ciforge/provider/internal/httpclient"

// SortDirection specifies the ordering of list results supported by the GitHub API.
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

// ListArtifactsOptions configures pagination and filtering when listing
// artifacts for a repository.
//
// Zero-value fields are omitted from the request, allowing GitHub's default
// behavior to be used.
type ListArtifactsOptions struct {
	// Page is the page number of results to retrieve.
	//
	// Page numbering starts at 1.
	//
	// Default: 1.
	Page int `query:"page"`

	// PerPage is the maximum number of artifacts to return per page.
	//
	// Valid values range from 1 to 100.
	//
	// Default: 30.
	PerPage int `query:"per_page"`

	// Name filters the returned artifacts by their name.
	//
	// When specified, only artifacts whose name exactly matches this value
	// are returned.
	Name string `query:"name"`
}

// ListRunsForRepoOptions configures the query parameters for listing
// workflow runs in a repository.
//
// Zero-value fields are omitted from the request, allowing GitHub's default
// behavior to be used.
type ListRunsForRepoOptions struct {
	// Actor filters workflow runs by the login of the user who created
	// the push associated with the check suite or workflow run.
	Actor string `query:"actor"`

	// Branch filters workflow runs associated with a branch.
	Branch string `query:"branch"`

	// Event filters workflow runs triggered by the specified event, such
	// as "push", "pull_request", or "issue".
	Event string `query:"event"`

	// Status filters workflow runs by their check run status or
	// conclusion.
	//
	// Valid values include: "completed", "action_required", "cancelled",
	// "failure", "neutral", "skipped", "stale", "success", "timed_out",
	// "in_progress", "queued", "requested", "waiting", "pending".
	Status string `query:"status"`

	// PerPage is the number of results per page (max 100).
	//
	// Default: 30.
	PerPage int `query:"per_page"`

	// Page is the page number of the results to fetch.
	//
	// Default: 1.
	Page int `query:"page"`

	// Created returns workflow runs created within the given date-time
	// range. See "Understanding the search syntax" for the accepted
	// syntax.
	Created string `query:"created"`

	// ExcludePullRequests, when true, omits pull requests from the
	// response.
	//
	// Default: false.
	ExcludePullRequests bool `query:"exclude_pull_requests"`

	// CheckSuiteID filters workflow runs by the specified check suite
	// identifier.
	CheckSuiteID int `query:"check_suite_id"`

	// HeadSHA filters workflow runs to only those associated with the
	// specified commit SHA.
	HeadSHA string `query:"head_sha"`
}

// ListWorkflowRunArtifactsOptions configures pagination, filtering, and
// sorting when listing artifacts produced by a workflow run.
//
// Zero-value fields are omitted from the request, allowing GitHub's default
// behavior to be used.
type ListWorkflowRunArtifactsOptions struct {
	// Page is the page number of results to retrieve.
	//
	// Page numbering starts at 1.
	//
	// Default: 1.
	Page int `query:"page"`

	// PerPage is the maximum number of artifacts to return per page.
	//
	// Valid values range from 1 to 100.
	//
	// Default: 30.
	PerPage int `query:"per_page"`

	// Name filters the returned artifacts by their name.
	//
	// When specified, only artifacts whose name exactly matches this value
	// are returned.
	Name string `query:"name"`

	// Direction specifies the order in which artifacts are returned.
	//
	// Valid values are:
	//   - "asc"  : Oldest artifacts first.
	//   - "desc" : Newest artifacts first.
	//
	// Default: "desc".
	Direction SortDirection `query:"direction" query-default:"desc"`
}

// ListRunsForWorkflowOptions configures filtering when listing runs
// for a specific workflow.
type ListRunsForWorkflowOptions struct {
	// Actor filters workflow runs by the login of the user who created
	// the push associated with the check suite or workflow run.
	Actor string `query:"actor"`

	// Branch filters workflow runs associated with a branch.
	Branch string `query:"branch"`

	// Event filters workflow runs triggered by the specified event.
	Event string `query:"event"`

	// Status filters workflow runs by their check run status or conclusion.
	Status string `query:"status"`

	// PerPage is the number of results per page (max 100).
	//
	// Default: 30.
	PerPage int64 `query:"per_page"`

	// Page is the page number of the results to fetch.
	//
	// Default: 1.
	Page int64 `query:"page"`

	// Created returns workflow runs created within the given date-time range.
	Created string `query:"created"`

	// ExcludePullRequests, when true, omits pull requests from the response.
	//
	// Default: false.
	ExcludePullRequests bool `query:"exclude_pull_requests"`

	// CheckSuiteID filters workflow runs by the specified check suite identifier.
	CheckSuiteID int64 `query:"check_suite_id"`

	// HeadSHA filters workflow runs to only those associated with the
	// specified commit SHA.
	HeadSHA string `query:"head_sha"`
}

// ListConcurrencyGroupsForRepoOptions configures pagination when listing
// concurrency groups for a repository.
type ListConcurrencyGroupsForRepoOptions struct {
	// PerPage is the number of results per page (max 100).
	//
	// Default: 30.
	PerPage int `query:"per_page"`

	// After is a cursor for forward-only pagination. If specified, the
	// query only searches for results after this cursor.
	After string `query:"after"`
}

// GetConcurrencyGroupForRepoOptions filters the queue contents of a
// concurrency group.
//
// AheadOfRun and AheadOfJob are mutually exclusive.
type GetConcurrencyGroupForRepoOptions struct {
	// AheadOfRun filters to items ahead of this workflow run ID in
	// the queue, plus the run itself. Mutually exclusive with
	// AheadOfJob.
	AheadOfRun int64 `query:"ahead_of_run"`

	// AheadOfJob filters to items ahead of this job ID in the queue,
	// plus the job itself. Mutually exclusive with AheadOfRun.
	AheadOfJob int64 `query:"ahead_of_job"`
}

// ListConcurrencyGroupsForWorkflowOptions configures pagination when
// listing concurrency groups for a workflow run.
type ListConcurrencyGroupsForWorkflowOptions struct {
	// PerPage is the number of results per page (max 100).
	//
	// Default: 30.
	PerPage int `query:"per_page"`

	// Before is a cursor for backward-only pagination. If specified,
	// the query only searches for results before this cursor.
	Before string `query:"before"`

	// After is a cursor for forward-only pagination. If specified, the
	// query only searches for results after this cursor.
	After string `query:"after"`
}
