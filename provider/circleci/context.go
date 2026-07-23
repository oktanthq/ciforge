package circleci

import (
	"context"
	"fmt"
	"net/http"
)

// CreateContext creates a new context for storing environment variables.
//
// Contexts are used to share environment variables and configuration across
// multiple projects within an organization.
//
// The caller must have write access to the organization.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - body: The context configuration including name and owner.
//
// Returns the created context or an error if the request fails or the API
// returns a non-success status.
func (c *Client) CreateContext(ctx context.Context, body ContextBody) (Context, error) {
	path := "/context"
	var resp Context
	if err := c.do(ctx, http.MethodPost, path, body, &resp); err != nil {
		return Context{}, fmt.Errorf("failed to create context: %w", err)
	}
	return resp, nil
}
