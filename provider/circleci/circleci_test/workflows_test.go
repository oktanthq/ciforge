package circleci_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oktanthq/ciforge/provider/circleci"
)

func TestGetWorkflow(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/workflow/w1")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		json.NewEncoder(w).Encode(map[string]any{
			"id": "w1", "name": "build-test-deploy", "status": "running", "pipeline_id": "p1",
			"created_at": "2019-08-24T14:15:22Z",
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	wf, err := client.GetWorkflow(context.Background(), "w1")
	assertNoError(t, err)
	assertEqual(t, "w1", wf.ID)
	assertEqual(t, "build-test-deploy", wf.Name)
	assertEqual(t, "running", wf.Status)
}

func TestGetWorkflowJobs(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/workflow/w1/job")
		assertMethod(t, r, http.MethodGet)
		assertQuery(t, r, "page-token", "abc")

		json.NewEncoder(w).Encode(map[string]any{
			"items": []map[string]any{
				{"id": "j1", "name": "test", "status": "success", "job_number": 12},
				{"id": "j2", "name": "build", "status": "running", "job_number": 13},
			},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	jobs, err := client.GetWorkflowJobs(context.Background(), "w1", &circleci.ListWorkflowJobsOptions{PageToken: "abc"})
	assertNoError(t, err)
	assertEqual(t, 2, len(jobs))
	assertEqual(t, "test", jobs[0].Name)
}

func TestCancelWorkflow(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/workflow/w1/cancel")
		assertMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusAccepted)
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.CancelWorkflow(context.Background(), "w1")
	assertNoError(t, err)
}

func TestApproveJob(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/workflow/w1/approve/hold")
		assertMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusAccepted)
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.ApproveJob(context.Background(), "w1", "hold")
	assertNoError(t, err)
}

func TestRerunWorkflow(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/workflow/w1/rerun")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{"workflow_id": "w2"})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	id, err := client.RerunWorkflow(context.Background(), "w1", circleci.RerunWorkflowBody{FromFailed: true})
	assertNoError(t, err)
	assertEqual(t, circleci.WorkflowID("w2"), id)
}
