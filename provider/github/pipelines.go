package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/abdulahshoaib/ciforge/models"
)

func (c *Client) ListPipelines(ctx context.Context, repo, owner string) ([]models.Pipeline, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/workflows", owner, repo)
	// TODO(abdullah): add query params

	var resp listWorkflowsResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list workflows: %w", err)
	}
	return mapWorkflows(resp.Workflows), nil
}

func (c *Client) GetPipeline(ctx context.Context, repo, owner, id string) (*models.Pipeline, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/workflows/%s", owner, repo, id)
	var resp workflowDTO
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to get workflow: %w", err)
	}
	p := mapWorkflow(resp)
	return &p, nil
}

func (c *Client) DisablePipeline(ctx context.Context, repo, owner, workflowID string) error {
	path := fmt.Sprintf("/repos/%s/%s/actions/workflows/%s/disable", owner, repo, workflowID)

	if err := c.do(ctx, http.MethodPut, path, nil, nil); err != nil {
		return fmt.Errorf("failed to disable workflow: %w", err)
	}

	return nil
}

func (c *Client) CreateDispatchEvent(ctx context.Context, repo, owner, workflowID string, body workflowDispatchRequest) error {
	path := fmt.Sprintf("repos/%s/%s/actions/workflows/%s/dispatches", owner, repo, workflowID)

	var resp workflowDispatchResponse
	if err := c.do(ctx, http.MethodPost, path, body, &resp); err != nil {
		return fmt.Errorf("failed in creating a dispatch event: %w", err)
	}

	return nil
}

func (c *Client) EnablePipeline(ctx context.Context, repo, owner, workflowID string) error {
	path := fmt.Sprintf("repos/%s/%s/actions/workflows/%s/enable", owner, repo, workflowID)

	if err := c.do(ctx, http.MethodPut, path, nil, nil); err != nil {
		return fmt.Errorf("failed to enable pipeline: %w", err)
	}

	return nil
}
