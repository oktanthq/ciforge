package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/abdulahshoaib/ciforge/models"
)

// ListRuns retrieves all GitHub Actions workflow runs for a repository.
//
// By default, GitHub returns workflow runs across all workflows in the
// repository. Future versions may support filtering through query parameters.
//
// The caller must have read access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//
// Returns a slice of workflow runs or an error if the request fails or the API
// returns a non-success status.
func (c *Client) ListRuns(ctx context.Context, repo, owner string) ([]models.Run, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/runs", owner, repo)
	// TODO(abdullah): Add query params

	var resp listWorkflowRunsResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list workflow runs: %w", err)
	}
	return mapWorkflowRuns(resp.WorkflowRuns), nil
}

// GetRuns retrieves a single GitHub Actions workflow run by its identifier.
//
// The caller must have read access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - runID: Identifier of the workflow run.
//
// Returns the requested workflow run or an error if the request fails, the run
// does not exist, or the API returns a non-success status.
func (c *Client) GetRuns(ctx context.Context, repo, owner, runID string) (models.Run, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/runs/%s", repo, runID)
	// TODO(abdullah): Add query params

	var resp workflowRunDTO
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return models.Run{}, fmt.Errorf("failed to get workflow run: %w", err)
	}
	return mapWorkflowRun(resp), nil
}

// DeleteRun permanently deletes a GitHub Actions workflow run.
//
// Deleting a workflow run also removes its associated logs and metadata. This
// operation cannot be undone.
//
// The caller must have write access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - runID: Identifier of the workflow run.
//
// Returns an error if the deletion fails or the API returns a non-success
// status.
func (c *Client) DeleteRun(ctx context.Context, repo, owner, runID string) error {
	path := fmt.Sprintf("repos/%s/%s/actions/runs/%s", owner, repo, runID)

	if err := c.do(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("failed to delete workflow run: %w", err)
	}
	return nil
}

// GetReviewHistory retrieves the approval history for a workflow run that
// requires environment protection reviews.
//
// The returned approvals include reviewers and their approval decisions.
//
// The caller must have read access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - runID: Identifier of the workflow run.
//
// Returns the workflow approval history or an error if the request fails or the
// API returns a non-success status.
// TODO(abdulah): create a strategy to store the workflow approvals
func (c *Client) GetReviewHistory(ctx context.Context, repo, owner, runID string) ([]workflowApprovalDTO, error) {
	path := fmt.Sprintf("repos/%s/%s/actions/runs/%s/approvals", owner, repo, runID)

	var resp []workflowApprovalDTO
	if err := c.do(ctx, http.MethodGet, path, nil, resp); err != nil {
		return []workflowApprovalDTO{}, fmt.Errorf("failed to get review history: %w", err)
	}

	return resp, nil
}

// GetRunAttempt retrieves a specific execution attempt of a workflow run.
//
// Workflow runs may have multiple attempts when rerun. This method returns the
// details for the requested attempt number.
//
// The caller must have read access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - runID: Identifier of the workflow run.
//   - attemptNum: Attempt number to retrieve.
//
// Returns the requested workflow run attempt or an error if the request fails
// or the API returns a non-success status.
func (c *Client) GetRunAttempt(ctx context.Context, repo, owner, runID, attemptNum string) (models.Run, error) {
	path := fmt.Sprintf("repos/%s/%s/actions/runs/%s/attempts/%s", owner, repo, runID, attemptNum)

	var resp workflowRunDTO
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return models.Run{}, fmt.Errorf("failed to get run attempts: %w", err)
	}

	return mapWorkflowRun(resp), nil
}

// CancelRun requests cancellation of an in-progress workflow run.
//
// If the run has already completed, the request has no effect.
//
// The caller must have write access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - runID: Identifier of the workflow run.
//
// Returns an error if the cancellation request fails or the API returns a
// non-success status.
func (c *Client) CancelRun(ctx context.Context, repo, owner, runID string) error {
	path := fmt.Sprintf("/repos/%s/%s/actions/runs/%s/cancel", owner, repo, runID)

	if err := c.do(ctx, http.MethodPost, path, nil, nil); err != nil {
		return fmt.Errorf("failed to cancel workflow run: %w", err)
	}
	return nil
}

// ForceCancelRun forcefully cancels a workflow run that could not be cancelled
// through the standard cancellation endpoint.
//
// This endpoint is intended for runs that are stuck or unable to terminate
// normally.
//
// The caller must have write access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - runID: Identifier of the workflow run.
//
// Returns an error if the force cancellation request fails or the API returns
// a non-success status.
func (c *Client) ForceCancelRun(ctx context.Context, repo, owner, runID string) error {
	path := fmt.Sprintf("repos/%s/%s/actions/runs/%s/force-cancel", owner, repo, runID)

	if err := c.do(ctx, http.MethodPost, path, nil, nil); err != nil {
		return fmt.Errorf("failed to force cancel workflow run: %w", err)
	}

	return nil
}

// DownloadRunLogs downloads the logs generated by a workflow run.
//
// GitHub responds with a redirect to a downloadable archive containing the run
// logs.
//
// The caller must have read access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - runID: Identifier of the workflow run.
//
// Returns an error if the request fails or the API returns a non-success
// status.
func (c *Client) DownloadRunLogs(ctx context.Context, repo, owner, runID string) error {
	path := fmt.Sprintf("repos/%s/%s/actions/runs/%s/logs", owner, repo, runID)

	if err := c.do(ctx, http.MethodGet, path, nil, nil); err != nil {
		return fmt.Errorf("failed to get run logs: %w", err)
	}
	return nil
}

// DeleteRunLogs permanently deletes the logs associated with a workflow run.
//
// This operation only removes the logs and does not delete the workflow run
// itself.
//
// The caller must have write access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - runID: Identifier of the workflow run.
//
// Returns an error if the request fails or the API returns a non-success
// status.
func (c *Client) DeleteRunLogs(ctx context.Context, repo, owner, runID string) error {
	path := fmt.Sprintf("repos/%s/%s/actions/runs/%s/logs", owner, repo, runID)

	if err := c.do(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("failed to delete workflow run logs: %w", err)
	}
	return nil
}

// GetPendingRunDeployments retrieves all pending deployment requests for a
// workflow run.
//
// Pending deployments occur when a workflow is waiting for approval or other
// deployment protection rules before continuing.
//
// The caller must have read access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - runID: Identifier of the workflow run.
//
// Returns an error if the request fails or the API returns a non-success
// status.
func (c *Client) GetPendingRunDeployments(ctx context.Context, repo, owner, runID string) error {
	// TODO(abdullah): handle the incoming data
	path := fmt.Sprintf("repos/%s/%s/actions/runs/%s/pending_deployments", owner, repo, runID)

	// TODO(abdullah): there is a response to this request, make sure to capture that later
	if err := c.do(ctx, http.MethodGet, path, nil, nil); err != nil {
		return fmt.Errorf("failed to get pending deployments on given run{%s}: %w", runID, err)
	}
	return nil
}

// ReRunFailedJob reruns only the failed jobs within a workflow run.
//
// Successful jobs are not executed again.
//
// The caller must have write access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - runID: Identifier of the workflow run.
//
// Returns an error if the rerun request fails or the API returns a non-success
// status.
func (c *Client) ReRunFailedJob(ctx context.Context, repo, owner, runID string) error {
	path := fmt.Sprintf("repos/%s/%s/actions/runs/%s/rerun-failed-jobs", owner, repo, runID)

	if err := c.do(ctx, http.MethodPost, path, nil, nil); err != nil {
		return fmt.Errorf("failed to rerun failed workflow run{%s}: %w", runID, err)
	}
	return nil
}

// ReRun reruns an entire GitHub Actions workflow run.
//
// All jobs in the workflow are executed again regardless of their previous
// status.
//
// The caller must have write access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - runID: Identifier of the workflow run.
//
// Returns an error if the rerun request fails or the API returns a non-success
// status.
func (c *Client) ReRun(ctx context.Context, repo, owner, runID string) error {
	path := fmt.Sprintf("/repos/%s/%s/actions/run/%s/rerun", owner, repo, runID)

	if err := c.do(ctx, http.MethodPost, path, nil, nil); err != nil {
		return fmt.Errorf("failed to rerun workflow run{%s}: %w", runID, err)
	}
	return nil
}

// ListWorkflowRuns retrieves all workflow runs for a specific GitHub Actions
// workflow.
//
// Unlike ListRuns, which returns runs across all workflows in the repository,
// this method only returns runs associated with the specified workflow.
//
// The caller must have read access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - workflowID: Identifier or filename of the workflow.
//
// Returns a slice of workflow runs associated with the specified workflow or an
// error if the request fails or the API returns a non-success status.
func (c *Client) ListWorkflowRuns(ctx context.Context, repo, owner, workflowID string) ([]models.Run, error) {
	path := fmt.Sprintf("repos/OWNER/REPO/actions/workflows/WORKFLOW_ID/runs", owner, repo, workflowID)

	var resp listWorkflowRunsResponse
	if err := c.do(ctx, http.MethodGet, path, nil, nil); err != nil {
		return []models.Run{}, fmt.Errorf("failed to get runs for workflow{%s}: %w", err)
	}
	return mapWorkflowRuns(resp.WorkflowRuns), nil
}
