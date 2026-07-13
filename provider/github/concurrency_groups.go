package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/abdulahshoaib/ciforge/models"
)

// ListConcurrencyGroupsForRepo retrieves all GitHub Actions concurrency groups
// defined for a repository.
//
// Concurrency groups are used to limit the number of workflow runs or jobs that
// can execute simultaneously for a given concurrency key.
//
// The caller must have read access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//
// Returns a slice of concurrency groups defined for the repository or an error
// if the request fails or the API returns a non-success status.
func (c *Client) ListConcurrencyGroupsForRepo(ctx context.Context, repo, owner string) ([]models.ConcurencyGroup, error) {
	path := fmt.Sprintf("repos/%s/%s/actions/concurrency_groups", owner, repo)
	// TODO(abdulahshoaib): add query params

	var resp listConcurrencyGroups
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return []models.ConcurencyGroup{}, fmt.Errorf("failed to list concurrency groups: %w", err)
	}
	return mapConcurencyGroups(resp.ConcurrencyGroups), nil
}

// GetConcurrencyGroupsForRepo retrieves information about a specific GitHub
// Actions concurrency group in a repository.
//
// Concurrency groups ensure that only a limited number of workflow runs or jobs
// sharing the same group name execute simultaneously.
//
// The caller must have read access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - concurrency_group_name: Name of the concurrency group to retrieve.
//
// Returns the requested concurrency group or an error if the request fails,
// the concurrency group does not exist, or the API returns a non-success
// status.
func (c *Client) GetConcurrencyGroupsForRepo(ctx context.Context, repo, owner, concurrency_group_name string) (models.ConcurencyGroup, error) {
	path := fmt.Sprintf("repo/%s/%s/actions/concurrency_groups/%s", owner, repo, concurrency_group_name)
	// TODO(abdulahshoaib): add query params

	var resp ConcurrencyGroup
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return models.ConcurencyGroup{}, fmt.Errorf("failed to get concurrency group{%s}: %w", concurrency_group_name, err)
	}
	return mapConcurencyGroup(resp), nil
}

// ListConcurrencyGroupsForWorkflow retrieves all concurrency groups associated
// with a specific GitHub Actions workflow run.
//
// Unlike ListConcurrencyGroupsForRepo, which returns all concurrency groups
// defined in a repository, this method only returns the concurrency groups used
// by the specified workflow run.
//
// The caller must have read access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - runID: Identifier of the workflow run.
//
// Returns a slice of concurrency groups associated with the workflow run or an
// error if the request fails or the API returns a non-success status.
func (c *Client) ListConcurrencyGroupsForWorflow(ctx context.Context, repo, owner, runID string) ([]models.ConcurencyGroup, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/runs/%s/concurrency_groups", owner, repo, runID)
	// TODO(abdulahshoaib): add query params

	var resp listConcurrencyGroups
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return []models.ConcurencyGroup{}, fmt.Errorf("failed to get concurency group for workflow run{%s}: %w", runID, err)
	}

	return mapConcurencyGroups(resp.ConcurrencyGroups), nil
}
