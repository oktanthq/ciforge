package ci

import "errors"

var (
	ErrUnauthorized        = errors.New("unauthorized")
	ErrNotFound            = errors.New("resource not found")
	ErrRateLimited         = errors.New("rate limited")
	ErrProviderUnavailable = errors.New("provider unavailable")
	ErrInvalidProvider     = errors.New("invalid provider")
)
