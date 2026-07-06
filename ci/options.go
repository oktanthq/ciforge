package ci

import "time"

type Option func(*Client)

func WithProvider(p Provider) Option
func WithTimeout(timeout time.Duration) Option
func WithUserAgent(agent string) Option
func WithBaseURL(url string) Option
func WithDebug(enabled bool) Option
func WithMiddleware(m ...Middleware) Option
