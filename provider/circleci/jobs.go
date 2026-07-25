package circleci

import (
	"context"
	"fmt"
	"net/http"
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
func (c *Client) GetJob(ctx context.Context, jobID string) (Job, error) {
	path := fmt.Sprintf("/jobs/%s", jobID)

	var resp Job
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return Job{}, fmt.Errorf("failed to get job: %w", err)
	}
	return resp, nil
}

// CancelJobByJobNumber cancels a running job.
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
func (c *Client) CancelJobByJobNumber(ctx context.Context, projectSlug string, jobNumber int64) error {
	path := fmt.Sprintf("/project/%s/job/%d/cancel", projectSlug, jobNumber)

	if err := c.do(ctx, http.MethodPost, path, nil, nil); err != nil {
		return fmt.Errorf("failed to cancel job{%d}: %w", jobNumber, err)
	}
	return nil
}

// CancelJob cancels a running job by its ID.
//
// The caller must have write access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - jobID: The unique job ID.
//
// Returns an error if the request fails or the API returns a non-success status.
func (c *Client) CancelJob(ctx context.Context, jobID string) error {
	path := fmt.Sprintf("/jobs/%s/cancel", jobID)

	if err := c.do(ctx, http.MethodPost, path, nil, nil); err != nil {
		return fmt.Errorf("failed to cancel job{%s}: %w", jobID, err)
	}
	return nil
}
