package github

import (
	"context"
	"fmt"
	"net/http"
)

// ListRunsForRepo retrieves all GitHub Actions workflow runs for a repository.
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
func (c *Client) ListRunsForRepo(ctx context.Context, repo, owner string, opts ListRunsForRepoOptions) ([]WorkflowRun, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/runs", owner, repo)
	appendQueryParams(path, opts)

	var resp workflowRunListResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list workflow runs: %w", err)
	}
	return resp.WorkflowRuns, nil
}

// GetWorkflowRun retrieves a single GitHub Actions workflow run by its identifier.
//
// The caller must have read access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - runID: Numeric identifier of the workflow run.
//
// Returns the requested workflow run or an error if the request fails, the run
// does not exist, or the API returns a non-success status.
func (c *Client) GetWorkflowRun(ctx context.Context, repo, owner string, runID int64) (WorkflowRun, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/runs/%d", owner, repo, runID)

	var resp WorkflowRun
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return WorkflowRun{}, fmt.Errorf("failed to get workflow run: %w", err)
	}
	return resp, nil
}

// DeleteWorkflowRun permanently deletes a GitHub Actions workflow run.
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
//   - runID: Numeric identifier of the workflow run.
//
// Returns an error if the deletion fails or the API returns a non-success
// status.
func (c *Client) DeleteWorkflowRun(ctx context.Context, repo, owner string, runID int64) error {
	path := fmt.Sprintf("/repos/%s/%s/actions/runs/%d", owner, repo, runID)

	if err := c.do(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("failed to delete workflow run: %w", err)
	}
	return nil
}

// GetWorkflowRunReviewHistory retrieves the approval history for a workflow run
// that requires environment protection reviews.
//
// The returned approvals include reviewers and their approval decisions.
//
// The caller must have read access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - runID: Numeric identifier of the workflow run.
//
// Returns the workflow approval history or an error if the request fails or the
// API returns a non-success status.
func (c *Client) GetWorkflowRunReviewHistory(ctx context.Context, repo, owner string, runID int64) ([]WorkflowApproval, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/runs/%d/approvals", owner, repo, runID)

	var resp []WorkflowApproval
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return []WorkflowApproval{}, fmt.Errorf("failed to get review history: %w", err)
	}

	return resp, nil
}

// ApproveWorflowRunForForkPR approves a GitHub Actions workflow run triggered
// by a pull request from a public fork submitted by a first-time contributor.
//
// GitHub may require manual approval before workflows from first-time
// contributors are allowed to execute. This method approves the pending
// workflow run.
//
// The caller must have write access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - runID: Numeric identifier of the workflow run to approve.
//
// Returns an error if the approval request fails or the API returns a
// non-success status.
func (c *Client) ApproveWorflowRunForForkPR(ctx context.Context, repo, owner string, runID int64) error {
	path := fmt.Sprintf("/repos/%s/%s/actions/runs/%d/approve", owner, repo, runID)

	if err := c.do(ctx, http.MethodPost, path, nil, nil); err != nil {
		return fmt.Errorf("failed to approve Run{%d} for a fork PR: %w", runID, err)
	}
	return nil
}

// GetWorkflowRunAttempt retrieves a specific execution attempt of a workflow run.
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
//   - runID: Numeric identifier of the workflow run.
//   - attemptNum: Numeric attempt number to retrieve.
//
// Returns the requested workflow run attempt or an error if the request fails
// or the API returns a non-success status.
func (c *Client) GetWorkflowRunAttempt(ctx context.Context, repo, owner string, runID, attemptNum int64) (WorkflowRun, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/runs/%d/attempts/%d", owner, repo, runID, attemptNum)

	var resp WorkflowRun
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return WorkflowRun{}, fmt.Errorf("failed to get run attempts: %w", err)
	}

	return resp, nil
}

// DownloadWorkflowRunAttemptLogs downloads the log archive for a specific
// GitHub Actions workflow run attempt.
//
// GitHub returns a temporary redirect to a downloadable archive containing the
// logs for the requested workflow run attempt. The download URL expires after
// a short period, but this method automatically follows the redirect and
// returns the downloaded archive contents.
//
// The caller must have read access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - runID: Numeric identifier of the workflow run.
//   - attemptNum: Numeric attempt number of the workflow run.
//
// Returns the downloaded log archive as a byte slice or an error if the request
// fails or the API returns a non-success status.
func (c *Client) DownloadWorkflowRunAttemptLogs(ctx context.Context, repo, owner string, runID, attemptNum int64) ([]byte, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/runs/%d/attempts/%d/logs", owner, repo, runID, attemptNum)

	data, err := c.doRaw(ctx, http.MethodGet, path)
	if err != nil {
		return nil, fmt.Errorf("failed to download workflow run logs: %w", err)
	}

	return data, nil
}

// CancelWorkflowRun requests cancellation of an in-progress workflow run.
//
// If the run has already completed, the request has no effect.
//
// The caller must have write access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - runID: Numeric identifier of the workflow run.
//
// Returns an error if the cancellation request fails or the API returns a
// non-success status.
func (c *Client) CancelWorkflowRun(ctx context.Context, repo, owner string, runID int64) error {
	path := fmt.Sprintf("/repos/%s/%s/actions/runs/%d/cancel", owner, repo, runID)

	if err := c.do(ctx, http.MethodPost, path, nil, nil); err != nil {
		return fmt.Errorf("failed to cancel workflow run: %w", err)
	}
	return nil
}

// ForceCancelWorkflowRun forcefully cancels a workflow run that could not be
// cancelled through the standard cancellation endpoint.
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
//   - runID: Numeric identifier of the workflow run.
//
// Returns an error if the force cancellation request fails or the API returns
// a non-success status.
func (c *Client) ForceCancelWorkflowRun(ctx context.Context, repo, owner string, runID int64) error {
	path := fmt.Sprintf("/repos/%s/%s/actions/runs/%d/force-cancel", owner, repo, runID)

	if err := c.do(ctx, http.MethodPost, path, nil, nil); err != nil {
		return fmt.Errorf("failed to force cancel workflow run: %w", err)
	}

	return nil
}

// DownloadWorkflowRunLogs downloads the logs generated by a workflow run.
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
//   - runID: Numeric identifier of the workflow run.
//
// Returns an error if the request fails or the API returns a non-success
// status.
func (c *Client) DownloadWorkflowRunLogs(ctx context.Context, repo, owner string, runID int64) error {
	path := fmt.Sprintf("/repos/%s/%s/actions/runs/%d/logs", owner, repo, runID)

	if err := c.do(ctx, http.MethodGet, path, nil, nil); err != nil {
		return fmt.Errorf("failed to get run logs: %w", err)
	}
	return nil
}

// DeleteWorkflowRunLogs permanently deletes the logs associated with a workflow run.
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
//   - runID: Numeric identifier of the workflow run.
//
// Returns an error if the request fails or the API returns a non-success
// status.
func (c *Client) DeleteWorkflowRunLogs(ctx context.Context, repo, owner string, runID int64) error {
	path := fmt.Sprintf("/repos/%s/%s/actions/runs/%d/logs", owner, repo, runID)

	if err := c.do(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("failed to delete workflow run logs: %w", err)
	}
	return nil
}

// GetPendingDeployments retrieves all pending deployment requests for a
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
//   - runID: Numeric identifier of the workflow run.
//
// Returns the pending deployments or an error if the request fails or the API
// returns a non-success status.
func (c *Client) GetPendingDeployments(ctx context.Context, repo, owner string, runID int64) ([]PendingDeployment, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/runs/%d/pending_deployments", owner, repo, runID)

	var resp []PendingDeployment
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to get pending deployments on given run{%d}: %w", runID, err)
	}
	return resp, nil
}

// ReviewPendingDeploymentsBody defines the review submitted for one or more
// pending deployments in a GitHub Actions workflow run.
//
// The review specifies the target deployment environments, the review decision,
// and an optional comment explaining the decision.
type ReviewPendingDeploymentsBody struct {
	EnvironmentIDs []int64 `json:"environment_ids"`
	State          string  `json:"state"`
	Comment        string  `json:"comment,omitempty"`
}

// ReviewPendingDeployments reviews one or more pending deployments for a
// GitHub Actions workflow run.
//
// This endpoint approves or rejects pending deployments waiting for deployment
// protection rules. The request body specifies the target environments, the
// review state, and an optional review comment.
//
// The caller must have write access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - runID: Numeric identifier of the workflow run.
//   - body: Review details, including the target environment IDs, review
//     decision, and optional comment.
//
// Returns a slice of deployments updated as a result of the review or an error
// if the request fails or the API returns a non-success status.
func (c *Client) ReviewPendingDeployments(ctx context.Context, repo, owner string, runID int64, body ReviewPendingDeploymentsBody) ([]Deployment, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/runs/%d/pending_deployments", owner, repo, runID)

	var resp []Deployment
	if err := c.do(ctx, http.MethodPost, path, body, &resp); err != nil {
		return nil, fmt.Errorf("failed to review pending deployments on workflow run{%d}: %w", runID, err)
	}
	return resp, nil
}

// ReRunWorkflow reruns an entire GitHub Actions workflow run.
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
//   - runID: Numeric identifier of the workflow run.
//
// Returns an error if the rerun request fails or the API returns a non-success
// status.
func (c *Client) ReRunWorkflow(ctx context.Context, repo, owner string, runID int64) error {
	path := fmt.Sprintf("/repos/%s/%s/actions/runs/%d/rerun", owner, repo, runID)

	if err := c.do(ctx, http.MethodPost, path, nil, nil); err != nil {
		return fmt.Errorf("failed to rerun workflow run{%d}: %w", runID, err)
	}
	return nil
}

// ReRunFailedJobs reruns only the failed jobs within a workflow run.
//
// Successful jobs are not executed again.
//
// The caller must have write access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - runID: Numeric identifier of the workflow run.
//
// Returns an error if the rerun request fails or the API returns a non-success
// status.
func (c *Client) ReRunFailedJobs(ctx context.Context, repo, owner string, runID int64) error {
	path := fmt.Sprintf("/repos/%s/%s/actions/runs/%d/rerun-failed-jobs", owner, repo, runID)

	if err := c.do(ctx, http.MethodPost, path, nil, nil); err != nil {
		return fmt.Errorf("failed to rerun failed workflow run{%d}: %w", runID, err)
	}
	return nil
}

// GetWorkflowRunUsage retrieves the billable usage for a GitHub Actions
// workflow run.
//
// The returned usage is grouped by runner operating system and includes the
// total billable execution time, number of jobs, and per-job execution
// durations where available.
//
// The caller must have read access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - runID: Numeric identifier of the workflow run.
//
// Returns the billable usage for the specified workflow run or an error if the
// request fails or the API returns a non-success status.
func (c *Client) GetWorkflowRunUsage(ctx context.Context, repo, owner string, runID int64) (Billable, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/runs/%d/timing", owner, repo, runID)

	var resp BillableRunResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return Billable{}, fmt.Errorf("failed to get workflow run{%d} usage: %w", runID, err)
	}

	return resp.Billable, nil
}

// ListRunsForWorkflow retrieves all workflow runs for a specific GitHub Actions
// workflow.
//
// Unlike ListRunsForRepo, which returns runs across all workflows in the
// repository, this method only returns runs associated with the specified
// workflow.
//
// The caller must have read access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - workflowID: Numeric identifier of the workflow.
//
// Returns a slice of workflow runs associated with the specified workflow or an
// error if the request fails or the API returns a non-success status.
func (c *Client) ListRunsForWorkflow(ctx context.Context, repo, owner string, workflowID int64, opts ListRunsForWorkflowOptions) ([]WorkflowRun, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/workflows/%d/runs", owner, repo, workflowID)
	appendQueryParams(path, opts)

	var resp workflowRunListResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return []WorkflowRun{}, fmt.Errorf("failed to get runs for workflow{%d}: %w", workflowID, err)
	}
	return resp.WorkflowRuns, nil
}
