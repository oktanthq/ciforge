package gitlab

import (
	"context"
	"fmt"
	"net/http"
)

// ListRunners retrieves all runners available to the user.
//
// Prerequisites:
//   - For group runners, you must have the Owner role in the owner namespace.
//   - For project runners, you must have the Security Manager, Maintainer or
//     Owner role in a project assigned to the runner.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - opts: Optional filtering parameters.
//
// Returns a slice of runners or an error if the request fails or the API
// returns a non-success status.
func (c *Client) ListRunners(ctx context.Context, opts *ListRunnersOptions) ([]Runner, error) {
	path := "/runners"
	path = appendQueryParams(path, opts)

	var resp []Runner
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list runners: %w", err)
	}
	return resp, nil
}

// ListAllRunners retrieves all runners in the GitLab instance.
//
// This endpoint requires administrator or auditor access.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - opts: Optional filtering parameters.
//
// Returns a slice of runners or an error if the request fails or the API
// returns a non-success status.
func (c *Client) ListAllRunners(ctx context.Context, opts *ListRunnersOptions) ([]Runner, error) {
	path := "/runners/all"
	path = appendQueryParams(path, opts)

	var resp []Runner
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list all runners: %w", err)
	}
	return resp, nil
}

// GetRunner retrieves details of a runner.
//
// Instance runner details are available to all authenticated users.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - runnerID: The ID of the runner.
//
// Returns the requested runner or an error if the request fails, the runner
// does not exist, or the API returns a non-success status.
func (c *Client) GetRunner(ctx context.Context, runnerID int) (Runner, error) {
	path := fmt.Sprintf("/runners/%d", runnerID)

	var resp Runner
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return Runner{}, fmt.Errorf("failed to get runner: %w", err)
	}
	return resp, nil
}

// UpdateRunner updates details of a runner.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - runnerID: The ID of the runner.
//   - opts: The fields to update.
//
// Returns the updated runner or an error if the request fails or the API
// returns a non-success status.
func (c *Client) UpdateRunner(ctx context.Context, runnerID int, opts *UpdateRunnerRequest) (Runner, error) {
	path := fmt.Sprintf("/runners/%d", runnerID)

	var resp Runner
	if err := c.do(ctx, http.MethodPut, path, opts, &resp); err != nil {
		return Runner{}, fmt.Errorf("failed to update runner: %w", err)
	}
	return resp, nil
}

// ListRunnerJobs lists all jobs that are being processed or were processed
// by the specified runner.
//
// The list of jobs is limited to projects where the user has the Reporter,
// Developer, Maintainer, or Owner role.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - runnerID: The ID of the runner.
//   - opts: Optional filtering parameters.
//
// Returns a slice of jobs or an error if the request fails or the API
// returns a non-success status.
func (c *Client) ListRunnerJobs(ctx context.Context, runnerID int, opts *ListRunnerJobsOptions) ([]RunnerJob, error) {
	path := fmt.Sprintf("/runners/%d/jobs", runnerID)
	path = appendQueryParams(path, opts)

	var resp []RunnerJob
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list runner jobs: %w", err)
	}
	return resp, nil
}

// ListRunnerManagers lists all the managers of a runner.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - runnerID: The ID of the runner.
//
// Returns a slice of runner managers or an error if the request fails or the
// API returns a non-success status.
func (c *Client) ListRunnerManagers(ctx context.Context, runnerID int) ([]RunnerManager, error) {
	path := fmt.Sprintf("/runners/%d/managers", runnerID)

	var resp []RunnerManager
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list runner managers: %w", err)
	}
	return resp, nil
}

// ListProjectRunners lists all runners available in the project, including
// from ancestor groups and any allowed instance runners.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - opts: Optional filtering parameters.
//
// Returns a slice of runners or an error if the request fails or the API
// returns a non-success status.
func (c *Client) ListProjectRunners(ctx context.Context, projectID int, opts *ListRunnersOptions) ([]Runner, error) {
	path := fmt.Sprintf("/projects/%d/runners", projectID)
	path = appendQueryParams(path, opts)

	var resp []Runner
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list project runners: %w", err)
	}
	return resp, nil
}

// AssignRunnerToProject assigns an available project runner to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - runnerID: The ID of the runner to assign.
//
// Returns the assigned runner or an error if the request fails or the API
// returns a non-success status.
func (c *Client) AssignRunnerToProject(ctx context.Context, projectID, runnerID int) (Runner, error) {
	path := fmt.Sprintf("/projects/%d/runners", projectID)

	body := map[string]int{"runner_id": runnerID}

	var resp Runner
	if err := c.do(ctx, http.MethodPost, path, body, &resp); err != nil {
		return Runner{}, fmt.Errorf("failed to assign runner to project: %w", err)
	}
	return resp, nil
}

// UnassignRunnerFromProject unassigns a project runner from the project.
//
// You cannot unassign a runner from the owner project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - runnerID: The ID of the runner to unassign.
//
// Returns an error if the unassignment fails or the API returns a non-success
// status.
func (c *Client) UnassignRunnerFromProject(ctx context.Context, projectID, runnerID int) error {
	path := fmt.Sprintf("/projects/%d/runners/%d", projectID, runnerID)

	if err := c.do(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("failed to unassign runner from project: %w", err)
	}
	return nil
}

// ListGroupRunners lists all runners available in the group and its ancestor
// groups, including any allowed instance runners.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - groupID: The ID of the group.
//   - opts: Optional filtering parameters.
//
// Returns a slice of runners or an error if the request fails or the API
// returns a non-success status.
func (c *Client) ListGroupRunners(ctx context.Context, groupID int, opts *ListRunnersOptions) ([]Runner, error) {
	path := fmt.Sprintf("/groups/%d/runners", groupID)
	path = appendQueryParams(path, opts)

	var resp []Runner
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list group runners: %w", err)
	}
	return resp, nil
}

// CreateRunner creates a new runner with a registration token.
//
// This endpoint returns an HTTP 410 Gone status code if registration with
// runner registration tokens is disabled in the project or group settings.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - opts: The runner configuration including token and optional settings.
//
// Returns the created runner token response or an error if the request fails
// or the API returns a non-success status.
func (c *Client) CreateRunner(ctx context.Context, opts *CreateRunnerRequest) (RunnerTokenResponse, error) {
	path := "/runners"

	var resp RunnerTokenResponse
	if err := c.do(ctx, http.MethodPost, path, opts, &resp); err != nil {
		return RunnerTokenResponse{}, fmt.Errorf("failed to create runner: %w", err)
	}
	return resp, nil
}

// DeleteRunner deletes a runner by ID.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - runnerID: The ID of the runner to delete.
//
// Returns an error if the deletion fails or the API returns a non-success
// status.
func (c *Client) DeleteRunner(ctx context.Context, runnerID int) error {
	path := fmt.Sprintf("/runners/%d", runnerID)

	if err := c.do(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("failed to delete runner: %w", err)
	}
	return nil
}

// DeleteRunnerByToken deletes a runner using its authentication token.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - token: The runner's authentication token.
//
// Returns an error if the deletion fails or the API returns a non-success
// status.
func (c *Client) DeleteRunnerByToken(ctx context.Context, token string) error {
	path := "/runners"

	body := map[string]string{"token": token}

	if err := c.do(ctx, http.MethodDelete, path, body, nil); err != nil {
		return fmt.Errorf("failed to delete runner by token: %w", err)
	}
	return nil
}

// VerifyRunner validates authentication credentials for a registered runner.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - token: The runner's authentication token.
//   - systemID: Optional system ID (required if token starts with "glrt-").
//
// Returns the runner token response or an error if the request fails or the
// API returns a non-success status.
func (c *Client) VerifyRunner(ctx context.Context, token, systemID string) (RunnerTokenResponse, error) {
	path := "/runners/verify"

	body := map[string]string{"token": token}
	if systemID != "" {
		body["system_id"] = systemID
	}

	var resp RunnerTokenResponse
	if err := c.do(ctx, http.MethodPost, path, body, &resp); err != nil {
		return RunnerTokenResponse{}, fmt.Errorf("failed to verify runner: %w", err)
	}
	return resp, nil
}

// ResetInstanceRegistrationToken resets the runner registration token for
// the GitLab instance.
//
// This endpoint requires administrator access.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//
// Returns the new registration token or an error if the request fails or the
// API returns a non-success status.
func (c *Client) ResetInstanceRegistrationToken(ctx context.Context) (string, error) {
	path := "/runners/reset_registration_token"

	var resp struct {
		Token string `json:"token"`
	}
	if err := c.do(ctx, http.MethodPost, path, nil, &resp); err != nil {
		return "", fmt.Errorf("failed to reset instance registration token: %w", err)
	}
	return resp.Token, nil
}

// ResetProjectRegistrationToken resets the runner registration token for a
// project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//
// Returns the new registration token or an error if the request fails or the
// API returns a non-success status.
func (c *Client) ResetProjectRegistrationToken(ctx context.Context, projectID int) (string, error) {
	path := fmt.Sprintf("/projects/%d/runners/reset_registration_token", projectID)

	var resp struct {
		Token string `json:"token"`
	}
	if err := c.do(ctx, http.MethodPost, path, nil, &resp); err != nil {
		return "", fmt.Errorf("failed to reset project registration token: %w", err)
	}
	return resp.Token, nil
}

// ResetGroupRegistrationToken resets the runner registration token for a
// group.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - groupID: The ID of the group.
//
// Returns the new registration token or an error if the request fails or the
// API returns a non-success status.
func (c *Client) ResetGroupRegistrationToken(ctx context.Context, groupID int) (string, error) {
	path := fmt.Sprintf("/groups/%d/runners/reset_registration_token", groupID)

	var resp struct {
		Token string `json:"token"`
	}
	if err := c.do(ctx, http.MethodPost, path, nil, &resp); err != nil {
		return "", fmt.Errorf("failed to reset group registration token: %w", err)
	}
	return resp.Token, nil
}

// ResetRunnerAuthenticationToken resets the runner's authentication token
// using its runner ID.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - runnerID: The ID of the runner.
//
// Returns the new authentication token response or an error if the request
// fails or the API returns a non-success status.
func (c *Client) ResetRunnerAuthenticationToken(ctx context.Context, runnerID int) (RunnerTokenResponse, error) {
	path := fmt.Sprintf("/runners/%d/reset_authentication_token", runnerID)

	var resp RunnerTokenResponse
	if err := c.do(ctx, http.MethodPost, path, nil, &resp); err != nil {
		return RunnerTokenResponse{}, fmt.Errorf("failed to reset runner authentication token: %w", err)
	}
	return resp, nil
}

// ResetRunnerAuthenticationTokenByToken resets the runner's authentication
// token using the current token's value.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - token: The current authentication token of the runner.
//
// Returns the new authentication token response or an error if the request
// fails or the API returns a non-success status.
func (c *Client) ResetRunnerAuthenticationTokenByToken(ctx context.Context, token string) (RunnerTokenResponse, error) {
	path := "/runners/reset_authentication_token"

	body := map[string]string{"token": token}

	var resp RunnerTokenResponse
	if err := c.do(ctx, http.MethodPost, path, body, &resp); err != nil {
		return RunnerTokenResponse{}, fmt.Errorf("failed to reset runner authentication token by token: %w", err)
	}
	return resp, nil
}

// GetJobRouterDiscovery retrieves Job Router discovery information for a
// runner.
//
// This endpoint requires a valid runner authentication token.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//
// Returns the job router discovery information or an error if the request
// fails or the API returns a non-success status.
func (c *Client) GetJobRouterDiscovery(ctx context.Context) (JobRouterDiscovery, error) {
	path := "/runners/router/discovery"

	var resp JobRouterDiscovery
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return JobRouterDiscovery{}, fmt.Errorf("failed to get job router discovery: %w", err)
	}
	return resp, nil
}
