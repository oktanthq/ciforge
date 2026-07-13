package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/abdulahshoaib/ciforge/models"
)

func (c *Client) ListConcurrencyGroupsForRepo(ctx context.Context, repo, owner string) ([]models.ConcurencyGroup, error) {
	path := fmt.Sprintf("repos/%s/%s/actions/concurrency_groups", owner, repo)
	// TODO(abdulahshoaib): add query params

	var resp listConcurrencyGroups
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return []models.ConcurencyGroup{}, fmt.Errorf("failed to list concurrency groups: %w", err)
	}
	return mapConcurencyGroups(resp.ConcurrencyGroups), nil
}

func (c *Client) GetConcurrencyGroupsForRepo(ctx context.Context, repo, owner, concurrency_group_name string) (models.ConcurencyGroup, error) {
	path := fmt.Sprintf("repo/%s/%s/actions/concurrency_groups/%s", owner, repo, concurrency_group_name)
	// TODO(abdulahshoaib): add query params

	var resp ConcurrencyGroup
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return models.ConcurencyGroup{}, fmt.Errorf("failed to get concurrency group{%s}: %w", concurrency_group_name, err)
	}
	return mapConcurencyGroup(resp), nil
}

func (c *Client) ListConcurrencyGroupsForWorflow(ctx context.Context, repo, owner, runID string) ([]models.ConcurencyGroup, error) {
	path := fmt.Sprintf("/repo/%s/%s/actions/runs/%s/concurrency_groups", owner, repo, runID)
	// TODO(abdulahshoaib): add query params

	var resp listConcurrencyGroups
	if err := c.do(ctx, http.MethodGet, path, nil, resp); err != nil {
		return []models.ConcurencyGroup{}, fmt.Errorf("failed to get concurency group for workflow run{%s}: %w", runID, err)
	}

	return mapConcurencyGroups(resp.ConcurrencyGroups), nil
}
