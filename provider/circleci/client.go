package circleci

import (
	"net/http"

	"github.com/abdulahshoaib/ciforge/provider/internal/httpclient"
)

// DefaultBaseURL is the default CircleCI REST API endpoint used by the client.
//
// This value is used unless overridden with the WithBaseURL option.
const DefaultBaseURL = "https://circleci.com/api/v2"

// Client is a CircleCI REST API client.
//
// A Client manages authentication, HTTP communication, and request execution
// for the CircleCI REST API. A single Client is safe to reuse across multiple
// requests.
type Client struct {
	httpclient.Base
}

// Option configures a Client during construction.
//
// Options allow callers to customize client behavior, such as providing a
// custom HTTP client or overriding the CircleCI API base URL.
type Option = httpclient.Option

// WithHTTPClient configures the Client to use the provided HTTP client.
//
// This can be used to customize request behavior, such as setting timeouts,
// transport settings, proxies, or custom middleware.
func WithHTTPClient(httpClient *http.Client) Option {
	return httpclient.WithHTTPClient(httpClient)
}

// WithBaseURL configures the Client to use a custom CircleCI API base URL.
//
// This is primarily intended for CircleCI Server or testing against
// mock HTTP servers.
func WithBaseURL(baseURL string) Option {
	return httpclient.WithBaseURL(baseURL)
}

// New creates a new CircleCI API client.
//
// The provided token is used to authenticate requests using the CircleCI REST
// API. By default, the client uses http.DefaultClient and DefaultBaseURL.
//
// Client behavior may be customized by supplying one or more Option values.
//
// Parameters:
//   - token: CircleCI personal API token.
//   - opts: Optional client configuration functions.
//
// Returns a configured CircleCI API client.
func New(token string, opts ...Option) *Client {
	return &Client{
		Base: httpclient.NewBase(token, DefaultBaseURL, opts...),
	}
}

// BaseURL returns the CircleCI API base URL used by the client.
//
// This is either DefaultBaseURL or the value configured with WithBaseURL.
func (c *Client) BaseURL() string {
	return c.Base.GetBaseURL()
}

// HTTPClient returns the HTTP client used to execute requests.
//
// If no custom HTTP client was provided during construction, this returns
// http.DefaultClient.
func (c *Client) HTTPClient() *http.Client {
	return c.Base.GetHTTPClient()
}
