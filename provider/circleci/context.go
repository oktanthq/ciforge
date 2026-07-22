package circleci

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) CreateContext(ctx context.Context, body ContextBody) (Context, error) {
	path := "/context"
	var resp Context
	if err := c.do(ctx, http.MethodPost, path, body, &resp); err != nil {
		return Context{}, fmt.Errorf("failed to create context: %w", err)
	}
	return resp, nil
}
