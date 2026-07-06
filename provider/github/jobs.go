package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/abdulahshoaib/ciforge/models"
)

func (c *Client) ListJobs(ctx context.Context, repo string, runID string) ([]models.Job, error) {
	path := fmt.Sprintf("/repos/%s/actions/runs/%s/jobs", repo, runID)
	var resp listJobsResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list jobs: %w", err)
	}
	return mapJobs(resp.Jobs), nil
}
