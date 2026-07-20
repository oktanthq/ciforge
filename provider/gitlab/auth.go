package gitlab

import "net/http"

// setAuthHeaders adds the authentication and required API headers to an
// outgoing HTTP request.
//
// It sets the Authorization header using the client's private token and
// includes the headers required for communicating with the GitLab REST API.
//
// Headers added:
//   - PRIVATE-TOKEN: <token>
func (c *Client) setAuthHeaders(req *http.Request) {
	req.Header.Set("PRIVATE-TOKEN", c.token)
}
