package circleci

import "net/http"

// SetAuthHeaders adds the authentication and required API headers to an
// outgoing HTTP request.
//
// It sets the Circle-Token header using the client's API token.
//
// Headers added:
//   - Circle-Token: <token>
func (c *Client) SetAuthHeaders(req *http.Request) {
	req.Header.Set("Circle-Token", c.Token)
}
