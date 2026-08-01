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
