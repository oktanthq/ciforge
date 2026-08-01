package circleci_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abdulahshoaib/ciforge/provider/circleci"
)

func TestGetPipelineByID_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/pipeline/5034460f-c7c4-4c43-9457-de07e2029e7b")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
  "id": "5034460f-c7c4-4c43-9457-de07e2029e7b",
  "errors": [
    {
      "type": "config",
      "message": "string"
    }
  ],
  "project_slug": "gh/CircleCI-Public/api-preview-docs",
  "updated_at": "2019-08-24T14:15:22Z",
  "number": 25,
  "trigger_parameters": {
    "property1": "string",
    "property2": "string"
  },
  "state": "created",
  "created_at": "2019-08-24T14:15:22Z",
  "warnings": [
    {
      "type": "config-deprecated-syntax",
      "message": "string"
    }
  ],
  "trigger": {
    "type": "scheduled_pipeline",
    "received_at": "2019-08-24T14:15:22Z",
    "actor": {
      "login": "string",
      "avatar_url": "string"
    }
  },
  "vcs": {
    "provider_name": "GitHub",
    "target_repository_url": "https://github.com/CircleCI-Public/api-preview-docs",
    "branch": "feature/design-new-api",
    "review_id": "123",
    "review_url": "https://github.com/CircleCI-Public/api-preview-docs/pull/123",
    "revision": "f454a02b5d10fcccfd7d9dd7608a76d6493a98b4",
    "tag": "v3.1.4159",
    "commit": {
      "subject": "string",
      "body": "string"
    },
    "origin_repository_url": "https://github.com/CircleCI-Public/api-preview-docs"
  }
}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	pipeline, err := client.GetPipelineByID(context.Background(), "5034460f-c7c4-4c43-9457-de07e2029e7b")

	assertNoError(t, err)
	assertEqual(t, "5034460f-c7c4-4c43-9457-de07e2029e7b", pipeline.ID)
	assertEqual(t, "gh/CircleCI-Public/api-preview-docs", pipeline.ProjectSlug)
	assertEqual(t, int64(25), pipeline.Number)
	assertEqual(t, "created", pipeline.State)

	if len(pipeline.Errors) != 1 {
		t.Fatalf("errors len = %d, want 1", len(pipeline.Errors))
	}
	assertEqual(t, "config", pipeline.Errors[0].Type)

	if len(pipeline.Warnings) != 1 {
		t.Fatalf("warnings len = %d, want 1", len(pipeline.Warnings))
	}
	assertEqual(t, "config-deprecated-syntax", pipeline.Warnings[0].Type)

	assertEqual(t, "scheduled_pipeline", pipeline.Trigger.Type)
	assertEqual(t, "string", pipeline.Trigger.Actor.Login)

	assertEqual(t, "GitHub", pipeline.VCS.ProviderName)
	assertEqual(t, "feature/design-new-api", pipeline.VCS.Branch)
	assertEqual(t, "v3.1.4159", pipeline.VCS.Tag)
	assertEqual(t, "f454a02b5d10fcccfd7d9dd7608a76d6493a98b4", pipeline.VCS.Revision)
}

func TestGetPipelineByID_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Entity not found."}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	_, err := client.GetPipelineByID(context.Background(), "bad-id")

	assertError(t, err)
	if !errors.Is(err, circleci.ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}

func TestListPipeline_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/pipeline")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
  "items": [
    {
      "id": "5034460f-c7c4-4c43-9457-de07e2029e7b",
      "errors": [
        {
          "type": "config",
          "message": "string"
        }
      ],
      "project_slug": "gh/CircleCI-Public/api-preview-docs",
      "updated_at": "2019-08-24T14:15:22Z",
      "number": 25,
      "trigger_parameters": {
        "property1": "string",
        "property2": "string"
      },
      "state": "created",
      "created_at": "2019-08-24T14:15:22Z",
      "warnings": [
        {
          "type": "config-deprecated-syntax",
          "message": "string"
        }
      ],
      "trigger": {
        "type": "scheduled_pipeline",
        "received_at": "2019-08-24T14:15:22Z",
        "actor": {
          "login": "string",
          "avatar_url": "string"
        }
      },
      "vcs": {
        "provider_name": "GitHub",
        "target_repository_url": "https://github.com/CircleCI-Public/api-preview-docs",
        "branch": "feature/design-new-api",
        "review_id": "123",
        "review_url": "https://github.com/CircleCI-Public/api-preview-docs/pull/123",
        "revision": "f454a02b5d10fcccfd7d9dd7608a76d6493a98b4",
        "tag": "v3.1.4159",
        "commit": {
          "subject": "string",
          "body": "string"
        },
        "origin_repository_url": "https://github.com/CircleCI-Public/api-preview-docs"
      }
    }
  ],
  "next_page_token": "string"
}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	pipelines, err := client.ListPipelines(context.Background(), &circleci.ListPipelinesOptions{
		OrgSlug:   "gh%2FCircleCI-Public",
		PageToken: "",
		Mine:      true,
	})

	assertNoError(t, err)
	assertEqual(t, 1, len(pipelines))
	if len(pipelines) > 0 {
		assertEqual(t, "5034460f-c7c4-4c43-9457-de07e2029e7b", pipelines[0].ID)
		assertEqual(t, "gh/CircleCI-Public/api-preview-docs", pipelines[0].ProjectSlug)
		assertEqual(t, int64(25), pipelines[0].Number)
		assertEqual(t, "created", pipelines[0].State)
	}
}

func TestListPipeline_Failed(t *testing.T) {

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/pipeline")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Internal Server Error"}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	_, err := client.ListPipelines(context.Background(), &circleci.ListPipelinesOptions{
		OrgSlug:   "gh%2FCircleCI-Public",
		PageToken: "",
		Mine:      true,
	})

	assertError(t, err)
}

func TestContinuePipeline_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/pipeline/continue")
		assertMethod(t, r, http.MethodPost)
		assertAuth(t, r)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{ "message": "Pipeline continued successfully." }`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.ContinuePipeline(context.Background(), &circleci.ContinuePipelineBody{
		ContinuationKey: "5034460f-c7c4-4c43-9457-de07e2029e7b",
		Configuration:   "version: 2.1",
	})

	assertNoError(t, err)
}

func TestContinuePipeline_Failed(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/pipeline/continue")
		assertMethod(t, r, http.MethodPost)
		assertAuth(t, r)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Internal Server Error"}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.ContinuePipeline(context.Background(), &circleci.ContinuePipelineBody{
		ContinuationKey: "5034460f-c7c4-4c43-9457-de07e2029e7b",
		Configuration:   "version: 2.1",
	})

	assertError(t, err)
}

func TestGetPipelineConfig_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/pipeline/5034460f-c7c4-4c43-9457-de07e2029e7b/config")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
  "source": "version: 2.1",
  "compiled": "version: 2.1",
  "setup-config": "version: 2.1",
  "compiled-setup-config": "version: 2.1"
}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	cfg, err := client.GetPipelineConfig(context.Background(), "5034460f-c7c4-4c43-9457-de07e2029e7b")

	assertNoError(t, err)
	assertEqual(t, "version: 2.1", cfg.Source)
	assertEqual(t, "version: 2.1", cfg.Compiled)
	assertEqual(t, "version: 2.1", cfg.SetupConfig)
	assertEqual(t, "version: 2.1", cfg.CompiledSetupConfig)
}

func TestGetPipelineConfig_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Entity not found."}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	_, err := client.GetPipelineConfig(context.Background(), "bad-id")

	assertError(t, err)
	if !errors.Is(err, circleci.ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}

func TestListPipelineWorkflows_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/pipeline/5034460f-c7c4-4c43-9457-de07e2029e7b/workflow")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
  "items": [
    {
      "pipeline_id": "5034460f-c7c4-4c43-9457-de07e2029e7b",
      "id": "de6a4c5f-8b1a-4e3d-9c2f-1a2b3c4d5e6f",
      "name": "build-test-deploy",
      "project_slug": "gh/CircleCI-Public/api-preview-docs",
      "status": "running",
      "pipeline_number": 25,
      "created_at": "2019-08-24T14:15:22Z"
    }
  ],
  "next_page_token": "string"
}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	workflows, err := client.ListPipelineWorkflows(context.Background(), "5034460f-c7c4-4c43-9457-de07e2029e7b")

	assertNoError(t, err)
	assertEqual(t, 1, len(workflows))
	if len(workflows) > 0 {
		assertEqual(t, "de6a4c5f-8b1a-4e3d-9c2f-1a2b3c4d5e6f", workflows[0].ID)
		assertEqual(t, "build-test-deploy", workflows[0].Name)
		assertEqual(t, "running", workflows[0].Status)
	}
}

func TestListPipelineWorkflows_Failed(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/pipeline/5034460f-c7c4-4c43-9457-de07e2029e7b/workflow")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Internal Server Error"}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	_, err := client.ListPipelineWorkflows(context.Background(), "5034460f-c7c4-4c43-9457-de07e2029e7b")

	assertError(t, err)
}

func TestTriggerPipeline_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/api-preview-docs/pipeline")
		assertMethod(t, r, http.MethodPost)
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{
  "id": "5034460f-c7c4-4c43-9457-de07e2029e7b",
  "number": 26,
  "state": "pending",
  "created_at": "2019-08-24T14:15:22Z"
}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.TriggerPipeline(context.Background(), "gh/CircleCI-Public/api-preview-docs", circleci.TriggerPipelineBody{
		Branch:     "main",
		Parameters: map[string]any{"deploy": true},
	})

	assertNoError(t, err)
}

func TestTriggerPipeline_Failed(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/api-preview-docs/pipeline")
		assertMethod(t, r, http.MethodPost)
		assertAuth(t, r)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Internal Server Error"}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.TriggerPipeline(context.Background(), "gh/CircleCI-Public/api-preview-docs", circleci.TriggerPipelineBody{
		Branch: "main",
	})

	assertError(t, err)
}

func TestGetAllPipelines_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/api-preview-docs/pipeline")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
  "items": [
    {
      "id": "5034460f-c7c4-4c43-9457-de07e2029e7b",
      "project_slug": "gh/CircleCI-Public/api-preview-docs",
      "number": 25,
      "state": "created",
      "created_at": "2019-08-24T14:15:22Z"
    }
  ],
  "next_page_token": "string"
}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	pipelines, err := client.GetAllPipelines(context.Background(), "gh/CircleCI-Public/api-preview-docs", circleci.GetPipelineOptions{
		Branch: "main",
	})

	assertNoError(t, err)
	assertEqual(t, 1, len(pipelines))
	if len(pipelines) > 0 {
		assertEqual(t, "5034460f-c7c4-4c43-9457-de07e2029e7b", pipelines[0].ID)
		assertEqual(t, "created", pipelines[0].State)
	}
}

func TestGetAllPipelines_Failed(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/api-preview-docs/pipeline")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Internal Server Error"}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	_, err := client.GetAllPipelines(context.Background(), "gh/CircleCI-Public/api-preview-docs", circleci.GetPipelineOptions{})

	assertError(t, err)
}

func TestGetYourPipelines_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/api-preview-docs/pipeline/mine")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
  "items": [
    {
      "id": "5034460f-c7c4-4c43-9457-de07e2029e7b",
      "project_slug": "gh/CircleCI-Public/api-preview-docs",
      "number": 25,
      "state": "created",
      "created_at": "2019-08-24T14:15:22Z"
    }
  ],
  "next_page_token": "string"
}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	pipelines, err := client.GetYourPipelines(context.Background(), "gh/CircleCI-Public/api-preview-docs", circleci.GetYourPipelineOptions{})

	assertNoError(t, err)
	assertEqual(t, 1, len(pipelines))
	if len(pipelines) > 0 {
		assertEqual(t, "5034460f-c7c4-4c43-9457-de07e2029e7b", pipelines[0].ID)
		assertEqual(t, int64(25), pipelines[0].Number)
	}
}

func TestGetYourPipelines_Failed(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/api-preview-docs/pipeline/mine")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Internal Server Error"}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	_, err := client.GetYourPipelines(context.Background(), "gh/CircleCI-Public/api-preview-docs", circleci.GetYourPipelineOptions{})

	assertError(t, err)
}

func TestGetPipelineByNumber_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/api-preview-docs/pipeline/25")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
  "id": "5034460f-c7c4-4c43-9457-de07e2029e7b",
  "project_slug": "gh/CircleCI-Public/api-preview-docs",
  "number": 25,
  "state": "created",
  "created_at": "2019-08-24T14:15:22Z"
}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	pipeline, err := client.GetPipelineByNumber(context.Background(), "gh/CircleCI-Public/api-preview-docs", 25)

	assertNoError(t, err)
	assertEqual(t, "5034460f-c7c4-4c43-9457-de07e2029e7b", pipeline.ID)
	assertEqual(t, int64(25), pipeline.Number)
	assertEqual(t, "created", pipeline.State)
}

func TestGetPipelineByNumber_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Entity not found."}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	_, err := client.GetPipelineByNumber(context.Background(), "gh/CircleCI-Public/api-preview-docs", 999)

	assertError(t, err)
	if !errors.Is(err, circleci.ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}

func TestTriggerPipelineProject_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/api-preview-docs/pipeline/run")
		assertMethod(t, r, http.MethodPost)
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.TriggerPipelineProject(context.Background(), "gh", "CircleCI-Public", "api-preview-docs", circleci.TriggerProjPipelineBody{
		DefinitionID: "def-123",
		Config:       circleci.Config{Branch: "main"},
		Checkout:     circleci.Checkout{Branch: "main"},
		Parameters:   map[string]any{"deploy": true},
	})

	assertNoError(t, err)
}

func TestTriggerPipelineProject_Failed(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/api-preview-docs/pipeline/run")
		assertMethod(t, r, http.MethodPost)
		assertAuth(t, r)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Internal Server Error"}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.TriggerPipelineProject(context.Background(), "gh", "CircleCI-Public", "api-preview-docs", circleci.TriggerProjPipelineBody{
		DefinitionID: "def-123",
	})

	assertError(t, err)
}

func TestGetPipelineValues_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/pipeline/5034460f-c7c4-4c43-9457-de07e2029e7b/values")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{ "key1": "value1", "key2": true }`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.GetPipelineValues(context.Background(), "5034460f-c7c4-4c43-9457-de07e2029e7b")

	assertNoError(t, err)
}

func TestGetPipelineValues_Failed(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/pipeline/5034460f-c7c4-4c43-9457-de07e2029e7b/values")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Internal Server Error"}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.GetPipelineValues(context.Background(), "5034460f-c7c4-4c43-9457-de07e2029e7b")

	assertError(t, err)
}
