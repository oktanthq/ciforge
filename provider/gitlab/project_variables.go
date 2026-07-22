package gitlab

import (
	"context"
	"fmt"
	"net/http"
)

// ListProjectVariables retrieves all CI/CD variables for a project.
//
// Use the page and per_page pagination parameters to control the pagination
// of results.
//
// The caller must have at least maintainer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - opts: Optional filtering and pagination parameters.
//
// Returns a slice of project variables or an error if the request fails or
// the API returns a non-success status.
func (c *Client) ListProjectVariables(ctx context.Context, projectID int, opts *ListProjectVariablesOptions) ([]ProjectVariable, error) {
	path := fmt.Sprintf("/projects/%d/variables", projectID)
	path = appendQueryParams(path, opts)

	var resp []ProjectVariable
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list project variables: %w", err)
	}
	return resp, nil
}

// GetProjectVariable retrieves details of a single project variable.
//
// If there are multiple variables with the same key, use the filter option
// to select the correct environment_scope.
//
// The caller must have at least maintainer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - key: The key of the variable.
//   - opts: Optional filter parameters.
//
// Returns the requested project variable or an error if the request fails,
// the variable does not exist, or the API returns a non-success status.
func (c *Client) GetProjectVariable(ctx context.Context, projectID int, key string, opts *GetProjectVariableOptions) (ProjectVariable, error) {
	path := fmt.Sprintf("/projects/%d/variables/%s", projectID, key)
	path = appendQueryParams(path, opts)

	var resp ProjectVariable
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return ProjectVariable{}, fmt.Errorf("failed to get project variable: %w", err)
	}
	return resp, nil
}

// CreateProjectVariable creates a new CI/CD variable for a project.
//
// If a variable with the same key already exists, the new variable must have
// a different environment_scope. Otherwise, GitLab returns an error similar to:
// "VARIABLE_NAME has already been taken".
//
// The caller must have at least maintainer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - opts: The variable configuration including key and value.
//
// Returns the newly created project variable or an error if the request fails
// or the API returns a non-success status.
func (c *Client) CreateProjectVariable(ctx context.Context, projectID int, opts *CreateProjectVariableBody) (ProjectVariable, error) {
	path := fmt.Sprintf("/projects/%d/variables", projectID)

	var resp ProjectVariable
	if err := c.do(ctx, http.MethodPost, path, opts, &resp); err != nil {
		return ProjectVariable{}, fmt.Errorf("failed to create project variable: %w", err)
	}
	return resp, nil
}

// UpdateProjectVariable updates a project CI/CD variable.
//
// If there are multiple variables with the same key, use the filter option
// to select the correct environment_scope.
//
// The caller must have at least maintainer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - key: The key of the variable.
//   - opts: The fields to update.
//
// Returns the updated project variable or an error if the request fails or
// the API returns a non-success status.
func (c *Client) UpdateProjectVariable(ctx context.Context, projectID int, key string, opts *UpdateProjectVariableBody) (ProjectVariable, error) {
	path := fmt.Sprintf("/projects/%d/variables/%s", projectID, key)

	var resp ProjectVariable
	if err := c.do(ctx, http.MethodPut, path, opts, &resp); err != nil {
		return ProjectVariable{}, fmt.Errorf("failed to update project variable: %w", err)
	}
	return resp, nil
}

// DeleteProjectVariable deletes a project CI/CD variable.
//
// If there are multiple variables with the same key, use the filter option
// to select the correct environment_scope.
//
// The caller must have at least maintainer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - key: The key of the variable.
//   - opts: Optional filter parameters.
//
// Returns an error if the deletion fails or the API returns a non-success
// status.
func (c *Client) DeleteProjectVariable(ctx context.Context, projectID int, key string, opts *DeleteProjectVariableOptions) error {
	path := fmt.Sprintf("/projects/%d/variables/%s", projectID, key)
	path = appendQueryParams(path, opts)

	if err := c.do(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("failed to delete project variable: %w", err)
	}
	return nil
}
