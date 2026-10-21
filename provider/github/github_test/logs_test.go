package github_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oktanthq/ciforge/provider/github"
)

func TestGetWorkflowRunLogs(t *testing.T) {
	t.Parallel()

	zipData := createTestZip(map[string]string{
		"step_1.log": "Hello from step 1\n",
		"step_2.log": "Hello from step 2\n",
	})

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/runs/100/logs")
		assertMethod(t, r, http.MethodGet)
		w.Header().Set("Content-Type", "application/zip")
		w.Write(zipData)
	}))
	defer srv.Close()

	client := github.New("token", github.WithBaseURL(srv.URL), github.WithHTTPClient(srv.Client()))
	logs, err := client.GetWorkflowRunLogs(context.Background(), "repo", "owner", 100)
	assertNoError(t, err)
	assertEqual(t, true, len(logs) > 0)
}
