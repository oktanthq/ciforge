package circleci

import (
	"context"
	"fmt"
	"net/http"
)

// ListEnvironments retrieves all deployment environments for an organization.
//
// The caller must have read access to the organization.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - opts: Filtering and pagination options.
//
// Returns a slice of environments or an error if the request fails or the API
// returns a non-success status.
func (c *Client) ListEnvironments(ctx context.Context, opts ListEnvironmentsOptions) ([]Environment, error) {
	path := "/deploy/environments"

	var resp listEnvironmentResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return []Environment{}, fmt.Errorf("failed to get deployment environment: %w", err)
	}
	return resp.Items, nil
}

// GetEnvironment retrieves a deployment environment by its ID.
//
// The caller must have read access to the organization.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - envID: The unique environment ID.
//
// Returns the requested environment or an error if the request fails, the
// environment does not exist, or the API returns a non-success status.
func (c *Client) GetEnvironment(ctx context.Context, envID string) (Environment, error) {
	path := fmt.Sprintf("/deploy/environments/%s", envID)

	var resp Environment
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return Environment{}, fmt.Errorf("failed to get environment{%s}: %w", envID, err)
	}
	return resp, nil
}

// ListComponents retrieves all deployment components for an organization or project.
//
// The caller must have read access to the organization or project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - opts: Filtering and pagination options.
//
// Returns a slice of components or an error if the request fails or the API
// returns a non-success status.
func (c *Client) ListComponents(ctx context.Context, opts ListComponentsOptions) ([]Component, error) {
	path := "/deploy/component"

	var resp listComponentsResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return []Component{}, fmt.Errorf("failed to get components: %w", err)
	}
	return resp.Items, nil
}

// GetComponents retrieves a deployment component by its ID.
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - componentID: The unique component ID.
//
// Returns the requested component or an error if the request fails, the
// component does not exist, or the API returns a non-success status.
func (c *Client) GetComponents(ctx context.Context, componentID string) (Component, error) {
	path := fmt.Sprintf("/deploy/component/%s", componentID)

	var resp Component
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return Component{}, fmt.Errorf("failed to get component{%s}: %w", componentID, err)
	}
	return resp, nil
}
