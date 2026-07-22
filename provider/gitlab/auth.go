package gitlab

import "net/http"

// SetAuthHeaders adds the authentication and required API headers to an
// outgoing HTTP request.
//
// It sets the Authorization header using the client's private token and
// includes the headers required for communicating with the GitLab REST API.
//
// Headers added:
//   - PRIVATE-TOKEN: <token>
func (c *Client) SetAuthHeaders(req *http.Request) {
	req.Header.Set("PRIVATE-TOKEN", c.Token)
}
