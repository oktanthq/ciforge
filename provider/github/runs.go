package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/abdulahshoaib/ciforge/models"
)

func (c *Client) ListRuns(ctx context.Context, repo string, pipelineID string) ([]models.Run, error) {
	path := fmt.Sprintf("/repos/%s/actions/workflows/%s/runs", repo, pipelineID)
	var resp listWorkflowRunsResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list workflow runs: %w", err)
	}
	return mapWorkflowRuns(resp.WorkflowRuns), nil
}

func (c *Client) GetRuns(ctx context.Context, repo string, runID string) (*models.Run, error) {
	path := fmt.Sprintf("/repos/%s/actions/runs/%s", repo, runID)
	var resp workflowRunDTO
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to get workflow run: %w", err)
	}
	r := mapWorkflowRun(resp)
	return &r, nil
}

func (c *Client) CancelRun(ctx context.Context, repo string, runID string) error {
	path := fmt.Sprintf("/repos/%s/actions/runs/%s/cancel", repo, runID)
	if err := c.do(ctx, http.MethodPost, path, nil, nil); err != nil {
		return fmt.Errorf("failed to cancel workflow run: %w", err)
	}
	return nil
}

func (c *Client) RetryRun(ctx context.Context, repo string, runID string) error {
	path := fmt.Sprintf("/repos/%s/actions/runs/%s/rerun", repo, runID)
	if err := c.do(ctx, http.MethodPost, path, nil, nil); err != nil {
		return fmt.Errorf("failed to retry workflow run: %w", err)
	}
	return nil
}
