package circleci_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oktanthq/ciforge/provider/circleci"
)

func TestListWebhooks(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/webhook")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)
		assertQuery(t, r, "scope-id", "proj-1")
		assertQuery(t, r, "scope-type", "project")

		json.NewEncoder(w).Encode(map[string]any{
			"items": []map[string]any{
				{"id": "wh-1", "name": "notify", "url": "https://example.com/hook", "events": []string{"workflow-completed"}, "created_at": "2019-08-24T14:15:22Z", "updated_at": "2019-08-24T14:15:22Z"},
			},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	hooks, err := client.ListWebhooks(context.Background(), circleci.ListWebhooksOptions{
		ScopeID:   "proj-1",
		ScopeType: "project",
	})
	assertNoError(t, err)
	assertEqual(t, 1, len(hooks))
	assertEqual(t, "notify", hooks[0].Name)
}

func TestCreateWebhooks(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/webhook")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{
			"id": "wh-1", "name": "notify", "url": "https://example.com/hook", "events": []string{"job-completed"},
			"created_at": "2019-08-24T14:15:22Z", "updated_at": "2019-08-24T14:15:22Z",
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	hook, err := client.CreateWebhooks(context.Background(), circleci.CreateWebhooksBody{
		Name:   "notify",
		Events: []string{"job-completed"},
		URL:    "https://example.com/hook",
		Scope:  circleci.Scope{ID: "proj-1", Type: "project"},
	})
	assertNoError(t, err)
	assertEqual(t, "wh-1", hook.ID)
}

func TestGetWebhook(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/webhook/wh-1")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"id": "wh-1", "name": "notify", "url": "https://example.com/hook",
			"created_at": "2019-08-24T14:15:22Z", "updated_at": "2019-08-24T14:15:22Z",
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	hook, err := client.GetWebhook(context.Background(), "wh-1")
	assertNoError(t, err)
	assertEqual(t, "notify", hook.Name)
}

func TestUpdateWebhook(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/webhook/wh-1")
		assertMethod(t, r, http.MethodPut)

		json.NewEncoder(w).Encode(map[string]any{
			"id": "wh-1", "name": "updated", "url": "https://example.com/hook",
			"created_at": "2019-08-24T14:15:22Z", "updated_at": "2019-08-24T14:15:22Z",
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	hook, err := client.UpdateWebhook(context.Background(), "wh-1", circleci.UpdateWebhookBody{Name: "updated"})
	assertNoError(t, err)
	assertEqual(t, "updated", hook.Name)
}

func TestDeleteWebhooks(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/webhook/wh-1")
		assertMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.DeleteWebhooks(context.Background(), "wh-1")
	assertNoError(t, err)
}
