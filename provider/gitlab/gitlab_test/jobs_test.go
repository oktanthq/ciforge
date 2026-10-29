package gitlab_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oktanthq/ciforge/provider/gitlab"
)

func TestListJobsForProject(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/jobs")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)
		if got := r.URL.Query()["scope[]"]; len(got) != 2 || got[0] != "failed" || got[1] != "running" {
			t.Errorf("scope[] = %v, want [failed running]", r.URL.Query()["scope[]"])
		}
		assertQuery(t, r, "order_by", "id")
		assertQuery(t, r, "sort", "asc")

		json.NewEncoder(w).Encode([]map[string]any{
			{"id": 1, "name": "test", "status": "failed", "stage": "test", "created_at": "2026-01-02T15:04:05Z"},
			{"id": 2, "name": "build", "status": "running", "stage": "build", "created_at": "2026-01-02T15:04:05Z"},
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	jobs, err := client.ListJobsForProject(context.Background(), 42, &gitlab.ListJobsForProjectOptions{
		Scope:   []gitlab.JobScope{gitlab.JobScopeFailed, gitlab.JobScopeRunning},
		OrderBy: "id",
		Sort:    gitlab.SortAscending,
	})
	assertNoError(t, err)
	assertEqual(t, 2, len(jobs))
	assertEqual(t, "test", jobs[0].Name)
	assertEqual(t, "failed", jobs[0].Status)
}

func TestListJobsForPipeline(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipelines/100/jobs")
		assertMethod(t, r, http.MethodGet)
		assertQuery(t, r, "include_retried", "true")

		json.NewEncoder(w).Encode([]map[string]any{
			{"id": 7, "name": "lint", "status": "success", "stage": "test"},
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	jobs, err := client.ListJobsForPipeline(context.Background(), 42, 100, &gitlab.ListJobsForPipelineOptions{
		IncludeRetried: true,
	})
	assertNoError(t, err)
	assertEqual(t, 1, len(jobs))
	assertEqual(t, 7, jobs[0].ID)
}

func TestListTriggerJobsForPipeline(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipelines/100/trigger_jobs")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode([]map[string]any{
			{"id": 8, "name": "trigger-child", "status": "success", "downstream_pipeline": map[string]any{"id": 9, "status": "success"}},
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	jobs, err := client.ListTriggerJobsForPipeline(context.Background(), 42, 100, nil)
	assertNoError(t, err)
	assertEqual(t, 1, len(jobs))
	assertEqual(t, "trigger-child", jobs[0].Name)
	assertEqual(t, 9, jobs[0].DownstreamPipeline.ID)
}

func TestGetJob(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/jobs/7")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{"id": 7, "name": "lint", "status": "success", "stage": "test"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	job, err := client.GetJob(context.Background(), 42, 7)
	assertNoError(t, err)
	assertEqual(t, 7, job.ID)
	assertEqual(t, "lint", job.Name)
}

func TestGetJobByToken(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/job")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{"id": 11, "name": "current", "status": "running"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	job, err := client.GetJobByToken(context.Background())
	assertNoError(t, err)
	assertEqual(t, 11, job.ID)
}

func TestGetAllowedAgents(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/job/allowed_agents")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"allowed_agents": []map[string]any{{"id": 1, "config_project": map[string]any{"id": 42, "name": "cfg"}}},
			"job":            map[string]any{"id": 11},
			"pipeline":       map[string]any{"id": 100},
			"project":        map[string]any{"id": 42, "groups": []map[string]any{{"id": 3}}},
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	resp, err := client.GetAllowedAgents(context.Background())
	assertNoError(t, err)
	assertEqual(t, 1, len(resp.AllowedAgents))
	assertEqual(t, 1, resp.AllowedAgents[0].ID)
}

func TestGetJobTrace(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/jobs/7/trace")
		assertMethod(t, r, http.MethodGet)
		w.Write([]byte("job log line\n"))
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	data, err := client.GetJobTrace(context.Background(), 42, 7)
	assertNoError(t, err)
	assertEqual(t, "job log line\n", string(data))
}

func TestCancelJob(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/jobs/7/cancel")
		assertMethod(t, r, http.MethodPost)
		assertQuery(t, r, "force", "true")

		json.NewEncoder(w).Encode(map[string]any{"id": 7, "status": "canceled"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	job, err := client.CancelJob(context.Background(), 42, 7, &gitlab.CancelJobOptions{Force: true})
	assertNoError(t, err)
	assertEqual(t, "canceled", job.Status)
}

func TestRetryJob(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/jobs/7/retry")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{"id": 8, "name": "lint", "status": "pending"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	job, err := client.RetryJob(context.Background(), 42, 7, nil)
	assertNoError(t, err)
	assertEqual(t, 8, job.ID)
	assertEqual(t, "pending", job.Status)
}

func TestEraseJob(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/jobs/7/erase")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{"id": 7, "status": "success", "erased_at": "2026-01-02T15:04:05Z"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	job, err := client.EraseJob(context.Background(), 42, 7)
	assertNoError(t, err)
	assertEqual(t, 7, job.ID)
}

func TestPlayJob(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/jobs/7/play")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{"id": 7, "status": "pending", "name": "deploy"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	job, err := client.PlayJob(context.Background(), 42, 7, &gitlab.PlayJobBody{
		JobVariablesAttributes: []gitlab.Variable{{Key: "ENV", Value: "prod"}},
	})
	assertNoError(t, err)
	assertEqual(t, "pending", job.Status)
}
