package circleci

import (
	"context"
	"fmt"
)

// GetJob retrieves a job by its number within a project.
//
// The project slug is the three-segment slash-separated identifier: "{provider}/{organization}/{project}"
// (e.g., "gh/CircleCI-Public/circleci-cli").
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//   - jobNumber: The job number.
//
// Returns the requested job or an error if the request fails, the job
// does not exist, or the API returns a non-success status.
func (c *Client) GetJob(ctx context.Context, projectSlug string, jobNumber int64) (Job, error) {
	path := fmt.Sprintf("/project/%s/job/%d", projectSlug, jobNumber)

	var resp Job
	if err := c.do(ctx, "GET", path, nil, &resp); err != nil {
		return Job{}, fmt.Errorf("failed to get job: %w", err)
	}
	return resp, nil
}

// CancelJob cancels a running job.
//
// The project slug is the three-segment slash-separated identifier: "{provider}/{organization}/{project}"
// (e.g., "gh/CircleCI-Public/circleci-cli").
//
// The caller must have write access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//   - jobNumber: The job number.
//
// Returns an error if the request fails or the API returns a non-success status.
func (c *Client) CancelJob(ctx context.Context, projectSlug string, jobNumber int64) error {
	path := fmt.Sprintf("/project/%s/job/%d/cancel", projectSlug, jobNumber)

	if err := c.do(ctx, "POST", path, nil, nil); err != nil {
		return fmt.Errorf("failed to cancel job: %w", err)
	}
	return nil
}

// GetJobArtifacts retrieves the artifacts for a job.
//
// The project slug is the three-segment slash-separated identifier: "{provider}/{organization}/{project}"
// (e.g., "gh/CircleCI-Public/circleci-cli").
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//   - jobNumber: The job number.
//   - opts: Optional pagination configuration.
//
// Returns a slice of artifacts or an error if the request fails or the API
// returns a non-success status.
func (c *Client) GetJobArtifacts(ctx context.Context, projectSlug string, jobNumber int64, opts *GetJobArtifactsOptions) ([]Artifact, error) {
	path := fmt.Sprintf("/project/%s/%d/artifacts", projectSlug, jobNumber)
	path = appendQueryParams(path, opts)

	var resp artifactListResponse
	if err := c.do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to get job artifacts: %w", err)
	}
	return resp.Items, nil
}

// GetJobTests retrieves the test metadata for a job.
//
// The project slug is the three-segment slash-separated identifier: "{provider}/{organization}/{project}"
// (e.g., "gh/CircleCI-Public/circleci-cli").
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectSlug: The project slug (e.g., "gh/CircleCI-Public/circleci-cli").
//   - jobNumber: The job number.
//   - opts: Optional pagination configuration.
//
// Returns a slice of test results or an error if the request fails or the API
// returns a non-success status.
func (c *Client) GetJobTests(ctx context.Context, projectSlug string, jobNumber int64, opts *GetJobTestsOptions) ([]TestResult, error) {
	path := fmt.Sprintf("/project/%s/%d/tests", projectSlug, jobNumber)
	path = appendQueryParams(path, opts)

	var resp testListResponse
	if err := c.do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to get job tests: %w", err)
	}
	return resp.Items, nil
}
