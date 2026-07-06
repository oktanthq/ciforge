package ci

type Client struct {
	provider   Provider
	config     Config
	middleware []Middleware
}

func New(opts ...Option) *Client {
	cfg := defaultConfig()

	client := &Client{
		config: cfg,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
