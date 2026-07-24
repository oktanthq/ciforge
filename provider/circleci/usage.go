package circleci

import (
	"context"
	"fmt"
	"net/http"
)

// CreateUsageExport creates a new usage export job for an organization.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - orgID: The unique identifier of the organization.
//   - body: The export configuration including date range and shared orgs.
//
// Returns the usage export job or an error if the request fails.
func (c *Client) CreateUsageExport(ctx context.Context, orgID string, body UsageExportBody) (UsageExport, error) {
	path := fmt.Sprintf("organizations/%s/usage_export_job", orgID)

	var resp UsageExport
	if err := c.do(ctx, http.MethodPost, path, body, &resp); err != nil {
		return UsageExport{}, fmt.Errorf("failed to create usage export: %w", err)
	}
	return resp, nil
}

// GetUsageExport returns the status and download URLs for a usage export job.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - orgID: The unique identifier of the organization.
//   - usageExpJobID: The unique identifier of the usage export job.
//
// Returns the usage export response or an error if the request fails.
func (c *Client) GetUsageExport(ctx context.Context, orgID, usageExpJobID string) (usageExportResponse, error) {
	path := fmt.Sprintf("organizations/%s/usage_export_job/%s", orgID, usageExpJobID)

	var resp usageExportResponse
	if err := c.do(ctx, http.MethodPost, path, nil, &resp); err != nil {
		return usageExportResponse{}, fmt.Errorf("failed to create usage export: %w", err)
	}
	return resp, nil
}
