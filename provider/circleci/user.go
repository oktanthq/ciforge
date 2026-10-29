package circleci

import (
	"context"
	"fmt"
	"net/http"
)

// GetMyInfo retrieves the authenticated user's information.
//
// The caller must be authenticated.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//
// Returns the user information or an error if the request fails or the API
// returns a non-success status.
func (c *Client) GetMyInfo(ctx context.Context) (User, error) {
	path := "/me"

	var resp User
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return User{}, fmt.Errorf("failed to get user info: %w", err)
	}
	return resp, nil
}

// GetCollaborations retrieves the organizations the authenticated user belongs to.
//
// The caller must be authenticated.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//
// Returns a slice of collaborations or an error if the request fails or the API
// returns a non-success status.
func (c *Client) GetCollaborations(ctx context.Context) ([]Collaboration, error) {
	path := "/me/collaborations"

	var resp []Collaboration
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return []Collaboration{}, fmt.Errorf("failed to get collaborations: %w", err)
	}
	return resp, nil
}

// GetUserInfo retrieves a user's information by their ID.
//
// The caller must be authenticated.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - id: The unique user ID.
//
// Returns the user information or an error if the request fails, the user
// does not exist, or the API returns a non-success status.
func (c *Client) GetUserInfo(ctx context.Context, id string) (User, error) {
	path := fmt.Sprintf("/user/%s", id)

	var resp User
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return User{}, fmt.Errorf("failed to get user info: %w", err)
	}
	return resp, nil
}
