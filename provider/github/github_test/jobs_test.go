package github_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abdulahshoaib/ciforge/models"
	"github.com/abdulahshoaib/ciforge/provider/github"
)

func TestListJobsForWorkflowRun(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/runs/100/jobs")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"total_count": 2,
			"jobs": []map[string]interface{}{
				{"id": 1, "run_id": 100, "name": "test", "status": "completed", "conclusion": "success"},
				{"id": 2, "run_id": 100, "name": "build", "status": "completed", "conclusion": "failure"},
			},
		})
	}))
	defer srv.Close()

	client := github.New("token", github.WithBaseURL(srv.URL), github.WithHTTPClient(srv.Client()))
	jobs, err := client.ListJobsForWorkflowRun(context.Background(), "repo", "owner", "100")
	assertNoError(t, err)
	assertEqual(t, 2, len(jobs))
	assertEqual(t, "test", jobs[0].Name)
	assertEqual(t, models.RunResultSuccess, jobs[0].Result)
	assertEqual(t, models.RunResultFailed, jobs[1].Result)
}
