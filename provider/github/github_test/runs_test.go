package github_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/abdulahshoaib/ciforge/provider/github"
)

func TestListRunsForRepo(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/runs")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"total_count": 1,
			"workflow_runs": []map[string]interface{}{
				{
					"id": 100, "run_number": 5, "workflow_id": 1,
					"status": "completed", "conclusion": "success",
					"head_branch": "main", "event": "push",
					"created_at": time.Now().Add(-10 * time.Minute),
					"updated_at": time.Now(),
				},
			},
		})
	}))
	defer srv.Close()

	client := github.New("token", github.WithBaseURL(srv.URL), github.WithHTTPClient(srv.Client()))
	runs, err := client.ListRunsForRepo(context.Background(), "repo", "owner", github.ListRunsForRepoOptions{})
	assertNoError(t, err)
	assertEqual(t, 1, len(runs))
	assertEqual(t, int64(100), runs[0].ID)
	assertEqual(t, "completed", runs[0].Status)
	assertEqual(t, "success", runs[0].Conclusion)
	assertEqual(t, "push", runs[0].Event)
}

func TestGetWorkflowRun(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/runs/100")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"id": 100, "run_number": 5, "workflow_id": 1,
			"status": "in_progress", "conclusion": "",
			"head_branch": "feature", "event": "pull_request",
			"created_at": time.Now().Add(-5 * time.Minute),
			"updated_at": time.Now(),
		})
	}))
	defer srv.Close()

	client := github.New("token", github.WithBaseURL(srv.URL), github.WithHTTPClient(srv.Client()))
	run, err := client.GetWorkflowRun(context.Background(), "repo", "owner", 100)
	assertNoError(t, err)
	assertEqual(t, int64(100), run.ID)
	assertEqual(t, "in_progress", run.Status)
	assertEqual(t, "pull_request", run.Event)
}

func TestCancelWorkflowRun(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/runs/100/cancel")
		assertMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusAccepted)
	}))
	defer srv.Close()

	client := github.New("token", github.WithBaseURL(srv.URL), github.WithHTTPClient(srv.Client()))
	err := client.CancelWorkflowRun(context.Background(), "repo", "owner", 100)
	assertNoError(t, err)
}

func TestReRunWorkflow(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/runs/100/rerun")
		assertMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
	}))
	defer srv.Close()

	client := github.New("token", github.WithBaseURL(srv.URL), github.WithHTTPClient(srv.Client()))
	err := client.ReRunWorkflow(context.Background(), "repo", "owner", 100)
	assertNoError(t, err)
}
