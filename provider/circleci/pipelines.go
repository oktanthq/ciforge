package circleci

import (
	"context"
	"fmt"
	"net/http"
)

// ListPipelines retrieves all pipelines for an organization. (max 250)
//
// The org slug is the first two segments of the project slug: "{provider}/{organization}"
// (e.g., "gh/CircleCI-Public").
//
// The caller must have read access to the organization.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - opts: Configuration for filtering and pagination. OrgSlug is required.
//
// Returns a slice of pipelines or an error if the request fails or the API
// returns a non-success status.
func (c *Client) ListPipelines(ctx context.Context, opts *ListPipelinesOptions) ([]Pipeline, error) {
	path := "/pipeline"
	path = appendQueryParams(path, opts)

	var resp pipelineListResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list pipelines: %w", err)
	}
	return resp.Items, nil
}

// ContinuePipeline continues a pipeline that requires approval.
//
// This is used to continue a pipeline that was set up with an approval step.
// The pipeline must be in a "hold" status to be continued.
//
// The caller must have write access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - body: The continue configuration including the pipeline ID and parameters.
//
// Returns an error if the request fails or the API returns a non-success status.
func (c *Client) ContinuePipeline(ctx context.Context, body *ContinuePipelineBody) error {
	path := "/pipeline/continue"

	if err := c.do(ctx, http.MethodPost, path, body, nil); err != nil {
		return fmt.Errorf("failed to continue pipeline: %w", err)
	}
	return nil
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
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return Pipeline{}, fmt.Errorf("failed to get pipeline: %w", err)
	}
	return resp, nil
}

// GetPipelineConfig retrieves the configuration for a pipeline.
//
// This returns the source, compiled, and setup configurations for the pipeline.
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - pipelineID: The unique pipeline ID.
//
// Returns the pipeline configuration or an error if the request fails or the API
// returns a non-success status.
func (c *Client) GetPipelineConfig(ctx context.Context, pipelineID string) (PipelineConfig, error) {
	path := fmt.Sprintf("/pipeline/%s/config", pipelineID)

	var resp PipelineConfig
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return PipelineConfig{}, fmt.Errorf("failed to get pipeline config: %w", err)
	}
	return resp, nil
}

// ListPipelineWorkflows retrieves all workflows for a pipeline.
//
// This returns the workflows associated with a specific pipeline, including
// auto-rerun information and detailed status. The caller must have read
// access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - pipelineID: The unique pipeline ID.
//
// Returns a slice of pipeline workflows or an error if the request fails or the API
// returns a non-success status.
func (c *Client) ListPipelineWorkflows(ctx context.Context, pipelineID string) ([]PipelineWorkflow, error) {
	path := fmt.Sprintf("/pipeline/%s/workflow", pipelineID)

	var resp listPipelineWorkflowResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return []PipelineWorkflow{}, fmt.Errorf("failed to get Pipeline{%s} workflows: %w", pipelineID, err)
	}
	return resp.Items, nil
}

// TriggerPipeline triggers a new pipeline for a project.
//
// The project slug is the three-segment slash-separated identifier: "{provider}/{organization}/{project}"
// (e.g., "gh/CircleCI-Public/circleci-cli").
//
// The caller must have write access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//   - body: The trigger configuration including branch/tag and parameters.
//
// Returns an error if the request fails or the API returns a non-success status.
func (c *Client) TriggerPipeline(ctx context.Context, projectSlug string, body TriggerPipelineBody) error {
	path := fmt.Sprintf("/project/%s/pipeline", projectSlug)

	var resp triggerPipelineResponse
	if err := c.do(ctx, http.MethodPost, path, body, &resp); err != nil {
		return fmt.Errorf("failed to trigger pipeline: %w", err)
	}
	return nil
}

// GetAllPipelines retrieves all pipelines for a project.
//
// The project slug is the three-segment slash-separated identifier: "{provider}/{organization}/{project}"
// (e.g., "gh/CircleCI-Public/circleci-cli").
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//   - opts: Optional filtering and pagination configuration.
//
// Returns a slice of pipelines or an error if the request fails or the API
// returns a non-success status.
func (c *Client) GetAllPipelines(ctx context.Context, projectSlug string, opts GetPipelineOptions) ([]Pipeline, error) {
	path := fmt.Sprintf("/project/%s/pipeline", projectSlug)
	path = appendQueryParams(path, opts)

	var resp pipelineListResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return []Pipeline{}, fmt.Errorf("failed to get pipelines for %s: %w", projectSlug, err)
	}
	return resp.Items, nil
}

// GetYourPipelines retrieves pipelines triggered by the authenticated user.
//
// The project slug is the three-segment slash-separated identifier: "{provider}/{organization}/{project}"
// (e.g., "gh/CircleCI-Public/circleci-cli").
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//   - opts: Optional pagination configuration.
//
// Returns a slice of pipelines or an error if the request fails or the API
// returns a non-success status.
func (c *Client) GetYourPipelines(ctx context.Context, projectSlug string, opts GetYourPipelineOptions) ([]Pipeline, error) {
	path := fmt.Sprintf("/project/%s/pipeline/mine", projectSlug)
	path = appendQueryParams(path, opts)

	var resp pipelineListResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return []Pipeline{}, fmt.Errorf("failed to get pipelines for %s: %w", projectSlug, err)
	}
	return resp.Items, nil
}

// GetPipelineByNumber retrieves a pipeline by its number within a project.
//
// The project slug is the three-segment slash-separated identifier: "{provider}/{organization}/{project}"
// (e.g., "gh/CircleCI-Public/circleci-cli").
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//   - pipelineNumber: The pipeline number.
//
// Returns the requested pipeline or an error if the request fails, the pipeline
// does not exist, or the API returns a non-success status.
func (c *Client) GetPipelineByNumber(ctx context.Context, projectSlug string, pipelineNumber int) (Pipeline, error) {
	path := fmt.Sprintf("/project/%s/pipeline/%d", projectSlug, pipelineNumber)

	var resp Pipeline
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return Pipeline{}, fmt.Errorf("failed to get Pipeline{%d}: %w", pipelineNumber, err)
	}
	return resp, nil
}

// TriggerPipelineProject triggers a new pipeline for a project using the project-based endpoint.
//
// Unlike TriggerPipeline which takes a full project slug, this method takes the provider,
// organization, and project as separate parameters.
//
// The caller must have write access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - provider: The VCS provider ("github"/"gh", "bitbucket"/"bb", or "circleci").
//   - org: The organization name or ID.
//   - project: The project name or ID.
//   - body: The trigger configuration including definition ID, config, checkout, and parameters.
//
// Returns an error if the request fails or the API returns a non-success status.
func (c *Client) TriggerPipelineProject(ctx context.Context, provider, org, project string, body TriggerProjPipelineBody) error {
	path := fmt.Sprintf("/project/%s/%s/%s/pipeline/run", provider, org, project)

	if err := c.do(ctx, http.MethodPost, path, body, nil); err != nil {
		return fmt.Errorf("failed to trigger project pipeline: %w", err)
	}
	return nil
}

// GetPipelineValues retrieves the values/parameters for a pipeline.
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - pipelineID: The unique pipeline ID.
//
// Returns an error if the request fails or the API returns a non-success status.
func (c *Client) GetPipelineValues(ctx context.Context, pipelineID string) error {
	path := fmt.Sprintf("/pipeline/%s/values", pipelineID)

	var resp map[string]any
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return fmt.Errorf("failed to get pipeline{%s} values: %w", pipelineID, err)
	}
	return nil
}
