package github_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abdulahshoaib/ciforge/provider/github"
)

func TestAuthHeaders(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, "Bearer test-token", r.Header.Get("Authorization"))
		assertEqual(t, "application/vnd.github+json", r.Header.Get("Accept"))
		assertEqual(t, "2026-03-10", r.Header.Get("X-GitHub-Api-Version"))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"total_count": 0, "workflows": []interface{}{}})
	}))
	defer srv.Close()

	client := github.New("test-token", github.WithBaseURL(srv.URL), github.WithHTTPClient(srv.Client()))
	_, err := client.ListWorkflows(context.Background(), "repo", "owner")
	assertNoError(t, err)
}

func TestNotFoundError(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "Not Found"})
	}))
	defer srv.Close()

	client := github.New("token", github.WithBaseURL(srv.URL), github.WithHTTPClient(srv.Client()))
	_, err := client.GetWorkflow(context.Background(), "repo", "owner", "999")
	assertError(t, err)
}

func TestUnauthorizedError(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"message": "Bad credentials"})
	}))
	defer srv.Close()

	client := github.New("bad-token", github.WithBaseURL(srv.URL), github.WithHTTPClient(srv.Client()))
	_, err := client.ListRunsForRepo(context.Background(), "repo", "owner")
	assertError(t, err)
}
