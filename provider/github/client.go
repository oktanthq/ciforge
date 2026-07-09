package github

import (
	"net/http"

	"github.com/abdulahshoaib/ciforge/ci"
)

const DefaultBaseURL = "https://api.github.com"

type Client struct {
	token      string
	httpClient *http.Client
	baseURL    string
}

type Option func(*Client)

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

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

func (c *Client) BaseURL() string {
	return c.baseURL
}

func (c *Client) HTTPClient() *http.Client {
	return c.httpClient
}

var _ ci.Provider = (*Client)(nil)
