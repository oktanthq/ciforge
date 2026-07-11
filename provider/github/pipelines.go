package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/abdulahshoaib/ciforge/models"
)

// ListWorkflows retrieves all GitHub Actions workflows for a repository.
//
// The caller must have read access to the repository. Each workflow is
// returned as a provider-agnostic Pipeline model.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//
// Returns a slice of pipelines or an error if the request fails or the API
// returns a non-success status.
//
// TODO: Support pagination.
func (c *Client) ListWorkflows(ctx context.Context, repo, owner string) ([]models.Pipeline, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/workflows", owner, repo)

	var resp listWorkflowsResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list workflows: %w", err)
	}
	return mapWorkflows(resp.Workflows), nil
}

// GetWorkflow retrieves a GitHub Actions workflow by its ID or file name.
//
// GitHub accepts either the numeric workflow ID or the workflow file name
// (for example, "build.yml").
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - id: Workflow identifier or workflow file name.
//
// Returns the requested workflow or an error if the request fails, the
// workflow does not exist, or the API returns a non-success status.
func (c *Client) GetWorkflow(ctx context.Context, repo, owner, id string) (*models.Pipeline, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/workflows/%s", owner, repo, id)
	var resp workflowDTO
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to get workflow: %w", err)
	}
	p := mapWorkflow(resp)
	return &p, nil
}

// DisableWorkflow disables a GitHub Actions workflow.
//
// Once disabled, the workflow will no longer be triggered automatically or
// manually until it is re-enabled.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - workflowID: Identifier or file name of the workflow to disable.
//
// Returns an error if the request fails or the API returns a non-success
// status.
func (c *Client) DisableWorkflow(ctx context.Context, repo, owner, workflowID string) error {
	path := fmt.Sprintf("/repos/%s/%s/actions/workflows/%s/disable", owner, repo, workflowID)

	if err := c.do(ctx, http.MethodPut, path, nil, nil); err != nil {
		return fmt.Errorf("failed to disable workflow: %w", err)
	}

	return nil
}

// CreateWorkflowDispatch manually triggers a GitHub Actions workflow.
//
// The workflow is dispatched against the specified Git reference. If the
// workflow defines inputs, they may be supplied in the request body.
//
// On success, GitHub returns information about the newly created workflow
// run, including its identifier and API and web URLs.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - workflowID: Identifier or file name of the workflow to dispatch.
//   - body: Dispatch configuration containing the target Git reference and
//     optional workflow inputs.
//
// Returns the dispatched workflow run or an error if the request fails or
// the API returns a non-success status.
func (c *Client) CreateWorkflowDispatch(ctx context.Context, repo, owner, workflowID string, body workflowDispatchRequest) (models.DispatchEvent, error) {
	path := fmt.Sprintf("repos/%s/%s/actions/workflows/%s/dispatches", owner, repo, workflowID)

	var resp workflowDispatchResponse
	if err := c.do(ctx, http.MethodPost, path, body, &resp); err != nil {
		return models.DispatchEvent{}, fmt.Errorf("failed in creating a dispatch event: %w", err)
	}

	return mapDispatchEvent(resp), nil
}

// EnableWorkflow enables a previously disabled GitHub Actions workflow.
//
// Once enabled, the workflow may be triggered automatically or manually
// according to its configured events.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - workflowID: Identifier or file name of the workflow to enable.
//
// Returns an error if the request fails or the API returns a non-success
// status.
func (c *Client) EnableWorkflow(ctx context.Context, repo, owner, workflowID string) error {
	path := fmt.Sprintf("/repos/%s/%s/actions/workflows/%s/enable", owner, repo, workflowID)

	if err := c.do(ctx, http.MethodPut, path, nil, nil); err != nil {
		return fmt.Errorf("failed to enable workflow: %w", err)
	}

	return nil
}
