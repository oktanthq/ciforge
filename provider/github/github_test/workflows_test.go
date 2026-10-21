package github_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/oktanthq/ciforge/provider/github"
)

func TestListWorkflows(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/workflows")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"total_count": 2,
			"workflows": []map[string]interface{}{
				{"id": 1, "name": "CI", "state": "active", "created_at": time.Now(), "updated_at": time.Now(), "html_url": "https://github.com/owner/repo/actions/workflows/ci.yml"},
				{"id": 2, "name": "Deploy", "state": "disabled", "created_at": time.Now(), "updated_at": time.Now(), "html_url": "https://github.com/owner/repo/actions/workflows/deploy.yml"},
			},
		})
	}))
	defer srv.Close()

	client := github.New("test-token", github.WithBaseURL(srv.URL), github.WithHTTPClient(srv.Client()))
	workflows, err := client.ListWorkflows(context.Background(), "repo", "owner")
	assertNoError(t, err)
	assertEqual(t, 2, len(workflows))
	assertEqual(t, int64(1), workflows[0].ID)
	assertEqual(t, "CI", workflows[0].Name)
	assertEqual(t, "active", workflows[0].State)
	assertEqual(t, "disabled", workflows[1].State)
}

func TestGetWorkflow(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/workflows/42")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"id": 42, "name": "CI", "state": "active", "created_at": time.Now(), "updated_at": time.Now(),
			"html_url": "https://github.com/owner/repo/actions/workflows/ci.yml",
		})
	}))
	defer srv.Close()

	client := github.New("test-token", github.WithBaseURL(srv.URL), github.WithHTTPClient(srv.Client()))
	workflow, err := client.GetWorkflow(context.Background(), "repo", "owner", 42)
	assertNoError(t, err)
	assertEqual(t, int64(42), workflow.ID)
	assertEqual(t, "CI", workflow.Name)
}
