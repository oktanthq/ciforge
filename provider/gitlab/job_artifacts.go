package gitlab

import (
	"context"
	"fmt"
	"net/http"
)

// DownloadJobArtifacts downloads a job's artifacts archive by job ID.
//
// If you use cURL to download artifacts from GitLab.com, use the --location
// parameter as the request might redirect through a CDN.
//
// The caller must have at least reporter access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - jobID: The ID of the job.
//   - jobToken: Optional CI/CD job token for multi-project pipelines.
//
// Returns the raw artifacts bytes or an error if the request fails or the
// API returns a non-success status.
func (c *Client) DownloadJobArtifacts(ctx context.Context, projectID, jobID int, jobToken string) ([]byte, error) {
	path := fmt.Sprintf("/projects/%d/jobs/%d/artifacts", projectID, jobID)
	if jobToken != "" {
		path += "?job_token=" + jobToken
	}

	data, err := c.doRaw(ctx, http.MethodGet, path)
	if err != nil {
		return nil, fmt.Errorf("failed to download job artifacts: %w", err)
	}
	return data, nil
}

// DownloadJobArtifactsByRef downloads a job's artifacts archive from the latest
// successful pipeline using a reference name.
//
// For parent and child pipelines, artifacts are searched in hierarchical order
// from parent to child. If both parent and child pipelines have a job with the
// same name, the artifact from the parent pipeline is returned.
//
// Prerequisites:
//   - You must have a completed pipeline with a success status.
//   - If the pipeline includes manual jobs, they must either complete
//     successfully or have allow_failure: true set.
//
// The caller must have at least reporter access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - refName: Branch or tag name in repository.
//   - job: The name of the job.
//   - opts: Optional configuration for the download.
//
// Returns the raw artifacts bytes or an error if the request fails or the
// API returns a non-success status.
func (c *Client) DownloadJobArtifactsByRef(ctx context.Context, projectID int, refName, job string, opts *DownloadJobArtifactsByRefOptions) ([]byte, error) {
	path := fmt.Sprintf("/projects/%d/jobs/artifacts/%s/download?job=%s", projectID, refName, job)
	path = appendQueryParams(path, opts)

	data, err := c.doRaw(ctx, http.MethodGet, path)
	if err != nil {
		return nil, fmt.Errorf("failed to download job artifacts by ref: %w", err)
	}
	return data, nil
}

// DownloadSingleArtifactFile downloads a single file from a job's artifacts
// by job ID. The file is extracted from the archive and streamed to the client.
//
// The caller must have at least reporter access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - jobID: The ID of the job.
//   - artifactPath: Path to a file inside the artifacts archive.
//   - jobToken: Optional CI/CD job token for multi-project pipelines.
//
// Returns the raw file bytes or an error if the request fails or the API
// returns a non-success status.
func (c *Client) DownloadSingleArtifactFile(ctx context.Context, projectID, jobID int, artifactPath, jobToken string) ([]byte, error) {
	path := fmt.Sprintf("/projects/%d/jobs/%d/artifacts/%s", projectID, jobID, artifactPath)
	if jobToken != "" {
		path += "?job_token=" + jobToken
	}

	data, err := c.doRaw(ctx, http.MethodGet, path)
	if err != nil {
		return nil, fmt.Errorf("failed to download single artifact file: %w", err)
	}
	return data, nil
}

// ListArtifactTree lists all files and directories in the artifacts archive
// of a specified job.
//
// This operation reads the artifact metadata without extracting the full
// archive, making it efficient for browsing large archives.
//
// The caller must have at least reporter access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - jobID: The ID of the job.
//   - opts: Optional filtering parameters.
//
// Returns a slice of artifact file entries or an error if the request fails
// or the API returns a non-success status.
func (c *Client) ListArtifactTree(ctx context.Context, projectID, jobID int, opts *ListArtifactTreeOptions) ([]ArtifactFileEntry, error) {
	path := fmt.Sprintf("/projects/%d/jobs/%d/artifacts/tree", projectID, jobID)
	path = appendQueryParams(path, opts)

	var resp []ArtifactFileEntry
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list artifact tree: %w", err)
	}
	return resp, nil
}

// DownloadSingleArtifactFileByRef downloads a single file from a job's
// artifacts in the latest successful pipeline using the reference name.
//
// The file is extracted from the archive and streamed to the client.
//
// For parent and child pipelines, artifacts are searched in hierarchical order
// from parent to child. If both parent and child pipelines have a job with the
// same name, the artifact from the parent pipeline is returned.
//
// Prerequisites:
//   - You must have a completed pipeline with a success status.
//   - If the pipeline includes manual jobs, they must either complete
//     successfully or have allow_failure: true set.
//
// The caller must have at least reporter access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - refName: Branch or tag name in repository.
//   - artifactPath: Path to a file inside the artifacts archive.
//   - job: The name of the job.
//   - opts: Optional configuration for the download.
//
// Returns the raw file bytes or an error if the request fails or the API
// returns a non-success status.
func (c *Client) DownloadSingleArtifactFileByRef(ctx context.Context, projectID int, refName, artifactPath, job string, opts *DownloadSingleArtifactFileByRefOptions) ([]byte, error) {
	path := fmt.Sprintf("/projects/%d/jobs/artifacts/%s/raw/%s?job=%s", projectID, refName, artifactPath, job)
	path = appendQueryParams(path, opts)

	data, err := c.doRaw(ctx, http.MethodGet, path)
	if err != nil {
		return nil, fmt.Errorf("failed to download single artifact file by ref: %w", err)
	}
	return data, nil
}

// KeepJobArtifacts prevents a job's artifacts from being automatically deleted
// when they reach their expiration date.
//
// The caller must have at least developer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - jobID: The ID of the job.
//
// Returns the updated job or an error if the request fails or the API
// returns a non-success status.
func (c *Client) KeepJobArtifacts(ctx context.Context, projectID, jobID int) (*Job, error) {
	path := fmt.Sprintf("/projects/%d/jobs/%d/artifacts/keep", projectID, jobID)

	var resp Job
	if err := c.do(ctx, http.MethodPost, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to keep job artifacts: %w", err)
	}
	return &resp, nil
}

// DeleteJobArtifacts deletes all artifacts associated with a specific job.
//
// Artifacts cannot be recovered after they are deleted.
//
// Prerequisites:
//   - You must have the Maintainer or Owner role for the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - jobID: The ID of the job.
//
// Returns an error if the deletion fails or the API returns a non-success
// status.
func (c *Client) DeleteJobArtifacts(ctx context.Context, projectID, jobID int) error {
	path := fmt.Sprintf("/projects/%d/jobs/%d/artifacts", projectID, jobID)

	if err := c.do(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("failed to delete job artifacts: %w", err)
	}
	return nil
}

// DeleteAllJobArtifacts deletes all job artifacts eligible for deletion in a
// project.
//
// Artifacts cannot be recovered after they are deleted. By default, artifacts
// from the most recent successful pipeline of each ref are not deleted.
//
// Requests to this endpoint set the expiry of all job artifacts that can be
// deleted to the current time. The files are then deleted from the system as
// part of the regular cleanup of expired job artifacts. Job logs are never
// deleted.
//
// The regular cleanup occurs asynchronously on a schedule, so there might be a
// short delay before artifacts are deleted.
//
// Prerequisites:
//   - You must have the Maintainer or Owner role for the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//
// Returns an error if the request fails or the API returns a non-success
// status.
func (c *Client) DeleteAllJobArtifacts(ctx context.Context, projectID int) error {
	path := fmt.Sprintf("/projects/%d/artifacts", projectID)

	if err := c.do(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("failed to delete all job artifacts: %w", err)
	}
	return nil
}
