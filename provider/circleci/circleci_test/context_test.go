package circleci_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oktanthq/ciforge/provider/circleci"
)

func TestCreateContext(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/context")
		assertMethod(t, r, http.MethodPost)
		assertAuth(t, r)

		json.NewEncoder(w).Encode(map[string]any{
			"id": "c1", "name": "prod", "created_at": "2019-08-24T14:15:22Z",
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	ctx, err := client.CreateContext(context.Background(), circleci.ContextBody{
		Name:  "prod",
		Owner: &circleci.ContextOwner{ID: "org-1", Type: "organization"},
	})
	assertNoError(t, err)
	assertEqual(t, "c1", ctx.ID)
	assertEqual(t, "prod", ctx.Name)
}

func TestListContexts(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/context")
		assertMethod(t, r, http.MethodGet)
		assertQuery(t, r, "owner-id", "org-1")
		assertQuery(t, r, "owner-type", "organization")

		json.NewEncoder(w).Encode(map[string]any{
			"items": []map[string]any{
				{"id": "c1", "name": "prod", "created_at": "2019-08-24T14:15:22Z"},
			},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	contexts, err := client.ListContexts(context.Background(), circleci.ListContextOptions{
		OwnerID:   "org-1",
		OwnerType: "organization",
	})
	assertNoError(t, err)
	assertEqual(t, 1, len(contexts))
	assertEqual(t, "prod", contexts[0].Name)
}

func TestGetContext(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/context/c1")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{"id": "c1", "name": "prod", "created_at": "2019-08-24T14:15:22Z"})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	ctx, err := client.GetContext(context.Background(), "c1")
	assertNoError(t, err)
	assertEqual(t, "c1", ctx.ID)
}

func TestDeleteContext(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/context/c1")
		assertMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.DeleteContext(context.Background(), "c1")
	assertNoError(t, err)
}

func TestListEnvironmentVariables(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/context/c1/environment-variable")
		assertMethod(t, r, http.MethodGet)
		assertQuery(t, r, "page-token", "tok")

		json.NewEncoder(w).Encode(map[string]any{
			"items": []map[string]any{
				{"variable": "API_KEY", "context_id": "c1", "created_at": "2019-08-24T14:15:22Z", "updated_at": "2019-08-24T14:15:22Z"},
			},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	vars, err := client.ListEnvironmentVariables(context.Background(), "c1", circleci.ListEnvVarOptions{PageToken: "tok"})
	assertNoError(t, err)
	assertEqual(t, 1, len(vars))
	assertEqual(t, "API_KEY", vars[0].Variable)
}

func TestUpsertEnviornmentVariable(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/context/c1/environment-variable/API_KEY")
		assertMethod(t, r, http.MethodPut)

		json.NewEncoder(w).Encode(map[string]any{
			"variable": "API_KEY", "context_id": "c1", "created_at": "2019-08-24T14:15:22Z", "updated_at": "2019-08-24T14:15:22Z",
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	v, err := client.UpsertEnviornmentVariable(context.Background(), "c1", "API_KEY", circleci.UpsertEnvVarBody{Value: "secret"})
	assertNoError(t, err)
	assertEqual(t, "API_KEY", v.Variable)
}

func TestRemoveEnviornmentVariable(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/context/c1/environment-variable/API_KEY")
		assertMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.RemoveEnviornmentVariable(context.Background(), "c1", "API_KEY")
	assertNoError(t, err)
}

func TestGetContextRestrictions(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/context/c1/restrictions")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"items": []map[string]any{
				{"id": "r1", "context_id": "c1", "restriction_type": "project", "restriction_value": "proj-1"},
			},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	restrictions, err := client.GetContextRestrictions(context.Background(), "c1")
	assertNoError(t, err)
	assertEqual(t, 1, len(restrictions))
	assertEqual(t, "project", restrictions[0].RestrictionType)
}

func TestCreateContextRestrictions(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/context/c1/restrictions")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{
			"id": "r1", "context_id": "c1", "restriction_type": "project", "restriction_value": "proj-1",
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	restriction, err := client.CreateContextRestrictions(context.Background(), "c1", circleci.CreateContextRestrictionsBody{
		RestrictionType:  "project",
		RestrictionValue: "proj-1",
	})
	assertNoError(t, err)
	assertEqual(t, "r1", restriction.ID)
}

func TestDeleteContextRestrictions(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/context/c1/restrictions/r1")
		assertMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.DeleteContextRestrictions(context.Background(), "c1", "r1")
	assertNoError(t, err)
}
