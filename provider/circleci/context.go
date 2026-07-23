package circleci

import (
	"context"
	"fmt"
	"net/http"
)

// CreateContext creates a new context for storing environment variables.
//
// Contexts are used to share environment variables and configuration across
// multiple projects within an organization.
//
// The caller must have write access to the organization.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - body: The context configuration including name and owner.
//
// Returns the created context or an error if the request fails or the API
// returns a non-success status.
func (c *Client) CreateContext(ctx context.Context, body ContextBody) (Context, error) {
	path := "/context"

	var resp Context
	if err := c.do(ctx, http.MethodPost, path, body, &resp); err != nil {
		return Context{}, fmt.Errorf("failed to create context: %w", err)
	}
	return resp, nil
}

// ListContexts returns all contexts for the given organization.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - opts: Filtering options including owner ID/slug and pagination.
//
// Returns the list of contexts or an error if the request fails.
func (c *Client) ListContexts(ctx context.Context, opts ListContextOptions) ([]Context, error) {
	path := "/context"
	appendQueryParams(path, opts)

	var resp listContextResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return []Context{}, fmt.Errorf("failed to list contexts: %w", err)
	}
	return resp.Items, nil
}

// GetContext returns a single context by its ID.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - contextID: The unique identifier of the context.
//
// Returns the context or an error if the request fails.
func (c *Client) GetContext(ctx context.Context, contextID string) (Context, error) {
	path := fmt.Sprintf("/context/%s", contextID)

	var resp Context
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return Context{}, fmt.Errorf("failed to get context{%s}: %w", contextID, err)
	}
	return resp, nil
}

// DeleteContext deletes a context by its ID.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - contextID: The unique identifier of the context to delete.
//
// Returns an error if the request fails.
func (c *Client) DeleteContext(ctx context.Context, contextID string) error {
	path := fmt.Sprintf("/context/%s", contextID)

	if err := c.do(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("failed to delete context: %w", err)
	}
	return nil
}

// ListEnvironmentVariables returns all environment variables for a context.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - contextID: The unique identifier of the context.
//   - opts: Pagination options.
//
// Returns the list of environment variables or an error if the request fails.
func (c *Client) ListEnvironmentVariables(ctx context.Context, contextID string, opts ListEnvVarOptions) ([]EnvVar, error) {
	path := fmt.Sprintf("/context/%s/environment-variable", contextID)
	appendQueryParams(path, opts)

	var resp listEnvVarsResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return []EnvVar{}, fmt.Errorf("failed to list environment variables: %w", err)
	}
	return resp.Items, nil
}

// UpsertEnviornmentVariable creates or updates an environment variable in a context.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - contextID: The unique identifier of the context.
//   - envVarName: The name of the environment variable.
//   - body: The new value for the environment variable.
//
// Returns the upserted environment variable or an error if the request fails.
func (c *Client) UpsertEnviornmentVariable(ctx context.Context, contextID, envVarName string, body UpsertEnvVarBody) (EnvVar, error) {
	path := fmt.Sprintf("/context/%s/environment-variable/%s", contextID, envVarName)

	var resp EnvVar
	if err := c.do(ctx, http.MethodPut, path, body, &resp); err != nil {
		return EnvVar{}, fmt.Errorf("failed to upsert environment variables: %w", err)
	}
	return resp, nil
}

// RemoveEnviornmentVariable deletes an environment variable from a context.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - contextID: The unique identifier of the context.
//   - envVarName: The name of the environment variable to remove.
//
// Returns an error if the request fails.
func (c *Client) RemoveEnviornmentVariable(ctx context.Context, contextID, envVarName string) error {
	path := fmt.Sprintf("/context/%s/environment-variable/%s", contextID, envVarName)

	if err := c.do(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("failed to upsert environment variables: %w", err)
	}
	return nil
}

// GetContextRestrictions returns all restrictions for a context.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - contextID: The unique identifier of the context.
//
// Returns the list of restrictions or an error if the request fails.
func (c *Client) GetContextRestrictions(ctx context.Context, contextID string) ([]Restriction, error) {
	path := fmt.Sprintf("/context/%s/restrictions", contextID)

	var resp listContextRestrictionResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return []Restriction{}, fmt.Errorf("failed to get context{%s} restrictions: %w", contextID, err)
	}
	return resp.Items, nil
}

// CreateContextRestrictions creates a new restriction for a context.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - contextID: The unique identifier of the context.
//   - body: The restriction configuration including type and value.
//
// Returns the created restriction or an error if the request fails.
func (c *Client) CreateContextRestrictions(ctx context.Context, contextID string, body CreateContextRestrictionsBody) (Restriction, error) {
	path := fmt.Sprintf("/context/%s/restrictions", contextID)

	var resp Restriction
	if err := c.do(ctx, http.MethodPost, path, body, &resp); err != nil {
		return Restriction{}, fmt.Errorf("failed to create context{%s} restrictions: %w", contextID, err)
	}
	return resp, nil
}

// DeleteContextRestrictions deletes a restriction from a context.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - contextID: The unique identifier of the context.
//   - restrictionID: The unique identifier of the restriction to delete.
//
// Returns an error if the request fails.
func (c *Client) DeleteContextRestrictions(ctx context.Context, contextID, restrictionID string) error {
	path := fmt.Sprintf("/context/%s/restrictions/%s", contextID, restrictionID)

	if err := c.do(ctx, http.MethodPost, path, nil, nil); err != nil {
		return fmt.Errorf("failed to delete context{%s} restriction{%s}: %w", contextID, restrictionID, err)
	}
	return nil
}
