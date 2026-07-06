package ci

import "time"

type Option func(*Client)

func WithProvider(p Provider) Option {
	return func(c *Client) {
		c.provider = p
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.config.Timeout = timeout
	}
}

func WithUserAgent(agent string) Option {
	return func(c *Client) {
		c.config.UserAgent = agent
	}
}

func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.config.BaseURL = url
	}
}

func WithDebug(enabled bool) Option {
	return func(c *Client) {
		c.config.Debug = enabled
	}
}

func WithMiddleware(m ...Middleware) Option {
	return func(c *Client) {
		c.middleware = append(c.middleware, m...)
	}
}
