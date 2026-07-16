package github

import (
	"context"
	"fmt"
	"net/http"
)

// ListWorkflows retrieves all GitHub Actions workflows for a repository.
//
// The caller must have read access to the repository. Each workflow is
// returned as a Workflow value.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//
// Returns a slice of workflows or an error if the request fails or the API
// returns a non-success status.
//
// TODO: Support pagination.
func (c *Client) ListWorkflows(ctx context.Context, repo, owner string) ([]Workflow, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/workflows", owner, repo)

	var resp workflowListResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list workflows: %w", err)
	}
	return resp.Workflows, nil
}

// GetWorkflow retrieves a GitHub Actions workflow by its ID.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - workflowID: Workflow identifier.
//
// Returns the requested workflow or an error if the request fails, the
// workflow does not exist, or the API returns a non-success status.
func (c *Client) GetWorkflow(ctx context.Context, repo, owner string, workflowID int64) (Workflow, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/workflows/%d", owner, repo, workflowID)
	var resp Workflow
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return Workflow{}, fmt.Errorf("failed to get workflow: %w", err)
	}
	return resp, nil
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
//   - workflowID: Identifier of the workflow to disable.
//
// Returns an error if the request fails or the API returns a non-success
// status.
func (c *Client) DisableWorkflow(ctx context.Context, repo, owner string, workflowID int64) error {
	path := fmt.Sprintf("/repos/%s/%s/actions/workflows/%d/disable", owner, repo, workflowID)

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
//   - workflowID: Identifier of the workflow to dispatch.
//   - body: Dispatch configuration containing the target Git reference and
//     optional workflow inputs.
//
// Returns the dispatched workflow event or an error if the request fails or
// the API returns a non-success status.
func (c *Client) CreateWorkflowDispatch(ctx context.Context, repo, owner string, workflowID int64, body WorkflowDispatchRequest) (WorkflowDispatchEvent, error) {
	path := fmt.Sprintf("repos/%s/%s/actions/workflows/%d/dispatches", owner, repo, workflowID)

	var resp workflowDispatchResponse
	if err := c.do(ctx, http.MethodPost, path, body, &resp); err != nil {
		return WorkflowDispatchEvent{}, fmt.Errorf("failed in creating a dispatch event: %w", err)
	}

	return WorkflowDispatchEvent{
		RunID:  resp.WorkflowRunID,
		RunURL: resp.RunURL,
		URL:    resp.HTMLURL,
	}, nil
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
//   - workflowID: Identifier of the workflow to enable.
//
// Returns an error if the request fails or the API returns a non-success
// status.
func (c *Client) EnableWorkflow(ctx context.Context, repo, owner string, workflowID int64) error {
	path := fmt.Sprintf("/repos/%s/%s/actions/workflows/%d/enable", owner, repo, workflowID)

	if err := c.do(ctx, http.MethodPut, path, nil, nil); err != nil {
		return fmt.Errorf("failed to enable workflow: %w", err)
	}

	return nil
}

// GetWorkflowUsage retrieves the billable usage for a GitHub Actions workflow.
//
// The returned usage is grouped by runner operating system and includes the
// total billable execution time for each runner type across all runs of the
// specified workflow.
//
// The caller must have read access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - workflowID: Identifier of the workflow.
//
// Returns the billable usage for the specified workflow or an error if the
// request fails or the API returns a non-success status.
func (c *Client) GetWorkflowUsage(ctx context.Context, repo, owner string, workflowID int64) (BillableWorkflow, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/workflows/%d/timing", owner, repo, workflowID)

	var resp BillableWorkflowResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return BillableWorkflow{}, fmt.Errorf("failed to get usage for workflow{%d}: %w", workflowID, err)
	}
	return resp.Billable, nil
}
