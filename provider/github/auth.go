package github

import "net/http"

// SetAuthHeaders adds the authentication and required API headers to an
// outgoing HTTP request.
//
// It sets the Authorization header using the client's bearer token and
// includes the headers required for communicating with the GitHub REST API.
//
// Headers added:
//   - Authorization: Bearer <token>
//   - Accept: application/vnd.github+json
//   - X-GitHub-Api-Version: 2026-03-10
func (c *Client) SetAuthHeaders(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2026-03-10")
}
