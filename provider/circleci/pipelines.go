package circleci

import (
	"context"
	"fmt"
)

// ListPipelines retrieves all pipelines for a project.
//
// The project slug must be in the form of "vcs/org/repo" (e.g., "gh/CircleCI-Public/circleci-cli").
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//   - opts: Optional list configuration for filtering and pagination.
//
// Returns a slice of pipelines or an error if the request fails or the API
// returns a non-success status.
func (c *Client) ListPipelines(ctx context.Context, projectSlug string, opts *ListPipelinesOptions) ([]Pipeline, error) {
	path := fmt.Sprintf("/project/%s/pipeline", projectSlug)
	path = appendQueryParams(path, opts)

	var resp pipelineListResponse
	if err := c.do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list pipelines: %w", err)
	}
	return resp.Items, nil
}

// GetPipelineByNumber retrieves a pipeline by its number within a project.
//
// The project slug must be in the form of "vcs/org/repo" (e.g., "gh/CircleCI-Public/circleci-cli").
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//   - pipelineNumber: The pipeline number.
//   - opts: Optional configuration for filtering.
//
// Returns the requested pipeline or an error if the request fails, the pipeline
// does not exist, or the API returns a non-success status.
func (c *Client) GetPipelineByNumber(ctx context.Context, projectSlug string, pipelineNumber int64, opts *GetPipelineOptions) (Pipeline, error) {
	path := fmt.Sprintf("/project/%s/pipeline/%d", projectSlug, pipelineNumber)
	path = appendQueryParams(path, opts)

	var resp Pipeline
	if err := c.do(ctx, "GET", path, nil, &resp); err != nil {
		return Pipeline{}, fmt.Errorf("failed to get pipeline: %w", err)
	}
	return resp, nil
}

// GetPipelineByID retrieves a pipeline by its unique ID.
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - pipelineID: The unique pipeline ID.
//
// Returns the requested pipeline or an error if the request fails, the pipeline
// does not exist, or the API returns a non-success status.
func (c *Client) GetPipelineByID(ctx context.Context, pipelineID string) (Pipeline, error) {
	path := fmt.Sprintf("/pipeline/%s", pipelineID)

	var resp Pipeline
	if err := c.do(ctx, "GET", path, nil, &resp); err != nil {
		return Pipeline{}, fmt.Errorf("failed to get pipeline: %w", err)
	}
	return resp, nil
}

// TriggerPipeline triggers a new pipeline for a project.
//
// The project slug must be in the form of "vcs/org/repo" (e.g., "gh/CircleCI-Public/circleci-cli").
//
// The caller must have write access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//   - body: The trigger configuration including branch/tag and parameters.
//
// Returns the newly created pipeline or an error if the request fails or the
// API returns a non-success status.
func (c *Client) TriggerPipeline(ctx context.Context, projectSlug string, body TriggerPipelineBody) (Pipeline, error) {
	path := fmt.Sprintf("/project/%s/pipeline", projectSlug)

	var resp Pipeline
	if err := c.do(ctx, "POST", path, body, &resp); err != nil {
		return Pipeline{}, fmt.Errorf("failed to trigger pipeline: %w", err)
	}
	return resp, nil
}

// ListPipelineWorkflows retrieves all workflows for a pipeline.
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - pipelineID: The unique pipeline ID.
//   - opts: Optional pagination configuration.
//
// Returns a slice of workflows or an error if the request fails or the API
// returns a non-success status.
func (c *Client) ListPipelineWorkflows(ctx context.Context, pipelineID string, opts *ListPipelineWorkflowsOptions) ([]Workflow, error) {
	path := fmt.Sprintf("/pipeline/%s/workflow", pipelineID)
	path = appendQueryParams(path, opts)

	var resp workflowListResponse
	if err := c.do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list pipeline workflows: %w", err)
	}
	return resp.Items, nil
}
