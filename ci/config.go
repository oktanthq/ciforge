package ci

import "time"

type Config struct {
	BaseURL   string
	Timeout   time.Duration
	UserAgent string
	Debug     bool
}

func defaultConfig() Config {
	return Config{
		Timeout:   30 * time.Second,
		UserAgent: "ciforge",
	}
}
