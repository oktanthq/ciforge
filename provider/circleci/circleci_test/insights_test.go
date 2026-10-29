package circleci_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oktanthq/ciforge/provider/circleci"
)

func TestListWorkflowMetrics(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/insights/gh/org/proj/workflows")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)
		assertQuery(t, r, "branch", "main")
		assertQuery(t, r, "date-range", "30d")

		json.NewEncoder(w).Encode(map[string]any{
			"items": []map[string]any{
				{"name": "build", "metrics": map[string]any{"total_runs": 10, "success_rate": 0.9, "throughput": 1.2, "total_credits_used": 100}},
			},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	metrics, err := client.ListWorkflowMetrics(context.Background(), "gh/org/proj", &circleci.ListWorkflowMetricsOptions{
		Branch:    "main",
		DateRange: "30d",
	})
	assertNoError(t, err)
	assertEqual(t, 1, len(metrics))
	assertEqual(t, "build", metrics[0].Name)
	assertEqual(t, int64(10), metrics[0].Metrics.TotalRuns)
}

func TestListWorkflowRuns(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/insights/gh/org/proj/workflows/build")
		assertMethod(t, r, http.MethodGet)
		assertQuery(t, r, "branch", "main")

		json.NewEncoder(w).Encode(map[string]any{
			"items": []map[string]any{
				{"id": "run-1", "status": "success", "credits_used": 12, "duration": 40.5},
			},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	runs, err := client.ListWorkflowRuns(context.Background(), "gh/org/proj", "build", &circleci.ListWorkflowRunsOptions{Branch: "main"})
	assertNoError(t, err)
	assertEqual(t, 1, len(runs))
	assertEqual(t, "success", runs[0].Status)
}

func TestListJobMetrics(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/insights/gh/org/proj/workflows/build/jobs")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"items": []map[string]any{
				{"name": "test", "metrics": map[string]any{"total_runs": 8, "success_rate": 1.0}},
			},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	metrics, err := client.ListJobMetrics(context.Background(), "gh/org/proj", "build", nil)
	assertNoError(t, err)
	assertEqual(t, 1, len(metrics))
	assertEqual(t, "test", metrics[0].Name)
}
