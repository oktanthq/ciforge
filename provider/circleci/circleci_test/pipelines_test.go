package circleci_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oktanthq/ciforge/provider/circleci"
)

func TestListPipelines(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/pipeline")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)
		assertQuery(t, r, "org-slug", "gh/CircleCI-Public")
		assertQuery(t, r, "mine", "true")

		json.NewEncoder(w).Encode(map[string]any{
			"items": []map[string]any{
				{"id": "p1", "project_slug": "gh/CircleCI-Public/api-preview-docs", "number": 25, "state": "created", "created_at": "2019-08-24T14:15:22Z", "updated_at": "2019-08-24T14:15:22Z"},
			},
			"next_page_token": "next",
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	pipelines, err := client.ListPipelines(context.Background(), &circleci.ListPipelinesOptions{
		OrgSlug: "gh/CircleCI-Public",
		Mine:    true,
	})
	assertNoError(t, err)
	assertEqual(t, 1, len(pipelines))
	assertEqual(t, "p1", pipelines[0].ID)
	assertEqual(t, int64(25), pipelines[0].Number)
}

func TestContinuePipeline(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/pipeline/continue")
		assertMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.ContinuePipeline(context.Background(), &circleci.ContinuePipelineBody{
		ContinuationKey: "key",
		Configuration:   "version: 2.1",
	})
	assertNoError(t, err)
}

func TestGetPipelineByID(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/pipeline/p1")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		json.NewEncoder(w).Encode(map[string]any{
			"id": "p1", "project_slug": "gh/org/proj", "number": 25, "state": "created",
			"created_at": "2019-08-24T14:15:22Z", "updated_at": "2019-08-24T14:15:22Z",
			"vcs": map[string]any{"provider_name": "GitHub", "branch": "main", "revision": "abc"},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	pipeline, err := client.GetPipelineByID(context.Background(), "p1")
	assertNoError(t, err)
	assertEqual(t, "p1", pipeline.ID)
	assertEqual(t, "main", pipeline.VCS.Branch)
}

func TestGetPipelineConfig(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/pipeline/p1/config")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"source": "version: 2.1", "compiled": "version: 2.1",
			"setup-config": "setup", "compiled-setup-config": "compiled-setup",
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	cfg, err := client.GetPipelineConfig(context.Background(), "p1")
	assertNoError(t, err)
	assertEqual(t, "version: 2.1", cfg.Source)
	assertEqual(t, "setup", cfg.SetupConfig)
}

func TestListPipelineWorkflows(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/pipeline/p1/workflow")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"items": []map[string]any{
				{"id": "w1", "name": "build", "status": "success", "pipeline_id": "p1", "created_at": "2019-08-24T14:15:22Z"},
			},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	workflows, err := client.ListPipelineWorkflows(context.Background(), "p1")
	assertNoError(t, err)
	assertEqual(t, 1, len(workflows))
	assertEqual(t, "build", workflows[0].Name)
}

func TestTriggerPipeline(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/org/proj/pipeline")
		assertMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{
			"id": "p2", "number": 26, "state": "pending", "created_at": "2019-08-24T14:15:22Z",
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.TriggerPipeline(context.Background(), "gh/org/proj", circleci.TriggerPipelineBody{Branch: "main"})
	assertNoError(t, err)
}

func TestGetAllPipelines(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/org/proj/pipeline")
		assertMethod(t, r, http.MethodGet)
		assertQuery(t, r, "branch", "main")

		json.NewEncoder(w).Encode(map[string]any{
			"items": []map[string]any{
				{"id": "p1", "number": 25, "state": "created", "created_at": "2019-08-24T14:15:22Z", "updated_at": "2019-08-24T14:15:22Z"},
			},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	pipelines, err := client.GetAllPipelines(context.Background(), "gh/org/proj", circleci.GetPipelineOptions{Branch: "main"})
	assertNoError(t, err)
	assertEqual(t, 1, len(pipelines))
}

func TestGetYourPipelines(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/org/proj/pipeline/mine")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"items": []map[string]any{
				{"id": "p1", "number": 25, "state": "created", "created_at": "2019-08-24T14:15:22Z", "updated_at": "2019-08-24T14:15:22Z"},
			},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	pipelines, err := client.GetYourPipelines(context.Background(), "gh/org/proj", circleci.GetYourPipelineOptions{})
	assertNoError(t, err)
	assertEqual(t, 1, len(pipelines))
}

func TestGetPipelineByNumber(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/org/proj/pipeline/25")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"id": "p1", "number": 25, "state": "created", "created_at": "2019-08-24T14:15:22Z", "updated_at": "2019-08-24T14:15:22Z",
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	pipeline, err := client.GetPipelineByNumber(context.Background(), "gh/org/proj", 25)
	assertNoError(t, err)
	assertEqual(t, int64(25), pipeline.Number)
}

func TestTriggerPipelineProject(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/org/proj/pipeline/run")
		assertMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusAccepted)
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.TriggerPipelineProject(context.Background(), "gh", "org", "proj", circleci.TriggerProjPipelineBody{
		DefinitionID: "def-1",
		Config:       circleci.Config{Branch: "main"},
		Checkout:     circleci.Checkout{Branch: "main"},
	})
	assertNoError(t, err)
}

func TestGetPipelineValues(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/pipeline/p1/values")
		assertMethod(t, r, http.MethodGet)
		json.NewEncoder(w).Encode(map[string]any{"pipeline.number": 25})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.GetPipelineValues(context.Background(), "p1")
	assertNoError(t, err)
}
