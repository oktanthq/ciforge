package circleci_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/oktanthq/ciforge/provider/circleci"
)

func TestNewDefaultBaseURL(t *testing.T) {
	t.Parallel()

	client := circleci.New("token")
	assertEqual(t, circleci.DefaultBaseURL, client.BaseURL())
}

func TestWithHTTPClient(t *testing.T) {
	t.Parallel()

	client := circleci.New("token", circleci.WithHTTPClient(&http.Client{Timeout: 5 * time.Second}))
	assertEqual(t, 5*time.Second, client.HTTPClient().Timeout)
}

func TestWithBaseURL(t *testing.T) {
	t.Parallel()

	client := circleci.New("token", circleci.WithBaseURL("https://circleci.example.com/api/v2"))
	assertEqual(t, "https://circleci.example.com/api/v2", client.BaseURL())
}
