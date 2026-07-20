package gitlab

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrNotFound     = errors.New("resource not found")
	ErrRateLimited  = errors.New("rate limited")
)

type apiError struct {
	Message string `json:"message"`
}

// do executes an HTTP request against the GitLab API.
//
// If body is non-nil, it is JSON-encoded and sent as the request body with
// the Content-Type header set to application/json.
//
// The request is authenticated using the client's configured credentials.
// Responses with a status code >= 300 are converted into typed errors via
// handleError.
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
	url := c.baseURL + path

	var reqBody io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	c.setAuthHeaders(req)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusMultipleChoices {
		return c.handleError(resp)
	}

	if out != nil {
		if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// doRaw executes an HTTP request and returns the raw response body.
//
// Unlike do, this method does not attempt to decode the response as JSON.
// It is intended for endpoints that return arbitrary content such as logs,
// configuration files, archives, or other binary/text data.
//
// The request is authenticated using the client's configured credentials.
// Responses with a status code >= 300 are converted into typed errors via
// handleError.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - method: HTTP method.
//   - path: API endpoint path relative to the client's base URL.
//
// Returns the response body as a byte slice or an error if the request fails,
// the API returns an error status, or the response body cannot be read.
func (c *Client) doRaw(ctx context.Context, method, path string) ([]byte, error) {
	url := c.baseURL + path

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	c.setAuthHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusMultipleChoices {
		return nil, c.handleError(resp)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return data, nil
}

// handleError converts a non-success HTTP response into a typed Go error.
//
// The response body is expected to contain a JSON error object matching
// apiError. If the body cannot be decoded, the HTTP status text is used as
// the error message instead.
//
// Common API status codes are mapped to package-specific sentinel errors,
// allowing callers to use errors.Is for error handling.
//
// Mappings:
//   - 401 -> ErrUnauthorized
//   - 404 -> ErrNotFound
//   - 429 -> ErrRateLimited
//
// All other status codes return an error containing the HTTP status text
// and the API's error message.
func (c *Client) handleError(resp *http.Response) error {
	var apiErr apiError
	if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
		apiErr.Message = http.StatusText(resp.StatusCode)
	}

	msg := apiErr.Message

	switch resp.StatusCode {
	case http.StatusUnauthorized:
		return fmt.Errorf("%w: %s", ErrUnauthorized, msg)
	case http.StatusNotFound:
		return fmt.Errorf("%w: %s", ErrNotFound, msg)
	case http.StatusTooManyRequests:
		return fmt.Errorf("%w: %s", ErrRateLimited, msg)
	default:
		return fmt.Errorf("%s: %s", http.StatusText(resp.StatusCode), msg)
	}
}
