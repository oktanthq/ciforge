package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/abdulahshoaib/ciforge/models"
)

// ListArtifacts retrieves all GitHub Actions artifacts for a repository.
//
// The caller must have read access to the repository. The returned artifacts
// include metadata such as their name, size, creation time, expiration time,
// and download URL.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//
// Returns a slice of artifacts or an error if the request fails or the API
// returns a non-success status.
func (c *Client) ListArtifacts(ctx context.Context, repo, owner string) ([]models.Artifact, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/artifacts", owner, repo)

	var resp listArtifactsResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list artifacts: %w", err)
	}

	return mapArtifacts(resp.Artifacts), nil
}

// GetArtifact retrieves a GitHub Actions artifact by its ID.
//
// The caller must have read access to the repository. The returned artifact
// includes metadata such as its name, size, creation time, expiration time,
// and download URL.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - artifactID: Identifier of the artifact to retrieve.
//
// Returns the requested artifact or an error if the request fails, the
// artifact does not exist, or the API returns a non-success status.
func (c *Client) GetArtifact(ctx context.Context, repo, owner, artifactID string) (models.Artifact, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/artifacts/%s", owner, repo, artifactID)

	var resp artifactDTO
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return models.Artifact{}, fmt.Errorf("failed to get artifact: %w", err)
	}

	return mapArtifact(resp), nil
}

// DownloadArtifact downloads a GitHub Actions artifact from its download URL.
//
// The artifact is returned as raw bytes. The caller is responsible for
// writing the data to a file or extracting the archive as needed.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - downloadURL: Fully qualified URL used to download the artifact.
//
// Returns the downloaded artifact as a byte slice or an error if the request
// fails or the download cannot be completed.
//
// The download URL is provided by `ListArtifacts`. GitHub artifacts are
// downloaded as ZIP archives.
func (c *Client) DownloadArtifact(ctx context.Context, downloadURL string) ([]byte, error) {
	data, err := c.doRawURL(ctx, http.MethodGet, downloadURL)
	if err != nil {
		return nil, fmt.Errorf("failed to download artifact: %w", err)
	}
	return data, nil
}

func (c *Client) DeleteArtifact(ctx context.Context, repo, owner, artifactID string) error {
	path := fmt.Sprintf("repos/%s/%s/actions/artifacts/%s", owner, repo, artifactID)

	if err := c.do(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("failed to delete artifact: %w", err)
	}

	return nil
}
