package github

import (
	"net/url"
	"reflect"
	"strconv"
)

// SortDirection specifies the ordering of list results supported by the GitHub API.
type SortDirection string

const (
	SortAscending  SortDirection = "asc"
	SortDescending SortDirection = "desc"
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

type ListRunsForWorkflowOpts struct {
	// Returns someone's workflow runs. Use the login for the user who created
	// the push associated with the check suite or workflow run.
	Actor string `query:"actor"`

	// Returns workflow runs associated with a branch. Use the name of the branch
	// of the push.
	Branch string `query:"branch"`

	// Returns workflow run triggered by the event you specify. For example, push,
	// pull_request or issue.
	Event string `query:"event"`

	// Returns workflow runs with the check run status or conclusion that you specify.
	// For example, a conclusion can be success or a status can be in_progress. Only
	// GitHub Actions can set a status of waiting, pending, or requested.

	// Can be one of: completed, action_required, cancelled, failure, neutral, skipped,
	// stale, success, timed_out, in_progress, queued, requested, waiting, pending
	Status string `query:"status"`

	// The number of results per page (max 100)
	// Default: 30
	PerPage int64 `query:"per_page"`

	// The page number of the results to fetch.
	// Default: 1
	Page int64 `query:"page"`

	// Returns workflow runs created within the given date-time range.
	Created string `query:"created"`

	// If true pull requests are omitted from the response (empty array).
	// Default: false
	ExcludePullRequests bool `query:"exclude_pull_requests"`

	// Returns workflow runs with the check_suite_id that you specify.
	CheckSuiteID int64 `query:"check_suite_id"`

	//Only returns workflow runs that are associated with the specified head_sha.
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
