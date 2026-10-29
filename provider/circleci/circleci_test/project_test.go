package circleci_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oktanthq/ciforge/provider/circleci"
)

func TestCreateProject(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/organization/gh/org/project")
		assertMethod(t, r, http.MethodPost)
		assertAuth(t, r)

		json.NewEncoder(w).Encode(map[string]any{
			"id": "proj-1", "name": "api", "slug": "gh/org/api", "organization_name": "org",
			"vcs_info": map[string]any{"provider": "GitHub", "default_branch": "main", "vcs_url": "https://github.com/org/api"},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	project, err := client.CreateProject(context.Background(), "gh/org", circleci.CreateProjectBody{Name: "api"})
	assertNoError(t, err)
	assertEqual(t, "api", project.Name)
	assertEqual(t, "gh/org/api", project.Slug)
}

func TestGetProject(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/org/api")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"id": "proj-1", "name": "api", "slug": "gh/org/api", "organization_name": "org",
			"vcs_info": map[string]any{"provider": "GitHub", "default_branch": "main"},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	project, err := client.GetProject(context.Background(), "gh/org/api")
	assertNoError(t, err)
	assertEqual(t, "proj-1", project.ID)
	assertEqual(t, "GitHub", project.VCSInfo.Provider)
}

func TestDeleteProject(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/org/api")
		assertMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.DeleteProject(context.Background(), "gh/org/api")
	assertNoError(t, err)
}

func TestGetAllCheckoutKeys(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/org/api/checkout-key")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"items": []map[string]any{
				{"public-key": "ssh-rsa AAA", "type": "deploy-key", "fingerprint": "aa:bb", "preferred": true, "created-at": "2019-08-24T14:15:22Z"},
			},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	keys, err := client.GetAllCheckoutKeys(context.Background(), "gh/org/api")
	assertNoError(t, err)
	assertEqual(t, 1, len(keys))
	assertEqual(t, "aa:bb", keys[0].Fingerprint)
}

func TestCreateCheckoutKey(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/org/api/checkout-key")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{
			"public-key": "ssh-rsa AAA", "type": "deploy-key", "fingerprint": "aa:bb", "preferred": true, "created-at": "2019-08-24T14:15:22Z",
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	key, err := client.CreateCheckoutKey(context.Background(), "gh/org/api", circleci.CreateCheckoutKeyBody{Type: circleci.CheckoutKeyTypeDeployKey})
	assertNoError(t, err)
	assertEqual(t, "deploy-key", key.Type)
}

func TestGetCheckoutKey(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/org/api/checkout-key/aa:bb")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"public-key": "ssh-rsa AAA", "type": "deploy-key", "fingerprint": "aa:bb", "preferred": true, "created-at": "2019-08-24T14:15:22Z",
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	key, err := client.GetCheckoutKey(context.Background(), "gh/org/api", "aa:bb")
	assertNoError(t, err)
	assertEqual(t, "aa:bb", key.Fingerprint)
}

func TestDeleteCheckoutKey(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/org/api/checkout-key/aa:bb")
		assertMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.DeleteCheckoutKey(context.Background(), "gh/org/api", "aa:bb")
	assertNoError(t, err)
}

func TestListEvironmentVariables(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/org/api/envvar")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"items": []map[string]any{
				{"name": "TOKEN", "value": "xxxx", "created-at": "2019-08-24T14:15:22Z"},
			},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	vars, err := client.ListEvironmentVariables(context.Background(), "gh/org/api")
	assertNoError(t, err)
	assertEqual(t, 1, len(vars))
	assertEqual(t, "TOKEN", vars[0].Name)
}

func TestCreateEnvironmentVariables(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/org/api/envvar")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{"name": "TOKEN", "value": "xxxx"})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	v, err := client.CreateEnvironmentVariables(context.Background(), "gh/org/api", circleci.CreateProjectEnvVarBody{Name: "TOKEN", Value: "secret"})
	assertNoError(t, err)
	assertEqual(t, "TOKEN", v.Name)
}

func TestGetMaskedEnvironmentVariable(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/org/api/envvar/TOKEN")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{"name": "TOKEN", "value": "xxxx"})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	v, err := client.GetMaskedEnvironmentVariable(context.Background(), "gh/org/api", "TOKEN")
	assertNoError(t, err)
	assertEqual(t, "xxxx", v.Value)
}

func TestDeleteEnvironmentVariables(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/org/api/envvar/TOKEN")
		assertMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.DeleteEnvironmentVariables(context.Background(), "gh/org/api", "TOKEN")
	assertNoError(t, err)
}

func TestGetProjectSettings(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/org/api/settings")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"advanced": map[string]any{"autocancel_builds": true, "oss": false, "setup_workflows": true},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	settings, err := client.GetProjectSettings(context.Background(), "gh", "org", "api")
	assertNoError(t, err)
	assertEqual(t, true, settings.AutocancelBuilds)
	assertEqual(t, true, settings.SetupWorkflows)
}

func TestUpdateProjectSettings(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/org/api/settings")
		assertMethod(t, r, http.MethodPatch)

		json.NewEncoder(w).Encode(map[string]any{
			"advanced": map[string]any{"autocancel_builds": false, "oss": true},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	settings, err := client.UpdateProjectSettings(context.Background(), "gh", "org", "api", circleci.UpdateProjectSettingsBody{
		Advanced: circleci.ProjectSettings{OSS: true},
	})
	assertNoError(t, err)
	assertEqual(t, true, settings.OSS)
}
