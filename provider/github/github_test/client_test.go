package github_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/abdulahshoaib/ciforge/provider/github"
)

func TestNewDefaultBaseURL(t *testing.T) {
	t.Parallel()

	client := github.New("token")
	assertEqual(t, github.DefaultBaseURL, client.BaseURL())
}

func TestWithHTTPClient(t *testing.T) {
	t.Parallel()

	client := github.New("token", github.WithHTTPClient(&http.Client{Timeout: 5 * time.Second}))
	assertEqual(t, 5*time.Second, client.HTTPClient().Timeout)
}

func TestWithBaseURL(t *testing.T) {
	t.Parallel()

	client := github.New("token", github.WithBaseURL("https://custom.url"))
	assertEqual(t, "https://custom.url", client.BaseURL())
}
