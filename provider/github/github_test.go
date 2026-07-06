package github

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/abdulahshoaib/ciforge/models"
)

func TestListPipelines(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/workflows")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		json.NewEncoder(w).Encode(listWorkflowsResponse{
			TotalCount: 2,
			Workflows: []workflowDTO{
				{ID: 1, Name: "CI", State: "active", CreatedAt: time.Now(), UpdatedAt: time.Now(), HTMLURL: "https://github.com/owner/repo/actions/workflows/ci.yml"},
				{ID: 2, Name: "Deploy", State: "disabled", CreatedAt: time.Now(), UpdatedAt: time.Now(), HTMLURL: "https://github.com/owner/repo/actions/workflows/deploy.yml"},
			},
		})
	}))
	defer srv.Close()

	client := New("test-token", WithBaseURL(srv.URL), WithHTTPClient(srv.Client()))
	pipelines, err := client.ListPipelines(context.Background(), "owner/repo")
	assertNoError(t, err)
	assertEqual(t, 2, len(pipelines))
	assertEqual(t, "1", pipelines[0].ID)
	assertEqual(t, "CI", pipelines[0].Name)
	assertEqual(t, false, pipelines[0].Disabled)
	assertEqual(t, true, pipelines[1].Disabled)
}

func TestGetPipeline(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/workflows/42")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		json.NewEncoder(w).Encode(workflowDTO{
			ID: 42, Name: "CI", State: "active", CreatedAt: time.Now(), UpdatedAt: time.Now(),
			HTMLURL: "https://github.com/owner/repo/actions/workflows/ci.yml",
		})
	}))
	defer srv.Close()

	client := New("test-token", WithBaseURL(srv.URL), WithHTTPClient(srv.Client()))
	pipeline, err := client.GetPipeline(context.Background(), "owner/repo", "42")
	assertNoError(t, err)
	assertEqual(t, "42", pipeline.ID)
	assertEqual(t, "CI", pipeline.Name)
}

func TestListRuns(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/workflows/1/runs")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(listWorkflowRunsResponse{
			TotalCount: 1,
			WorkflowRuns: []workflowRunDTO{
				{
					ID: 100, RunNumber: 5, WorkflowID: 1,
					Status: "completed", Conclusion: "success",
					HeadBranch: "main", Event: "push",
					CreatedAt: time.Now().Add(-10 * time.Minute),
					UpdatedAt: time.Now(),
				},
			},
		})
	}))
	defer srv.Close()

	client := New("token", WithBaseURL(srv.URL), WithHTTPClient(srv.Client()))
	runs, err := client.ListRuns(context.Background(), "owner/repo", "1")
	assertNoError(t, err)
	assertEqual(t, 1, len(runs))
	assertEqual(t, "100", runs[0].ID)
	assertEqual(t, models.RunStatusCompleted, runs[0].Status)
	assertEqual(t, models.RunResultSuccess, runs[0].Result)
	assertEqual(t, models.TriggerPush, runs[0].Trigger)
}

func TestGetRuns(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/runs/100")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(workflowRunDTO{
			ID: 100, RunNumber: 5, WorkflowID: 1,
			Status: "in_progress", Conclusion: "",
			HeadBranch: "feature", Event: "pull_request",
			CreatedAt: time.Now().Add(-5 * time.Minute),
			UpdatedAt: time.Now(),
		})
	}))
	defer srv.Close()

	client := New("token", WithBaseURL(srv.URL), WithHTTPClient(srv.Client()))
	run, err := client.GetRuns(context.Background(), "owner/repo", "100")
	assertNoError(t, err)
	assertEqual(t, "100", run.ID)
	assertEqual(t, models.RunStatusRunning, run.Status)
	assertEqual(t, models.TriggerPullRequest, run.Trigger)
}

func TestCancelRun(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/runs/100/cancel")
		assertMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusAccepted)
	}))
	defer srv.Close()

	client := New("token", WithBaseURL(srv.URL), WithHTTPClient(srv.Client()))
	err := client.CancelRun(context.Background(), "owner/repo", "100")
	assertNoError(t, err)
}

func TestRetryRun(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/runs/100/rerun")
		assertMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
	}))
	defer srv.Close()

	client := New("token", WithBaseURL(srv.URL), WithHTTPClient(srv.Client()))
	err := client.RetryRun(context.Background(), "owner/repo", "100")
	assertNoError(t, err)
}

func TestListJobs(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/runs/100/jobs")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(listJobsResponse{
			TotalCount: 2,
			Jobs: []jobDTO{
				{ID: 1, RunID: 100, Name: "test", Status: "completed", Conclusion: "success"},
				{ID: 2, RunID: 100, Name: "build", Status: "completed", Conclusion: "failure"},
			},
		})
	}))
	defer srv.Close()

	client := New("token", WithBaseURL(srv.URL), WithHTTPClient(srv.Client()))
	jobs, err := client.ListJobs(context.Background(), "owner/repo", "100")
	assertNoError(t, err)
	assertEqual(t, 2, len(jobs))
	assertEqual(t, "test", jobs[0].Name)
	assertEqual(t, models.RunResultSuccess, jobs[0].Result)
	assertEqual(t, models.RunResultFailed, jobs[1].Result)
}

func TestListArtifacts(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/runs/100/artifacts")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(listArtifactsResponse{
			TotalCount: 1,
			Artifacts: []artifactDTO{
				{
					ID: 1, Name: "build.zip", SizeInBytes: 1024,
					CreatedAt: time.Now(), Expired: false,
					ArchiveDownloadURL: "https://api.github.com/repos/owner/repo/actions/artifacts/1/zip",
				},
			},
		})
	}))
	defer srv.Close()

	client := New("token", WithBaseURL(srv.URL), WithHTTPClient(srv.Client()))
	artifacts, err := client.ListArtifacts(context.Background(), "owner/repo", "100")
	assertNoError(t, err)
	assertEqual(t, 1, len(artifacts))
	assertEqual(t, "1", artifacts[0].ID)
	assertEqual(t, "build.zip", artifacts[0].Name)
	assertEqual(t, "100", artifacts[0].RunID)
}

func TestDownloadArtifact(t *testing.T) {
	t.Parallel()

	expected := []byte("zip-content")
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/artifacts/1/zip")
		assertMethod(t, r, http.MethodGet)
		w.Header().Set("Content-Type", "application/zip")
		w.Write(expected)
	}))
	defer srv.Close()

	client := New("token", WithBaseURL(srv.URL), WithHTTPClient(srv.Client()))
	data, err := client.DownloadArtifact(context.Background(), "owner/repo", "1")
	assertNoError(t, err)
	assertEqual(t, string(expected), string(data))
}

func TestGetLogs(t *testing.T) {
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

	client := New("token", WithBaseURL(srv.URL), WithHTTPClient(srv.Client()))
	log, err := client.GetLogs(context.Background(), "owner/repo", "100")
	assertNoError(t, err)
	assertEqual(t, "100", log.RunID)
	assertEqual(t, true, len(log.Content) > 0)
}

func TestAuthHeaders(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, "Bearer test-token", r.Header.Get("Authorization"))
		assertEqual(t, "application/vnd.github+json", r.Header.Get("Accept"))
		assertEqual(t, "2022-11-28", r.Header.Get("X-GitHub-Api-Version"))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(listWorkflowsResponse{})
	}))
	defer srv.Close()

	client := New("test-token", WithBaseURL(srv.URL), WithHTTPClient(srv.Client()))
	_, err := client.ListPipelines(context.Background(), "owner/repo")
	assertNoError(t, err)
}

func TestNotFoundError(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(apiError{Message: "Not Found"})
	}))
	defer srv.Close()

	client := New("token", WithBaseURL(srv.URL), WithHTTPClient(srv.Client()))
	_, err := client.GetPipeline(context.Background(), "owner/repo", "999")
	assertError(t, err)
}

func TestUnauthorizedError(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(apiError{Message: "Bad credentials"})
	}))
	defer srv.Close()

	client := New("bad-token", WithBaseURL(srv.URL), WithHTTPClient(srv.Client()))
	_, err := client.ListRuns(context.Background(), "owner/repo", "1")
	assertError(t, err)
}

func TestMapperRunStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  models.RunStatus
	}{
		{"queued", models.RunStatusQueued},
		{"in_progress", models.RunStatusRunning},
		{"completed", models.RunStatusCompleted},
		{"waiting", models.RunStatusWaiting},
		{"pending", models.RunStatusPending},
		{"cancelled", models.RunStatusCancelled},
		{"unknown_status", models.RunStatusUnknown},
	}

	for _, tc := range tests {
		got := mapRunStatus(tc.input)
		if got != tc.want {
			t.Errorf("mapRunStatus(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

func TestMapperRunResult(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  models.RunResult
	}{
		{"success", models.RunResultSuccess},
		{"failure", models.RunResultFailed},
		{"cancelled", models.RunResultCancelled},
		{"skipped", models.RunResultSkipped},
		{"timed_out", models.RunResultTimedOut},
		{"neutral", models.RunResultNeutral},
		{"action_required", models.RunResultActionReq},
		{"unknown_result", models.RunResultUnknown},
	}

	for _, tc := range tests {
		got := mapRunResult(tc.input)
		if got != tc.want {
			t.Errorf("mapRunResult(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

func TestMapperTriggerType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  models.TriggerType
	}{
		{"push", models.TriggerPush},
		{"pull_request", models.TriggerPullRequest},
		{"release", models.TriggerRelease},
		{"schedule", models.TriggerSchedule},
		{"workflow_dispatch", models.TriggerManual},
		{"repository_dispatch", models.TriggerWebhook},
		{"workflow_call", models.TriggerWorkflowCall},
		{"unknown_event", models.TriggerUnknown},
	}

	for _, tc := range tests {
		got := mapTriggerType(tc.input)
		if got != tc.want {
			t.Errorf("mapTriggerType(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

func TestParseLinkHeader(t *testing.T) {
	t.Parallel()

	header := `<https://api.github.com/repos/owner/repo/actions/workflows?page=2>; rel="next", <https://api.github.com/repos/owner/repo/actions/workflows?page=5>; rel="last"`
	links := parseLinkHeader(header)

	assertEqual(t, "https://api.github.com/repos/owner/repo/actions/workflows?page=2", links.Next())
	assertEqual(t, "https://api.github.com/repos/owner/repo/actions/workflows?page=5", links.Last())
}

func TestParseLinkHeaderEmpty(t *testing.T) {
	t.Parallel()

	links := parseLinkHeader("")
	assertEqual(t, "", links.Next())
}

func TestNewDefaultBaseURL(t *testing.T) {
	t.Parallel()

	client := New("token")
	assertEqual(t, defaultBaseURL, client.baseURL)
}

func TestWithHTTPClient(t *testing.T) {
	t.Parallel()

	client := New("token", WithHTTPClient(&http.Client{Timeout: 5 * time.Second}))
	assertEqual(t, 5*time.Second, client.httpClient.Timeout)
}

func TestWithBaseURL(t *testing.T) {
	t.Parallel()

	client := New("token", WithBaseURL("https://custom.url"))
	assertEqual(t, "https://custom.url", client.baseURL)
}

// helpers

func assertPath(t *testing.T, r *http.Request, expected string) {
	t.Helper()
	if r.URL.Path != expected {
		t.Errorf("path = %q, want %q", r.URL.Path, expected)
	}
}

func assertMethod(t *testing.T, r *http.Request, expected string) {
	t.Helper()
	if r.Method != expected {
		t.Errorf("method = %q, want %q", r.Method, expected)
	}
}

func assertAuth(t *testing.T, r *http.Request) {
	t.Helper()
	if r.Header.Get("Authorization") == "" {
		t.Error("missing Authorization header")
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func createTestZip(files map[string]string) []byte {
	var buf bytes.Buffer
	w := zip.NewWriter(&buf)
	for name, content := range files {
		f, err := w.Create(name)
		if err != nil {
			panic(err)
		}
		f.Write([]byte(content))
	}
	w.Close()
	return buf.Bytes()
}
