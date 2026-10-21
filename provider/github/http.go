package github

import (
	"context"
	"net/http"

	"github.com/oktanthq/ciforge/provider/internal/httpclient"
)

// Sentinel errors for common API failure modes.
//
// Callers can use errors.Is to check for specific error conditions
// returned by the GitHub API.
var (
	ErrUnauthorized = httpclient.ErrUnauthorized
	ErrNotFound     = httpclient.ErrNotFound
	ErrRateLimited  = httpclient.ErrRateLimited
)

// HandleError converts a non-success HTTP response into a typed Go error.
//
// This method implements the httpclient.Doer interface, allowing the GitHub
// client to delegate HTTP operations to the shared httpclient package.
func (c *Client) HandleError(resp *http.Response) error {
	return httpclient.HandleAPIError(resp)
}

// do executes an HTTP request against the GitHub API.
//
// If body is non-nil, it is JSON-encoded and sent as the request body with
// the Content-Type header set to application/json.
//
// The request is authenticated using the client's configured credentials.
// Responses with a status code >= 300 are converted into typed errors via
// HandleError.
//
// If out is non-nil, the response body is expected to contain JSON and is
// decoded into out.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - method: HTTP method (GET, POST, PUT, DELETE, etc.).
//   - path: API endpoint path relative to the client's base URL.
//   - body: Request payload to be JSON-encoded, or nil.
//   - out: Destination for the decoded JSON response, or nil.
//
// Returns an error if the request cannot be created, executed, the response
// indicates an API error, or the response JSON cannot be decoded.
func (c *Client) do(ctx context.Context, method, path string, body, out any) error {
	return httpclient.Do(ctx, &c.Base, c, method, c.Base.BaseURL+path, body, out)
}

// doRaw executes an HTTP request and returns the raw response body.
//
// Unlike do, this method does not attempt to decode the response as JSON.
// It is intended for endpoints that return arbitrary content such as logs,
// configuration files, archives, or other binary/text data.
//
// The request is authenticated using the client's configured credentials.
// Responses with a status code >= 300 are converted into typed errors via
// HandleError.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - method: HTTP method.
//   - path: API endpoint path relative to the client's base URL.
//
// Returns the response body as a byte slice or an error if the request fails,
// the API returns an error status, or the response body cannot be read.
func (c *Client) doRaw(ctx context.Context, method, path string) ([]byte, error) {
	return httpclient.DoRaw(ctx, &c.Base, c, method, c.Base.BaseURL+path)
}

// doRawURL executes an HTTP request against the provided URL and returns the
// raw response body.
//
// Unlike doRaw, this method does not prepend the client's base URL, making it
// suitable for requests to fully qualified URLs such as artifact download
// links.
//
// The request is authenticated using the client's configured credentials.
// Responses with a status code >= 300 are converted into typed errors via
// HandleError.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - method: HTTP method.
//   - url: Fully qualified request URL.
//
// Returns the response body as a byte slice or an error if the request fails,
// the API returns a non-success status, or the response body cannot be read.
func (c *Client) doRawURL(ctx context.Context, method, url string) ([]byte, error) {
	return httpclient.DoRaw(ctx, &c.Base, c, method, url)
}
