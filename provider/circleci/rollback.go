package circleci

import (
	"context"
	"fmt"
	"net/http"
)

// RollbackProject rolls back a component deployment to a previous version.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The unique identifier of the project.
//   - body: The rollback configuration including component, versions, and reason.
//
// Returns the rollback response or an error if the request fails.
func (c *Client) RollbackProject(ctx context.Context, projectID string, body RollbackBody) (RollbackResponse, error) {
	path := fmt.Sprintf("/projects/%s/rollback", projectID)

	var resp RollbackResponse
	if err := c.do(ctx, http.MethodPost, path, body, &resp); err != nil {
		return RollbackResponse{}, fmt.Errorf("failed to rollback project{%s}: %w", projectID, err)
	}
	return resp, nil
}
