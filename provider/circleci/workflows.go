package circleci

import (
	"context"
	"fmt"
)

// GetWorkflow retrieves a workflow by its ID.
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - workflowID: The unique workflow ID.
//
// Returns the requested workflow or an error if the request fails, the workflow
// does not exist, or the API returns a non-success status.
func (c *Client) GetWorkflow(ctx context.Context, workflowID string) (Workflow, error) {
	path := fmt.Sprintf("/workflow/%s", workflowID)

	var resp Workflow
	if err := c.do(ctx, "GET", path, nil, &resp); err != nil {
		return Workflow{}, fmt.Errorf("failed to get workflow: %w", err)
	}
	return resp, nil
}

// ListWorkflowJobs retrieves all jobs for a workflow.
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - workflowID: The unique workflow ID.
//   - opts: Optional pagination configuration.
//
// Returns a slice of jobs or an error if the request fails or the API
// returns a non-success status.
func (c *Client) ListWorkflowJobs(ctx context.Context, workflowID string, opts *ListWorkflowJobsOptions) ([]Job, error) {
	path := fmt.Sprintf("/workflow/%s/job", workflowID)
	path = appendQueryParams(path, opts)

	var resp jobListResponse
	if err := c.do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list workflow jobs: %w", err)
	}
	return resp.Items, nil
}

// CancelWorkflow cancels all jobs in a workflow.
//
// The caller must have write access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - workflowID: The unique workflow ID.
//
// Returns an error if the request fails or the API returns a non-success status.
func (c *Client) CancelWorkflow(ctx context.Context, workflowID string) error {
	path := fmt.Sprintf("/workflow/%s/cancel", workflowID)

	if err := c.do(ctx, "POST", path, nil, nil); err != nil {
		return fmt.Errorf("failed to cancel workflow: %w", err)
	}
	return nil
}

// ApproveWorkflow approves a job in a workflow.
//
// This is used to approve jobs that require manual approval before continuing.
//
// The caller must have write access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - workflowID: The unique workflow ID.
//   - approvalName: The name of the approval job.
//
// Returns an error if the request fails or the API returns a non-success status.
func (c *Client) ApproveWorkflow(ctx context.Context, workflowID, approvalName string) error {
	path := fmt.Sprintf("/workflow/%s/approve/%s", workflowID, approvalName)

	if err := c.do(ctx, "POST", path, nil, nil); err != nil {
		return fmt.Errorf("failed to approve workflow: %w", err)
	}
	return nil
}
