package circleci

import (
	"context"
	"fmt"
)

// WorkflowMetrics represents metrics for a workflow.
type WorkflowMetrics struct {
	// Name is the workflow name.
	Name string `json:"name"`

	// Metrics contains the workflow metrics.
	Metrics WorkflowMetricsData `json:"metrics"`
}

// WorkflowMetricsData contains the actual metrics data.
type WorkflowMetricsData struct {
	// TotalRuns is the total number of workflow runs.
	TotalRuns int64 `json:"total_runs"`

	// SuccessfulRuns is the number of successful runs.
	SuccessfulRuns MetricData `json:"successful_runs"`

	// FailedRuns is the number of failed runs.
	FailedRuns MetricData `json:"failed_runs"`

	// TotalDuration is the total duration across all runs.
	TotalDuration MetricData `json:"total_duration"`

	// MedianDuration is the median duration of runs.
	MedianDuration MetricData `json:"median_duration"`

	// SuccessRate is the success rate percentage.
	SuccessRate float64 `json:"success_rate"`

	// Throughput is the average number of runs per day.
	Throughput float64 `json:"throughput"`

	// TotalCreditsUsed is the total credits used.
	TotalCreditsUsed int64 `json:"total_credits_used"`
}

// MetricData represents a metric with min, max, mean, median, and percentile values.
type MetricData struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Mean float64 `json:"mean"`
	Median float64 `json:"median"`
	P95 float64 `json:"p95"`
	StandardDeviation float64 `json:"standard_deviation"`
}

// WorkflowRunInsight represents a single workflow run in insights.
type WorkflowRunInsight struct {
	// ID is the run ID.
	ID string `json:"id"`

	// Status is the run status.
	Status string `json:"status"`

	// CreditsUsed is the number of credits used.
	CreditsUsed int64 `json:"credits_used"`

	// Duration is the run duration in seconds.
	Duration float64 `json:"duration"`

	// StartedAt is when the run started.
	StartedAt string `json:"started_at"`

	// StoppedAt is when the run stopped.
	StoppedAt string `json:"stopped_at"`
}

// workflowMetricsListResponse is the response from listing workflow metrics.
type workflowMetricsListResponse struct {
	Items []WorkflowMetrics `json:"items"`
}

// workflowRunsListResponse is the response from listing workflow runs.
type workflowRunsListResponse struct {
	Items         []WorkflowRunInsight `json:"items"`
	NextPageToken string               `json:"next_page_token"`
}

// ListWorkflowMetrics retrieves summary metrics for a project's workflows.
//
// The project slug must be in the form of "vcs/org/repo" (e.g., "gh/CircleCI-Public/circleci-cli").
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//   - opts: Optional filtering parameters.
//
// Returns a slice of workflow metrics or an error if the request fails or the API
// returns a non-success status.
func (c *Client) ListWorkflowMetrics(ctx context.Context, projectSlug string, opts *ListWorkflowMetricsOptions) ([]WorkflowMetrics, error) {
	path := fmt.Sprintf("/insights/%s/workflows", projectSlug)
	path = appendQueryParams(path, opts)

	var resp workflowMetricsListResponse
	if err := c.do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list workflow metrics: %w", err)
	}
	return resp.Items, nil
}

// ListWorkflowRuns retrieves recent runs of a workflow.
//
// The project slug must be in the form of "vcs/org/repo" (e.g., "gh/CircleCI-Public/circleci-cli").
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//   - workflowName: The name of the workflow.
//   - opts: Optional filtering parameters.
//
// Returns a slice of workflow run insights or an error if the request fails or the API
// returns a non-success status.
func (c *Client) ListWorkflowRuns(ctx context.Context, projectSlug, workflowName string, opts *ListWorkflowRunsOptions) ([]WorkflowRunInsight, error) {
	path := fmt.Sprintf("/insights/%s/workflows/%s", projectSlug, workflowName)
	path = appendQueryParams(path, opts)

	var resp workflowRunsListResponse
	if err := c.do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list workflow runs: %w", err)
	}
	return resp.Items, nil
}

// ListJobMetrics retrieves summary metrics for jobs in a workflow.
//
// The project slug must be in the form of "vcs/org/repo" (e.g., "gh/CircleCI-Public/circleci-cli").
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//   - workflowName: The name of the workflow.
//   - opts: Optional filtering parameters.
//
// Returns a slice of job metrics or an error if the request fails or the API
// returns a non-success status.
func (c *Client) ListJobMetrics(ctx context.Context, projectSlug, workflowName string, opts *ListJobMetricsOptions) ([]WorkflowMetrics, error) {
	path := fmt.Sprintf("/insights/%s/workflows/%s/jobs", projectSlug, workflowName)
	path = appendQueryParams(path, opts)

	var resp workflowMetricsListResponse
	if err := c.do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list job metrics: %w", err)
	}
	return resp.Items, nil
}
