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

func (c *Client) GetLogs(ctx context.Context, repo string, runID string) (*models.Log, error) {
	path := fmt.Sprintf("/repos/%s/actions/runs/%s/logs", repo, runID)
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
