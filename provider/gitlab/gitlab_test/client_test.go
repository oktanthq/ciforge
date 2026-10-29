package gitlab_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/oktanthq/ciforge/provider/gitlab"
)

func TestNewDefaultBaseURL(t *testing.T) {
	t.Parallel()

	client := gitlab.New("token")
	assertEqual(t, gitlab.DefaultBaseURL, client.BaseURL())
}

func TestWithHTTPClient(t *testing.T) {
	t.Parallel()

	client := gitlab.New("token", gitlab.WithHTTPClient(&http.Client{Timeout: 5 * time.Second}))
	assertEqual(t, 5*time.Second, client.HTTPClient().Timeout)
}

func TestWithBaseURL(t *testing.T) {
	t.Parallel()

	client := gitlab.New("token", gitlab.WithBaseURL("https://gitlab.example.com/api/v4"))
	assertEqual(t, "https://gitlab.example.com/api/v4", client.BaseURL())
}
