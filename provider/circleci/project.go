package circleci

import (
	"context"
	"fmt"
	"net/http"
)

// CreateProject creates a new project under an organization.
//
// The caller must have write access to the organization.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - orgSlugID: The organization slug (e.g., "gh/CircleCI-Public").
//   - body: The project configuration including name.
//
// Returns the created project or an error if the request fails or the API
// returns a non-success status.
func (c *Client) CreateProject(ctx context.Context, orgSlugID string, body CreateProjectBody) (Project, error) {
	path := fmt.Sprintf("/organization/%s/project", orgSlugID)

	var resp Project
	if err := c.do(ctx, http.MethodPost, path, body, &resp); err != nil {
		return Project{}, fmt.Errorf("failed to create project for Org{%s}: %w", orgSlugID, err)
	}
	return resp, nil
}

// DeleteProject deletes a project by its slug.
//
// The caller must have write access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//
// Returns an error if the request fails or the API returns a non-success status.
func (c *Client) DeleteProject(ctx context.Context, projectSlug string) error {
	path := fmt.Sprintf("/project/%s", projectSlug)

	if err := c.do(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("failed to delete project{%s}: %w", projectSlug, err)
	}
	return nil
}

// GetProject retrieves a project by its slug.
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//
// Returns the requested project or an error if the request fails, the project
// does not exist, or the API returns a non-success status.
func (c *Client) GetProject(ctx context.Context, projectSlug string) (Project, error) {
	path := fmt.Sprintf("/project/%s", projectSlug)

	var resp Project
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return Project{}, fmt.Errorf("failed to get project{%s}: %w", projectSlug, err)
	}
	return resp, nil
}

// GetAllCheckoutKeys returns all checkout keys for a project.
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//
// Returns a slice of checkout keys or an error if the request fails or the API
// returns a non-success status.
func (c *Client) GetAllCheckoutKeys(ctx context.Context, projectSlug string) ([]CheckoutKey, error) {
	path := fmt.Sprintf("/project/%s/checkout-key", projectSlug)

	var resp getAllCheckoutKeysResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return []CheckoutKey{}, fmt.Errorf("failed to get all checkout-keys for Project{%s}: %w", projectSlug, err)
	}
	return resp.Items, nil
}

// CreateCheckoutKey creates a new checkout key for a project.
//
// The caller must have write access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//   - body: The checkout key configuration including type.
//
// Returns the created checkout key or an error if the request fails or the API
// returns a non-success status.
func (c *Client) CreateCheckoutKey(ctx context.Context, projectSlug string, body CreateCheckoutKeyBody) (CheckoutKey, error) {
	path := fmt.Sprintf("/project/%s/checkout-key", projectSlug)

	var resp CheckoutKey
	if err := c.do(ctx, http.MethodPost, path, body, &resp); err != nil {
		return CheckoutKey{}, fmt.Errorf("failed to create checkout key for Project{%s}: %w", projectSlug, err)
	}
	return resp, nil
}

// DeleteCheckoutKey deletes a checkout key by its fingerprint.
//
// The caller must have write access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//   - fingerprint: The SSH key fingerprint to delete.
//
// Returns an error if the request fails or the API returns a non-success status.
func (c *Client) DeleteCheckoutKey(ctx context.Context, projectSlug, fingerprint string) error {
	path := fmt.Sprintf("/project/%s/checkout-key/%s", projectSlug, fingerprint)

	if err := c.do(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("failed to delete checkout-key for Project{%s}: %w", projectSlug, err)
	}
	return nil
}

// GetCheckoutKey retrieves a checkout key by its fingerprint.
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//   - fingerprint: The SSH key fingerprint.
//
// Returns the requested checkout key or an error if the request fails, the key
// does not exist, or the API returns a non-success status.
func (c *Client) GetCheckoutKey(ctx context.Context, projectSlug, fingerprint string) (CheckoutKey, error) {
	path := fmt.Sprintf("/project/%s/checkout-key/%s", projectSlug, fingerprint)

	var resp CheckoutKey
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return CheckoutKey{}, fmt.Errorf("failed to get checkout key for Project{%s}: %w", projectSlug, err)
	}
	return resp, nil
}

// ListEvironmentVariables returns all environment variables for a project.
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//
// Returns a slice of environment variables or an error if the request fails or the API
// returns a non-success status.
func (c *Client) ListEvironmentVariables(ctx context.Context, projectSlug string) ([]ProjectEnvVar, error) {
	path := fmt.Sprintf("/project/%s/envvar", projectSlug)

	var resp listProjectEnvVarsResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return []ProjectEnvVar{}, fmt.Errorf("failed to get Project{%s} Environment Variable: %w", projectSlug, err)
	}
	return resp.Items, nil
}

// CreateEnvironmentVariables creates a new environment variable for a project.
//
// The caller must have write access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//   - body: The environment variable configuration including name and value.
//
// Returns the created environment variable or an error if the request fails or the API
// returns a non-success status.
func (c *Client) CreateEnvironmentVariables(ctx context.Context, projectSlug string, body CreateProjectEnvVarBody) (ProjectEnvVar, error) {
	path := fmt.Sprintf("/project/%s/envvar", projectSlug)

	var resp ProjectEnvVar
	if err := c.do(ctx, http.MethodPost, path, body, &resp); err != nil {
		return ProjectEnvVar{}, fmt.Errorf("failed to create Environment Variable for Project{%s}: %w", projectSlug, err)
	}
	return resp, nil
}

// DeleteEnvironmentVariables deletes an environment variable from a project.
//
// The caller must have write access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//   - name: The name of the environment variable to delete.
//
// Returns an error if the request fails or the API returns a non-success status.
func (c *Client) DeleteEnvironmentVariables(ctx context.Context, projectSlug, name string) error {
	path := fmt.Sprintf("/project/%s/envvar/%s", projectSlug, name)

	if err := c.do(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("failed to delete Environment Variable{%s} for Project{%s}: %w", name, projectSlug, err)
	}
	return nil
}

// GetMaskedEnvironmentVariable retrieves a masked environment variable by name.
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//   - name: The name of the environment variable.
//
// Returns the requested environment variable or an error if the request fails, the variable
// does not exist, or the API returns a non-success status.
func (c *Client) GetMaskedEnvironmentVariable(ctx context.Context, projectSlug, name string) (ProjectEnvVar, error) {
	path := fmt.Sprintf("/project/%s/envvar/%s", projectSlug, name)

	var resp ProjectEnvVar
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return ProjectEnvVar{}, fmt.Errorf("failed to get masked Environment Variable{%s} for Project{%s}: %w", name, projectSlug, err)
	}
	return resp, nil
}

// GetProjectSettings retrieves the settings for a project.
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - provider: The VCS provider ("github"/"gh", "bitbucket"/"bb", or "circleci").
//   - organization: The organization name or ID.
//   - project: The project name or ID.
//
// Returns the project settings or an error if the request fails or the API
// returns a non-success status.
func (c *Client) GetProjectSettings(ctx context.Context, provider, organization, project string) (ProjectSettings, error) {
	path := fmt.Sprintf("/project/%s/%s/%s/settings", provider, organization, project)

	var resp projectSettingsResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return ProjectSettings{}, fmt.Errorf("failed to get Project Settings for{%s/%s}: %w", organization, project, err)
	}
	return resp.Advanced, nil
}

// UpdateProjectSettings updates the settings for a project.
//
// The caller must have write access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - provider: The VCS provider ("github"/"gh", "bitbucket"/"bb", or "circleci").
//   - organization: The organization name or ID.
//   - project: The project name or ID.
//   - body: The updated project settings.
//
// Returns the updated project settings or an error if the request fails or the API
// returns a non-success status.
func (c *Client) UpdateProjectSettings(ctx context.Context, provider, organization, project string, body UpdateProjectSettingsBody) (ProjectSettings, error) {
	path := fmt.Sprintf("/project/%s/%s/%s/settings", provider, organization, project)

	var resp projectSettingsResponse
	if err := c.do(ctx, http.MethodPatch, path, body, &resp); err != nil {
		return ProjectSettings{}, fmt.Errorf("failed to updated Project Settings{%s/%s}: %w", organization, project, err)
	}
	return resp.Advanced, nil
}
