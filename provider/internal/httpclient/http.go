package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Sentinel errors for common API failure modes.
//
// Callers can use errors.Is to check for specific error conditions
// returned by the API.
var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrNotFound     = errors.New("resource not found")
	ErrRateLimited  = errors.New("rate limited")
)

// apiError represents a standard JSON error response from a CI provider API.
type apiError struct {
	Message string `json:"message"`
}

// Doer is the interface that wraps the core HTTP execution methods.
//
// Provider-specific clients implement this interface to delegate HTTP
// operations while retaining control over authentication headers.
type Doer interface {
	// SetAuthHeaders adds provider-specific authentication headers to a request.
	SetAuthHeaders(req *http.Request)

	// HandleError converts a non-success HTTP response into a typed Go error.
	HandleError(resp *http.Response) error
}

// Do executes an HTTP request and optionally decodes the JSON response.
//
// If body is non-nil, it is JSON-encoded and sent as the request body with
// the Content-Type header set to application/json.
//
// If out is non-nil, the response body is expected to contain JSON and is
// decoded into out.
//
// Responses with a status code >= 300 are converted into typed errors via
// the doer's HandleError method.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - method: HTTP method (GET, POST, PUT, DELETE, etc.).
//   - url: Full request URL.
//   - body: Request payload to be JSON-encoded, or nil.
//   - out: Destination for the decoded JSON response, or nil.
//   - doer: The provider-specific client for auth and error handling.
//
// Returns an error if the request cannot be created, executed, the response
// indicates an API error, or the response JSON cannot be decoded.
func Do(ctx context.Context, b *Base, doer Doer, method, url string, body, out any) error {
	var reqBody io.Reader
	if body != nil {
		p, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewReader(p)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	doer.SetAuthHeaders(req)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := b.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusMultipleChoices {
		return doer.HandleError(resp)
	}

	if out != nil {
		if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// DoRaw executes an HTTP request and returns the raw response body.
//
// Unlike Do, this method does not attempt to decode the response as JSON.
// It is intended for endpoints that return arbitrary content such as logs,
// configuration files, archives, or other binary/text data.
//
// Responses with a status code >= 300 are converted into typed errors via
// the doer's HandleError method.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - method: HTTP method.
//   - url: Full request URL.
//   - doer: The provider-specific client for auth and error handling.
//
// Returns the response body as a byte slice or an error if the request fails,
// the API returns an error status, or the response body cannot be read.
func DoRaw(ctx context.Context, b *Base, doer Doer, method, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	doer.SetAuthHeaders(req)

	resp, err := b.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusMultipleChoices {
		return nil, doer.HandleError(resp)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return data, nil
}

// HandleAPIError converts a non-success HTTP response into a typed Go error
// using a standard JSON error body.
//
// The response body is expected to contain a JSON error object with a "message"
// field. If the body cannot be decoded, the HTTP status text is used instead.
//
// Common API status codes are mapped to package-level sentinel errors,
// allowing callers to use errors.Is for error handling:
//   - 401 -> ErrUnauthorized
//   - 404 -> ErrNotFound
//   - 429 -> ErrRateLimited
//
// All other status codes return an error containing the HTTP status text
// and the API's error message.
func HandleAPIError(resp *http.Response) error {
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
