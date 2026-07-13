package github

import "net/http"

// DefaultBaseURL is the default GitHub REST API endpoint used by the client.
//
// This value is used unless overridden with the WithBaseURL option.
const DefaultBaseURL = "https://api.github.com"

// Client is a GitHub REST API client.
//
// A Client manages authentication, HTTP communication, and request execution
// for the GitHub REST API. A single Client is safe to reuse across multiple
// requests.
type Client struct {
	token      string
	httpClient *http.Client
	baseURL    string
}

// Option configures a Client during construction.
//
// Options allow callers to customize client behavior, such as providing a
// custom HTTP client or overriding the GitHub API base URL.
type Option func(*Client)

// WithHTTPClient configures the Client to use the provided HTTP client.
//
// This can be used to customize request behavior, such as setting timeouts,
// transport settings, proxies, or custom middleware.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithBaseURL configures the Client to use a custom GitHub API base URL.
//
// This is primarily intended for GitHub Enterprise Server or testing against
// mock HTTP servers.
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// New creates a new GitHub API client.
//
// The provided token is used to authenticate requests using the GitHub REST
// API. By default, the client uses http.DefaultClient and DefaultBaseURL.
//
// Client behavior may be customized by supplying one or more Option values.
//
// Parameters:
//   - token: GitHub personal access token or other supported authentication token.
//   - opts: Optional client configuration functions.
//
// Returns a configured GitHub API client.
func New(token string, opts ...Option) *Client {
	c := &Client{
		token:      token,
		httpClient: http.DefaultClient,
		baseURL:    DefaultBaseURL,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// BaseURL returns the GitHub API base URL used by the client.
//
// This is either DefaultBaseURL or the value configured with WithBaseURL.
func (c *Client) BaseURL() string {
	return c.baseURL
}

// HTTPClient returns the HTTP client used to execute requests.
//
// If no custom HTTP client was provided during construction, this returns
// http.DefaultClient.
func (c *Client) HTTPClient() *http.Client {
	return c.httpClient
}
