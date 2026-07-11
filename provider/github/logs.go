package github

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/abdulahshoaib/ciforge/models"
)

// GetWorkflowRunLogs downloads the logs for a workflow run and extracts the
// contents from the ZIP archive.
//
// Unlike DownloadWorkflowRunLogs, which triggers a raw download, this method
// extracts all log files from the archive and returns them as a single string.
//
// The caller must have read access to the repository.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - owner: Repository owner or organization.
//   - repo: Repository name.
//   - runID: Identifier of the workflow run.
//
// Returns the extracted log contents or an error if the request fails, the
// archive cannot be read, or the API returns a non-success status.
func (c *Client) GetWorkflowRunLogs(ctx context.Context, repo, owner, runID string) (*models.Log, error) {
	path := fmt.Sprintf("/repos/%s/%s/actions/runs/%s/logs", owner, repo, runID)
	data, err := c.doRaw(ctx, http.MethodGet, path)
	if err != nil {
		return nil, fmt.Errorf("failed to download logs: %w", err)
	}

	reader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, fmt.Errorf("failed to read log archive: %w", err)
	}

	var content strings.Builder
	for _, f := range reader.File {
		if f.FileInfo().IsDir() {
			continue
		}
		rc, err := f.Open()
		if err != nil {
			continue
		}
		if _, err := io.Copy(&content, rc); err != nil {
			rc.Close()
			continue
		}
		rc.Close()
		content.WriteString("\n")
	}

	return &models.Log{
		RunID:   runID,
		Content: content.String(),
	}, nil
}
