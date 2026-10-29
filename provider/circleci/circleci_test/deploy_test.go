package circleci_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oktanthq/ciforge/provider/circleci"
)

func TestListEnvironments(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/deploy/environments")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)
		assertQuery(t, r, "org-id", "org-1")
		assertQuery(t, r, "page-size", "20")

		json.NewEncoder(w).Encode(map[string]any{
			"items": []map[string]any{
				{"id": "e1", "name": "production", "description": "prod", "created_at": "2019-08-24T14:15:22Z", "updated_at": "2019-08-24T14:15:22Z"},
			},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	envs, err := client.ListEnvironments(context.Background(), circleci.ListEnvironmentsOptions{
		OrgID:    "org-1",
		PageSize: 20,
	})
	assertNoError(t, err)
	assertEqual(t, 1, len(envs))
	assertEqual(t, "production", envs[0].Name)
}

func TestGetEnvironment(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/deploy/environments/e1")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"id": "e1", "name": "production", "created_at": "2019-08-24T14:15:22Z", "updated_at": "2019-08-24T14:15:22Z",
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	env, err := client.GetEnvironment(context.Background(), "e1")
	assertNoError(t, err)
	assertEqual(t, "e1", env.ID)
}

func TestListComponents(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/deploy/component")
		assertMethod(t, r, http.MethodGet)
		assertQuery(t, r, "org-id", "org-1")
		assertQuery(t, r, "project-id", "proj-1")

		json.NewEncoder(w).Encode(map[string]any{
			"items": []map[string]any{
				{"id": "comp-1", "name": "api", "project_id": "proj-1", "release_count": 3, "created_at": "2019-08-24T14:15:22Z", "updated_at": "2019-08-24T14:15:22Z"},
			},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	components, err := client.ListComponents(context.Background(), circleci.ListComponentsOptions{
		OrgID:     "org-1",
		ProjectID: "proj-1",
	})
	assertNoError(t, err)
	assertEqual(t, 1, len(components))
	assertEqual(t, "api", components[0].Name)
}

func TestGetComponents(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/deploy/component/comp-1")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"id": "comp-1", "name": "api", "project_id": "proj-1", "created_at": "2019-08-24T14:15:22Z", "updated_at": "2019-08-24T14:15:22Z",
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	component, err := client.GetComponents(context.Background(), "comp-1")
	assertNoError(t, err)
	assertEqual(t, "api", component.Name)
}

func TestListComponentVersions(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/deploy/components/comp-1/versions")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"items": []map[string]any{
				{"name": "1.2.3", "environment_id": "e1", "is_live": true, "last_deployed_at": "2019-08-24T14:15:22Z"},
			},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	versions, err := client.ListComponentVersions(context.Background(), "comp-1")
	assertNoError(t, err)
	assertEqual(t, 1, len(versions))
	assertEqual(t, "1.2.3", versions[0].Name)
	assertEqual(t, true, versions[0].IsLive)
}

func TestRollbackProject(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/proj-1/rollback")
		assertMethod(t, r, http.MethodPost)
		assertAuth(t, r)

		json.NewEncoder(w).Encode(map[string]any{"id": "rb-1", "rollback_type": "pipeline"})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	resp, err := client.RollbackProject(context.Background(), "proj-1", circleci.RollbackBody{
		ComponentName:   "api",
		CurrentVersion:  "1.2.3",
		TargetVersion:   "1.2.2",
		EnvironmentName: "production",
		Reason:          "regression",
	})
	assertNoError(t, err)
	assertEqual(t, "rb-1", resp.ID)
	assertEqual(t, "pipeline", resp.RollbackType)
}
