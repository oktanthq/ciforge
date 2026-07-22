package httpclient

import "net/http"

// DefaultBaseURL is the default API endpoint used by the client.
//
// This value is used unless overridden with the WithBaseURL option.
const DefaultBaseURL = ""

// Base is the base HTTP client shared across all CI provider implementations.
//
// A Base manages authentication, HTTP communication, and request execution.
// It is safe to reuse across multiple requests. Provider-specific clients
// should embed Base and implement setAuthHeaders for their API's auth scheme.
type Base struct {
	Token      string
	HTTPClient *http.Client
	BaseURL    string
}

// Option configures a Base during construction.
//
// Options allow callers to customize client behavior, such as providing a
// custom HTTP client or overriding the API base URL.
type Option func(*Base)

// WithHTTPClient configures the Base to use the provided HTTP client.
//
// This can be used to customize request behavior, such as setting timeouts,
// transport settings, proxies, or custom middleware.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(b *Base) {
		b.HTTPClient = httpClient
	}
}

// WithBaseURL configures the Base to use a custom API base URL.
//
// This is primarily intended for self-hosted instances or testing against
// mock HTTP servers.
func WithBaseURL(baseURL string) Option {
	return func(b *Base) {
		b.BaseURL = baseURL
	}
}

// NewBase creates a new Base with the provided token and options.
//
// By default, the Base uses http.DefaultClient and an empty base URL.
// Provider-specific clients should call this and then set their own BaseURL.
func NewBase(token string, baseURL string, opts ...Option) Base {
	b := Base{
		Token:      token,
		HTTPClient: http.DefaultClient,
		BaseURL:    baseURL,
	}
	for _, opt := range opts {
		opt(&b)
	}
	return b
}

// GetBaseURL returns the API base URL used by the client.
func (b *Base) GetBaseURL() string {
	return b.BaseURL
}

// GetHTTPClient returns the HTTP client used to execute requests.
//
// If no custom HTTP client was provided during construction, this returns
// http.DefaultClient.
func (b *Base) GetHTTPClient() *http.Client {
	return b.HTTPClient
}
