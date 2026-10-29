package gitlab_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oktanthq/ciforge/provider/gitlab"
)

func TestListPipelines(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipelines")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)
		assertQuery(t, r, "ref", "main")
		assertQuery(t, r, "status", "success")

		json.NewEncoder(w).Encode([]map[string]any{
			{"id": 1, "iid": 10, "project_id": 42, "status": "success", "ref": "main", "sha": "abc123", "created_at": "2026-01-02T15:04:05Z", "updated_at": "2026-01-02T15:04:05Z"},
			{"id": 2, "iid": 11, "project_id": 42, "status": "failed", "ref": "main", "sha": "def456", "created_at": "2026-01-02T15:04:05Z", "updated_at": "2026-01-02T15:04:05Z"},
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	pipelines, err := client.ListPipelines(context.Background(), 42, &gitlab.ListPipelinesOptions{
		Ref:    "main",
		Status: gitlab.PipelineStatusSuccess,
	})
	assertNoError(t, err)
	assertEqual(t, 2, len(pipelines))
	assertEqual(t, 1, pipelines[0].ID)
	assertEqual(t, "success", pipelines[0].Status)
	assertEqual(t, "main", pipelines[0].Ref)
}

func TestGetPipeline(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipelines/100")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		json.NewEncoder(w).Encode(map[string]any{
			"id": 100, "iid": 10, "project_id": 42, "status": "running", "ref": "main", "sha": "abc123",
			"created_at": "2026-01-02T15:04:05Z", "updated_at": "2026-01-02T15:04:05Z",
			"detailed_status": map[string]any{"text": "running", "label": "running", "group": "running"},
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	pipeline, err := client.GetPipeline(context.Background(), 42, 100)
	assertNoError(t, err)
	assertEqual(t, 100, pipeline.ID)
	assertEqual(t, "running", pipeline.Status)
	assertEqual(t, "running", pipeline.DetailedStatus.Text)
}

func TestGetLatestPipeline(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipelines/latest")
		assertMethod(t, r, http.MethodGet)
		assertQuery(t, r, "ref", "main")

		json.NewEncoder(w).Encode(map[string]any{
			"id": 200, "project_id": 42, "status": "success", "ref": "main",
			"created_at": "2026-01-02T15:04:05Z", "updated_at": "2026-01-02T15:04:05Z",
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	pipeline, err := client.GetLatestPipeline(context.Background(), 42, &gitlab.GetLatestPipelineOptions{Ref: "main"})
	assertNoError(t, err)
	assertEqual(t, 200, pipeline.ID)
	assertEqual(t, "success", pipeline.Status)
}

func TestGetPipelineVariables(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipelines/100/variables")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode([]map[string]any{
			{"key": "FOO", "variable_type": "env_var", "value": "bar"},
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	vars, err := client.GetPipelineVariables(context.Background(), 42, 100)
	assertNoError(t, err)
	assertEqual(t, 1, len(vars))
	assertEqual(t, "FOO", vars[0].Key)
	assertEqual(t, "bar", vars[0].Value)
}

func TestGetPipelineTestReport(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipelines/100/test_report")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"total_time": 12, "total_count": 3, "success_count": 2, "failed_count": 1, "skipped_count": 0, "error_count": 0,
			"test_suites": []map[string]any{
				{"name": "pkg", "total_count": 3, "success_count": 2, "failed_count": 1},
			},
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	report, err := client.GetPipelineTestReport(context.Background(), 42, 100)
	assertNoError(t, err)
	assertEqual(t, 3, report.TotalCount)
	assertEqual(t, 2, report.SuccessCount)
	assertEqual(t, 1, len(report.TestSuites))
	assertEqual(t, "pkg", report.TestSuites[0].Name)
}

func TestGetPipelineTestReportSummary(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipelines/100/test_report_summary")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"total": map[string]any{"time": 12, "count": 3, "success": 2, "failed": 1, "skipped": 0, "error": 0},
			"test_suites": []map[string]any{
				{"name": "pkg", "total_count": 3, "success_count": 2, "failed_count": 1, "build_ids": []int{9}},
			},
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	summary, err := client.GetPipelineTestReportSummary(context.Background(), 42, 100)
	assertNoError(t, err)
	assertEqual(t, 3, summary.Total.Count)
	assertEqual(t, 1, len(summary.TestSuites))
}

func TestCreatePipeline(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipeline")
		assertMethod(t, r, http.MethodPost)
		assertQuery(t, r, "ref", "main")

		json.NewEncoder(w).Encode(map[string]any{
			"id": 300, "project_id": 42, "status": "created", "ref": "main",
			"created_at": "2026-01-02T15:04:05Z", "updated_at": "2026-01-02T15:04:05Z",
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	pipeline, err := client.CreatePipeline(context.Background(), 42, "main")
	assertNoError(t, err)
	assertEqual(t, 300, pipeline.ID)
	assertEqual(t, "created", pipeline.Status)
}

func TestRetryPipeline(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipelines/100/retry")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{
			"id": 100, "status": "pending",
			"created_at": "2026-01-02T15:04:05Z", "updated_at": "2026-01-02T15:04:05Z",
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	pipeline, err := client.RetryPipeline(context.Background(), 42, 100)
	assertNoError(t, err)
	assertEqual(t, "pending", pipeline.Status)
}

func TestCancelPipeline(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipelines/100/cancel")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{
			"id": 100, "status": "canceled",
			"created_at": "2026-01-02T15:04:05Z", "updated_at": "2026-01-02T15:04:05Z",
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	pipeline, err := client.CancelPipeline(context.Background(), 42, 100)
	assertNoError(t, err)
	assertEqual(t, "canceled", pipeline.Status)
}

func TestDeletePipeline(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipelines/100")
		assertMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	err := client.DeletePipeline(context.Background(), 42, 100)
	assertNoError(t, err)
}

func TestUpdatePipelineMetadata(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipelines/100/metadata")
		assertMethod(t, r, http.MethodPut)

		var body map[string]string
		json.NewDecoder(r.Body).Decode(&body)
		assertEqual(t, "nightly", body["name"])

		json.NewEncoder(w).Encode(map[string]any{
			"id": 100, "name": "nightly",
			"created_at": "2026-01-02T15:04:05Z", "updated_at": "2026-01-02T15:04:05Z",
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	pipeline, err := client.UpdatePipelineMetadata(context.Background(), 42, 100, "nightly")
	assertNoError(t, err)
	assertEqual(t, "nightly", pipeline.Name)
}
