package gitlab_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oktanthq/ciforge/provider/gitlab"
)

func TestListTriggerTokens(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/triggers")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		json.NewEncoder(w).Encode([]map[string]any{
			{"id": 1, "description": "deploy", "token": "tok-aaaa", "created_at": "2026-01-02T15:04:05Z", "updated_at": "2026-01-02T15:04:05Z"},
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	tokens, err := client.ListTriggerTokens(context.Background(), 42)
	assertNoError(t, err)
	assertEqual(t, 1, len(tokens))
	assertEqual(t, "deploy", tokens[0].Description)
	assertEqual(t, "tok-aaaa", tokens[0].Token)
}

func TestGetTriggerToken(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/triggers/1")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"id": 1, "description": "deploy", "token": "tok-aaaa",
			"created_at": "2026-01-02T15:04:05Z", "updated_at": "2026-01-02T15:04:05Z",
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	token, err := client.GetTriggerToken(context.Background(), 42, 1)
	assertNoError(t, err)
	assertEqual(t, 1, token.ID)
	assertEqual(t, "deploy", token.Description)
}

func TestCreateTriggerToken(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/triggers")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{
			"id": 2, "description": "ci", "token": "tok-bbbb",
			"created_at": "2026-01-02T15:04:05Z", "updated_at": "2026-01-02T15:04:05Z",
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	token, err := client.CreateTriggerToken(context.Background(), 42, "ci")
	assertNoError(t, err)
	assertEqual(t, 2, token.ID)
	assertEqual(t, "ci", token.Description)
}

func TestUpdateTriggerToken(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/triggers/2")
		assertMethod(t, r, http.MethodPut)

		json.NewEncoder(w).Encode(map[string]any{
			"id": 2, "description": "ci-updated", "token": "tok-bbbb",
			"created_at": "2026-01-02T15:04:05Z", "updated_at": "2026-01-02T15:04:05Z",
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	token, err := client.UpdateTriggerToken(context.Background(), 42, 2, "ci-updated")
	assertNoError(t, err)
	assertEqual(t, "ci-updated", token.Description)
}

func TestDeleteTriggerToken(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/triggers/2")
		assertMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	err := client.DeleteTriggerToken(context.Background(), 42, 2)
	assertNoError(t, err)
}

func TestTriggerPipeline(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/trigger/pipeline")
		assertMethod(t, r, http.MethodPost)
		assertQuery(t, r, "ref", "main")
		assertQuery(t, r, "token", "tok-aaaa")

		json.NewEncoder(w).Encode(map[string]any{
			"id": 77, "status": "pending", "ref": "main",
			"created_at": "2026-01-02T15:04:05Z", "updated_at": "2026-01-02T15:04:05Z",
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	pipeline, err := client.TriggerPipeline(context.Background(), 42, &gitlab.TriggerPipelineBody{
		Ref:   "main",
		Token: "tok-aaaa",
		Variables: map[string]string{
			"ENV": "prod",
		},
	})
	assertNoError(t, err)
	assertEqual(t, 77, pipeline.ID)
	assertEqual(t, "pending", pipeline.Status)
}
