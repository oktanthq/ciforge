package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/abdulahshoaib/ciforge/models"
)

func (c *Client) ListArtifacts(ctx context.Context, repo string, runID string) ([]models.Artifact, error) {
	path := fmt.Sprintf("/repos/%s/actions/runs/%s/artifacts", repo, runID)
	var resp listArtifactsResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list artifacts: %w", err)
	}
	return mapArtifacts(resp.Artifacts, runID), nil
}

func (c *Client) DownloadArtifact(ctx context.Context, repo string, artifactID string) ([]byte, error) {
	path := fmt.Sprintf("/repos/%s/actions/artifacts/%s/zip", repo, artifactID)
	data, err := c.doRaw(ctx, http.MethodGet, path)
	if err != nil {
		return nil, fmt.Errorf("failed to download artifact: %w", err)
	}
	return data, nil
}
