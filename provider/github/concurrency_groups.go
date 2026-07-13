package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/abdulahshoaib/ciforge/models"
)

func (c *Client) ListConcurrencyGroupsForRepo(ctx context.Context, repo, owner string) ([]models.ConcurencyGroup, error) {
	path := fmt.Sprintf("repos/%s/%s/actions/concurrency_groups", owner, repo)

	var resp listConcurrencyGroups
	if err := c.do(ctx, http.MethodGet, path, nil, resp); err != nil {
		return []models.ConcurencyGroup{}, fmt.Errorf("failed to list concurrency groups: %w", err)
	}
	return mapConcurencyGroups(resp.ConcurrencyGroups), nil
}
