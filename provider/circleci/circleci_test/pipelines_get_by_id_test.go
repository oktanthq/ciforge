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
  "errors": [],
  "project_slug": "gh/CircleCI-Public/api-preview-docs",
  "updated_at": "2023-10-27T10:00:00Z",
  "number": 123,
  "state": "created",
  "created_at": "2023-10-27T09:59:00Z",
  "trigger": {
    "type": "scheduled_pipeline",
    "received_at": "2023-10-27T09:59:00Z",
    "actor": {
      "login": "octocat",
      "avatar_url": "https://avatars.githubusercontent.com/u/583231?v=4"
    }
  },
  "vcs": {
    "provider_name": "GitHub",
    "target_repository_url": "https://github.com/CircleCI-Public/api-preview-docs",
    "branch": "main",
    "review_id": "123",
    "review_url": "https://github.com/CircleCI-Public/api-preview-docs/pull/123",
    "revision": "f454a02b5d10fcccfd7d9dd7608a76d6493a98b4",
    "tag": "v3.1.4159",
    "commit": {
      "subject": "Update README",
      "body": "Detailed description of the update."
    },
    "origin_repository_url": "https://github.com/CircleCI-Public/api-preview-docs"
  }
}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	pipeline, err := client.GetPipelineByID(context.Background(), "5034460f-c7c4-4c43-9457-de07e2029e7b")

	assertNoError(t, err)
	assertEqual(t, pipeline.ID, "5034460f-c7c4-4c43-9457-de07e2029e7b")
	assertEqual(t, pipeline.State, "created")
	assertEqual(t, pipeline.ProjectSlug, "gh/CircleCI-Public/api-preview-docs")
	assertEqual(t, pipeline.Number, int64(123))
	assertEqual(t, pipeline.Trigger.Type, "scheduled_pipeline")
	assertEqual(t, pipeline.Trigger.Actor.Login, "octocat")
	assertEqual(t, pipeline.VCS.ProviderName, "GitHub")
	assertEqual(t, pipeline.VCS.Branch, "main")
	assertEqual(t, pipeline.VCS.Revision, "f454a02b5d10fcccfd7d9dd7608a76d6493a98b4")
	assertEqual(t, pipeline.VCS.Tag, "v3.1.4159")
	assertEqual(t, pipeline.VCS.ReviewID, "123")
	assertEqual(t, pipeline.VCS.ReviewURL, "https://github.com/CircleCI-Public/api-preview-docs/pull/123")
	assertEqual(t, pipeline.VCS.OriginRepositoryURL, "https://github.com/CircleCI-Public/api-preview-docs")
}

func TestGetPipelineByID_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{
  "message": "Entity not found."
}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	_, err := client.GetPipelineByID(context.Background(), "bad-id")

	assertError(t, err)
	if !errors.Is(err, circleci.ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}
