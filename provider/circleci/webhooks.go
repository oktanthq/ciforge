package circleci

import (
	"context"
	"fmt"
	"net/http"
)

// ListWebhooks returns all webhooks matching the given filter options.
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - opts: Filtering options including scope ID and scope type.
//
// Returns a slice of webhooks or an error if the request fails or the API
// returns a non-success status.
func (c *Client) ListWebhooks(ctx context.Context, opts ListWebhooksOptions) ([]Webhook, error) {
	path := "/webhook"
	appendQueryParams(path, opts)

	var resp listWebhookResponse
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return []Webhook{}, fmt.Errorf("failed to list all webhooks: %w", err)
	}
	return resp.Items, nil
}

// CreateWebhooks creates a new webhook for a project.
//
// The caller must have write access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - body: The webhook configuration including name, events, URL, and scope.
//
// Returns the created webhook or an error if the request fails or the API
// returns a non-success status.
func (c *Client) CreateWebhooks(ctx context.Context, body CreateWebhooksBody) (Webhook, error) {
	path := "/webhook"

	var resp Webhook
	if err := c.do(ctx, http.MethodPost, path, body, &resp); err != nil {
		return Webhook{}, fmt.Errorf("failed to create webhook: %w", err)
	}
	return resp, nil
}

// GetWebhook retrieves a webhook by its unique ID.
//
// The caller must have read access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - webhookID: The unique webhook ID.
//
// Returns the requested webhook or an error if the request fails, the webhook
// does not exist, or the API returns a non-success status.
func (c *Client) GetWebhook(ctx context.Context, webhookID string) (Webhook, error) {
	path := fmt.Sprintf("/webhook/%s", webhookID)

	var resp Webhook
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return Webhook{}, fmt.Errorf("failed to get webhook{%s}: %w", webhookID, err)
	}
	return resp, nil
}

// UpdateWebhook updates an existing webhook by its unique ID.
//
// The caller must have write access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - webhookID: The unique webhook ID.
//   - body: The updated webhook configuration.
//
// Returns the updated webhook or an error if the request fails or the API
// returns a non-success status.
func (c *Client) UpdateWebhook(ctx context.Context, webhookID string, body UpdateWebhookBody) (Webhook, error) {
	path := fmt.Sprintf("/webhook/%s", webhookID)

	var resp Webhook
	if err := c.do(ctx, http.MethodPut, path, body, &resp); err != nil {
		return Webhook{}, fmt.Errorf("failed to patch webhook{%s}: %w", webhookID, err)
	}
	return resp, nil
}

// DeleteWebhooks deletes a webhook by its unique ID.
//
// The caller must have write access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - webhookID: The unique webhook ID to delete.
//
// Returns an error if the request fails or the API returns a non-success status.
func (c *Client) DeleteWebhooks(ctx context.Context, webhookID string) error {
	path := fmt.Sprintf("/webhook/%s", webhookID)

	if err := c.do(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("failed to delete webhook{%s}: %w", webhookID, err)
	}
	return nil
}
