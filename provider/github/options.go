package github

import (
	"net/url"
	"strconv"
)

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
	Page int

	// PerPage is the maximum number of artifacts to return per page.
	//
	// Valid values range from 1 to 100.
	//
	// Default: 30.
	PerPage int

	// Name filters the returned artifacts by their name.
	//
	// When specified, only artifacts whose name exactly matches this value
	// are returned.
	Name string
}

// Append appends the configured query parameters to the provided API path.
//
// Only fields with non-zero values are included in the resulting query
// string. If no options are configured, the original path is returned
// unchanged.
func (o *ListArtifactsOptions) Append(path string) string {
	if o == nil {
		return path
	}

	q := url.Values{}

	if o.Page > 0 {
		q.Set("page", strconv.Itoa(o.Page))
	}
	if o.PerPage > 0 {
		q.Set("per_page", strconv.Itoa(o.PerPage))
	}
	if o.Name != "" {
		q.Set("name", o.Name)
	}

	if len(q) == 0 {
		return path
	}

	return path + "?" + q.Encode()
}

// SortDirection specifies the ordering of list results supported by the GitHub API.
type SortDirection string

const (
	SortAscending  SortDirection = "asc"
	SortDescending SortDirection = "desc"
)

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
	Page int

	// PerPage is the maximum number of artifacts to return per page.
	//
	// Valid values range from 1 to 100.
	//
	// Default: 30.
	PerPage int

	// Name filters the returned artifacts by their name.
	//
	// When specified, only artifacts whose name exactly matches this value
	// are returned.
	Name string

	// Direction specifies the order in which artifacts are returned.
	//
	// Valid values are:
	//   - "asc"  : Oldest artifacts first.
	//   - "desc" : Newest artifacts first.
	//
	// Default: "desc".
	Direction SortDirection
}

// AppendQueryParams appends the configured query parameters to the provided API path.
//
// Only fields with non-zero values are included in the resulting query
// string. If no options are configured, the original path is returned
// unchanged.
func (o *ListWorkflowRunArtifactsOptions) AppendQueryParams(path string) string {
	if o == nil {
		return path
	}

	q := url.Values{}

	if o.Page > 0 {
		q.Set("page", strconv.Itoa(o.Page))
	}
	if o.PerPage > 0 {
		q.Set("per_page", strconv.Itoa(o.PerPage))
	}
	if o.Name != "" {
		q.Set("name", o.Name)
	}

	if o.Direction == SortAscending {
		q.Set("direction", string(SortAscending))
	} else {
		q.Set("direction", string(SortDescending))
	}

	if len(q) == 0 {
		return path
	}

	return path + "?" + q.Encode()
}
