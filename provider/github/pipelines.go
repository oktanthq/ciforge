package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/abdulahshoaib/ciforge/models"
)

func (c *Client) ListPipelines(ctx context.Context, repo string) ([]models.Pipeline, error) {
	path := fmt.Sprintf("/repos/%s/actions/workflows", repo)
	var resp listWorkflowsResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list workflows: %w", err)
	}
	return mapWorkflows(resp.Workflows), nil
}

func (c *Client) GetPipeline(ctx context.Context, repo string, id string) (*models.Pipeline, error) {
	path := fmt.Sprintf("/repos/%s/actions/workflows/%s", repo, id)
	var resp workflowDTO
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to get workflow: %w", err)
	}
	p := mapWorkflow(resp)
	return &p, nil
}
