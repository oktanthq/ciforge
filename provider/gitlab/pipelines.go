package gitlab

import (
	"context"
	"fmt"
	"net/http"
)

// ListPipelines retrieves all CI/CD pipelines in a project.
//
// By default, child pipelines are not included in the results. To return
// child pipelines, set Source to PipelineSourceParentPipeline in the options.
//
// The caller must have at least reporter access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - opts: Optional list configuration for filtering and pagination.
//
// Returns a slice of pipelines or an error if the request fails or the API
// returns a non-success status.
func (c *Client) ListPipelines(ctx context.Context, projectID int, opts *ListPipelinesOptions) ([]Pipeline, error) {
	path := fmt.Sprintf("/projects/%d/pipelines", projectID)
	path = appendQueryParams(path, opts)

	var resp []Pipeline
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list pipelines: %w", err)
	}
	return resp, nil
}

// GetPipeline retrieves a single pipeline from a project.
//
// You can also get a single child pipeline using this method.
//
// The caller must have at least reporter access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - pipelineID: The ID of the pipeline.
//
// Returns the requested pipeline or an error if the request fails, the pipeline
// does not exist, or the API returns a non-success status.
func (c *Client) GetPipeline(ctx context.Context, projectID, pipelineID int) (PipelineDetailed, error) {
	path := fmt.Sprintf("/projects/%d/pipelines/%d", projectID, pipelineID)

	var resp PipelineDetailed
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return PipelineDetailed{}, fmt.Errorf("failed to get pipeline: %w", err)
	}
	return resp, nil
}

// GetLatestPipeline retrieves the latest pipeline for the most recent commit
// on a specific ref in a project.
//
// If no pipeline exists for the commit, a 403 status code is returned as an
// error.
//
// The caller must have at least reporter access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - opts: Optional query configuration for filtering by ref.
//
// Returns the latest pipeline or an error if the request fails, no pipeline
// exists for the ref, or the API returns a non-success status.
func (c *Client) GetLatestPipeline(ctx context.Context, projectID int, opts *GetLatestPipelineOptions) (PipelineDetailed, error) {
	path := fmt.Sprintf("/projects/%d/pipelines/latest", projectID)
	path = appendQueryParams(path, opts)

	var resp PipelineDetailed
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return PipelineDetailed{}, fmt.Errorf("failed to get latest pipeline: %w", err)
	}
	return resp, nil
}

// GetPipelineVariables retrieves the pipeline variables of a pipeline.
//
// Pipeline variables are key-value pairs that are available to jobs during
// pipeline execution.
//
// The caller must have at least reporter access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - pipelineID: The ID of the pipeline.
//
// Returns a slice of pipeline variables or an error if the request fails or
// the API returns a non-success status.
func (c *Client) GetPipelineVariables(ctx context.Context, projectID, pipelineID int) ([]PipelineVariable, error) {
	path := fmt.Sprintf("/projects/%d/pipelines/%d/variables", projectID, pipelineID)

	var resp []PipelineVariable
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to get pipeline variables: %w", err)
	}
	return resp, nil
}

// GetPipelineTestReport retrieves the test report for a pipeline.
//
// The test report contains information about test suites, test cases, and
// their execution results. This endpoint is part of the Unit test report
// feature.
//
// The caller must have at least reporter access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - pipelineID: The ID of the pipeline.
//
// Returns the pipeline test report or an error if the request fails or the
// API returns a non-success status.
func (c *Client) GetPipelineTestReport(ctx context.Context, projectID, pipelineID int) (PipelineTestReport, error) {
	path := fmt.Sprintf("/projects/%d/pipelines/%d/test_report", projectID, pipelineID)

	var resp PipelineTestReport
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return PipelineTestReport{}, fmt.Errorf("failed to get pipeline test report: %w", err)
	}
	return resp, nil
}

// GetPipelineTestReportSummary retrieves the test report summary for a pipeline.
//
// The summary contains aggregated counts and timing information for test suites
// and their build IDs. This endpoint is part of the Unit test report feature.
//
// The caller must have at least reporter access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - pipelineID: The ID of the pipeline.
//
// Returns the pipeline test report summary or an error if the request fails or
// the API returns a non-success status.
func (c *Client) GetPipelineTestReportSummary(ctx context.Context, projectID, pipelineID int) (PipelineTestReportSummary, error) {
	path := fmt.Sprintf("/projects/%d/pipelines/%d/test_report_summary", projectID, pipelineID)

	var resp PipelineTestReportSummary
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return PipelineTestReportSummary{}, fmt.Errorf("failed to get pipeline test report summary: %w", err)
	}
	return resp, nil
}

// CreatePipeline creates a new CI/CD pipeline in a project.
//
// The ref parameter specifies the branch or tag to run the pipeline on.
// For merge request pipelines, use the merge requests endpoint instead.
//
// The caller must have at least developer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - ref: The branch or tag to run the pipeline on.
//
// Returns the newly created pipeline or an error if the request fails or the
// API returns a non-success status.
func (c *Client) CreatePipeline(ctx context.Context, projectID int, ref string) (Pipeline, error) {
	path := fmt.Sprintf("/projects/%d/pipeline?ref=%s", projectID, ref)

	var resp Pipeline
	if err := c.do(ctx, http.MethodPost, path, nil, &resp); err != nil {
		return Pipeline{}, fmt.Errorf("failed to create pipeline: %w", err)
	}
	return resp, nil
}

// RetryPipeline retries failed or canceled jobs in a pipeline.
//
// If there are no failed or canceled jobs in the pipeline, calling this
// endpoint has no effect.
//
// The caller must have at least developer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - pipelineID: The ID of the pipeline.
//
// Returns the updated pipeline or an error if the request fails or the API
// returns a non-success status.
func (c *Client) RetryPipeline(ctx context.Context, projectID, pipelineID int) (Pipeline, error) {
	path := fmt.Sprintf("/projects/%d/pipelines/%d/retry", projectID, pipelineID)

	var resp Pipeline
	if err := c.do(ctx, http.MethodPost, path, nil, &resp); err != nil {
		return Pipeline{}, fmt.Errorf("failed to retry pipeline: %w", err)
	}
	return resp, nil
}

// CancelPipeline cancels all jobs for a pipeline.
//
// This endpoint returns a success response regardless of the pipeline's state.
//
// The caller must have at least developer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - pipelineID: The ID of the pipeline.
//
// Returns the updated pipeline or an error if the request fails or the API
// returns a non-success status.
func (c *Client) CancelPipeline(ctx context.Context, projectID, pipelineID int) (Pipeline, error) {
	path := fmt.Sprintf("/projects/%d/pipelines/%d/cancel", projectID, pipelineID)

	var resp Pipeline
	if err := c.do(ctx, http.MethodPost, path, nil, &resp); err != nil {
		return Pipeline{}, fmt.Errorf("failed to cancel pipeline: %w", err)
	}
	return resp, nil
}

// DeletePipeline permanently deletes a pipeline.
//
// Deleting a pipeline expires all pipeline caches and deletes all immediately
// related objects, such as builds, logs, artifacts, and triggers. This action
// cannot be undone.
//
// Deleting a pipeline does not automatically delete its child pipelines.
//
// The caller must have at least maintainer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - pipelineID: The ID of the pipeline.
//
// Returns an error if the deletion fails or the API returns a non-success
// status.
func (c *Client) DeletePipeline(ctx context.Context, projectID, pipelineID int) error {
	path := fmt.Sprintf("/projects/%d/pipelines/%d", projectID, pipelineID)

	if err := c.do(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("failed to delete pipeline: %w", err)
	}
	return nil
}

// UpdatePipelineMetadata updates the metadata of a pipeline.
//
// The metadata currently contains only the name of the pipeline.
//
// The caller must have at least developer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - pipelineID: The ID of the pipeline.
//   - name: The new name of the pipeline.
//
// Returns the updated pipeline metadata or an error if the request fails or
// the API returns a non-success status.
func (c *Client) UpdatePipelineMetadata(ctx context.Context, projectID, pipelineID int, name string) (Pipeline, error) {
	path := fmt.Sprintf("/projects/%d/pipelines/%d/metadata", projectID, pipelineID)

	body := PipelineMetadata{Name: name}

	var resp Pipeline
	if err := c.do(ctx, http.MethodPut, path, body, &resp); err != nil {
		return Pipeline{}, fmt.Errorf("failed to update pipeline metadata: %w", err)
	}
	return resp, nil
}
