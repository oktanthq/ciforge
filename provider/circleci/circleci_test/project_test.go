package circleci_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abdulahshoaib/ciforge/provider/circleci"
)

func TestCreateProject_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/organization/gh/CircleCI-Public/project")
		assertMethod(t, r, http.MethodPost)
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{
  "slug": "gh/CircleCI-Public/test-project",
  "name": "test-project",
  "id": "12345",
  "organization_name": "CircleCI-Public",
  "organization_slug": "gh/CircleCI-Public",
  "organization_id": "1111",
  "vcs_info": {
    "vcs_url": "https://github.com/CircleCI-Public/test-project",
    "provider": "GitHub",
    "default_branch": "main"
  }
}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	project, err := client.CreateProject(context.Background(), "gh/CircleCI-Public", circleci.CreateProjectBody{
		Name: "test-project",
	})

	assertNoError(t, err)
	assertEqual(t, "gh/CircleCI-Public/test-project", project.Slug)
	assertEqual(t, "test-project", project.Name)
	assertEqual(t, "12345", project.ID)
	assertEqual(t, "GitHub", project.VCSInfo.Provider)
}

func TestCreateProject_Failed(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/organization/gh/CircleCI-Public/project")
		assertMethod(t, r, http.MethodPost)
		assertAuth(t, r)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Internal Server Error"}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	_, err := client.CreateProject(context.Background(), "gh/CircleCI-Public", circleci.CreateProjectBody{
		Name: "test-project",
	})

	assertError(t, err)
}

func TestDeleteProject_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/test-project")
		assertMethod(t, r, http.MethodDelete)
		assertAuth(t, r)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Project deleted."}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.DeleteProject(context.Background(), "gh/CircleCI-Public/test-project")

	assertNoError(t, err)
}

func TestDeleteProject_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Project not found."}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.DeleteProject(context.Background(), "gh/CircleCI-Public/missing-project")

	assertError(t, err)
	if !errors.Is(err, circleci.ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}

func TestGetProject_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/test-project")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
  "slug": "gh/CircleCI-Public/test-project",
  "name": "test-project",
  "id": "12345",
  "organization_name": "CircleCI-Public",
  "organization_slug": "gh/CircleCI-Public",
  "organization_id": "1111",
  "vcs_info": {
    "vcs_url": "https://github.com/CircleCI-Public/test-project",
    "provider": "GitHub",
    "default_branch": "main"
  }
}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	project, err := client.GetProject(context.Background(), "gh/CircleCI-Public/test-project")

	assertNoError(t, err)
	assertEqual(t, "gh/CircleCI-Public/test-project", project.Slug)
	assertEqual(t, "test-project", project.Name)
	assertEqual(t, "12345", project.ID)
	assertEqual(t, "CircleCI-Public", project.OrganizationName)
	assertEqual(t, "main", project.VCSInfo.DefaultBranch)
}

func TestGetProject_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Project not found."}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	_, err := client.GetProject(context.Background(), "gh/CircleCI-Public/missing-project")

	assertError(t, err)
	if !errors.Is(err, circleci.ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}

func TestGetAllCheckoutKeys_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/test-project/checkout-key")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
  "items": [
    {
      "public-key": "ssh-rsa AAAAB3NzaC1yc2E...",
      "type": "deploy-key",
      "fingerprint": "5a:1b:2c:3d:4e:5f",
      "preferred": true,
      "created-at": "2019-08-24T14:15:22Z"
    }
  ],
  "next_page_token": null
}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	keys, err := client.GetAllCheckoutKeys(context.Background(), "gh/CircleCI-Public/test-project")

	assertNoError(t, err)
	assertEqual(t, 1, len(keys))
	if len(keys) > 0 {
		assertEqual(t, "deploy-key", keys[0].Type)
		assertEqual(t, "5a:1b:2c:3d:4e:5f", keys[0].Fingerprint)
		assertEqual(t, true, keys[0].Preferred)
	}
}

func TestGetAllCheckoutKeys_Failed(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/test-project/checkout-key")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Internal Server Error"}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	_, err := client.GetAllCheckoutKeys(context.Background(), "gh/CircleCI-Public/test-project")

	assertError(t, err)
}

func TestCreateCheckoutKey_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/test-project/checkout-key")
		assertMethod(t, r, http.MethodPost)
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{
  "public-key": "ssh-rsa AAAAB3NzaC1yc2E...",
  "type": "deploy-key",
  "fingerprint": "5a:1b:2c:3d:4e:5f",
  "preferred": true,
  "created-at": "2019-08-24T14:15:22Z"
}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	key, err := client.CreateCheckoutKey(context.Background(), "gh/CircleCI-Public/test-project", circleci.CreateCheckoutKeyBody{
		Type: circleci.CheckoutKeyTypeDeployKey,
	})

	assertNoError(t, err)
	assertEqual(t, "deploy-key", key.Type)
	assertEqual(t, "5a:1b:2c:3d:4e:5f", key.Fingerprint)
	assertEqual(t, true, key.Preferred)
}

func TestCreateCheckoutKey_Failed(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/test-project/checkout-key")
		assertMethod(t, r, http.MethodPost)
		assertAuth(t, r)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Internal Server Error"}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	_, err := client.CreateCheckoutKey(context.Background(), "gh/CircleCI-Public/test-project", circleci.CreateCheckoutKeyBody{
		Type: circleci.CheckoutKeyTypeDeployKey,
	})

	assertError(t, err)
}

func TestDeleteCheckoutKey_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/test-project/checkout-key/5a:1b:2c:3d:4e:5f")
		assertMethod(t, r, http.MethodDelete)
		assertAuth(t, r)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Checkout key deleted."}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.DeleteCheckoutKey(context.Background(), "gh/CircleCI-Public/test-project", "5a:1b:2c:3d:4e:5f")

	assertNoError(t, err)
}

func TestDeleteCheckoutKey_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Checkout key not found."}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.DeleteCheckoutKey(context.Background(), "gh/CircleCI-Public/test-project", "00:00:00")

	assertError(t, err)
	if !errors.Is(err, circleci.ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}

func TestGetCheckoutKey_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/test-project/checkout-key/5a:1b:2c:3d:4e:5f")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
  "public-key": "ssh-rsa AAAAB3NzaC1yc2E...",
  "type": "github-user-key",
  "fingerprint": "5a:1b:2c:3d:4e:5f",
  "preferred": false,
  "created-at": "2019-08-24T14:15:22Z"
}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	key, err := client.GetCheckoutKey(context.Background(), "gh/CircleCI-Public/test-project", "5a:1b:2c:3d:4e:5f")

	assertNoError(t, err)
	assertEqual(t, "github-user-key", key.Type)
	assertEqual(t, "5a:1b:2c:3d:4e:5f", key.Fingerprint)
	assertEqual(t, false, key.Preferred)
}

func TestGetCheckoutKey_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Checkout key not found."}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	_, err := client.GetCheckoutKey(context.Background(), "gh/CircleCI-Public/test-project", "00:00:00")

	assertError(t, err)
	if !errors.Is(err, circleci.ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}

func TestListEvironmentVariables_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/test-project/envvar")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
  "items": [
    {
      "name": "MY_ENV_VAR",
      "value": "some-value",
      "created-at": "2019-08-24T14:15:22Z"
    }
  ],
  "next_page_token": null
}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	vars, err := client.ListEvironmentVariables(context.Background(), "gh/CircleCI-Public/test-project")

	assertNoError(t, err)
	assertEqual(t, 1, len(vars))
	if len(vars) > 0 {
		assertEqual(t, "MY_ENV_VAR", vars[0].Name)
		assertEqual(t, "some-value", vars[0].Value)
	}
}

func TestListEvironmentVariables_Failed(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/test-project/envvar")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Internal Server Error"}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	_, err := client.ListEvironmentVariables(context.Background(), "gh/CircleCI-Public/test-project")

	assertError(t, err)
}

func TestCreateEnvironmentVariables_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/test-project/envvar")
		assertMethod(t, r, http.MethodPost)
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{
  "name": "MY_ENV_VAR",
  "value": "some-value",
  "created-at": "2019-08-24T14:15:22Z"
}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	envVar, err := client.CreateEnvironmentVariables(context.Background(), "gh/CircleCI-Public/test-project", circleci.CreateProjectEnvVarBody{
		Name:  "MY_ENV_VAR",
		Value: "some-value",
	})

	assertNoError(t, err)
	assertEqual(t, "MY_ENV_VAR", envVar.Name)
	assertEqual(t, "some-value", envVar.Value)
}

func TestCreateEnvironmentVariables_Failed(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/test-project/envvar")
		assertMethod(t, r, http.MethodPost)
		assertAuth(t, r)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Internal Server Error"}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	_, err := client.CreateEnvironmentVariables(context.Background(), "gh/CircleCI-Public/test-project", circleci.CreateProjectEnvVarBody{
		Name:  "MY_ENV_VAR",
		Value: "some-value",
	})

	assertError(t, err)
}

func TestDeleteEnvironmentVariables_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/test-project/envvar/MY_ENV_VAR")
		assertMethod(t, r, http.MethodDelete)
		assertAuth(t, r)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Environment variable deleted."}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.DeleteEnvironmentVariables(context.Background(), "gh/CircleCI-Public/test-project", "MY_ENV_VAR")

	assertNoError(t, err)
}

func TestDeleteEnvironmentVariables_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Environment variable not found."}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.DeleteEnvironmentVariables(context.Background(), "gh/CircleCI-Public/test-project", "MISSING")

	assertError(t, err)
	if !errors.Is(err, circleci.ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}

func TestGetMaskedEnvironmentVariable_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/test-project/envvar/MY_ENV_VAR")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
  "name": "MY_ENV_VAR",
  "value": null,
  "created-at": "2019-08-24T14:15:22Z"
}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	envVar, err := client.GetMaskedEnvironmentVariable(context.Background(), "gh/CircleCI-Public/test-project", "MY_ENV_VAR")

	assertNoError(t, err)
	assertEqual(t, "MY_ENV_VAR", envVar.Name)
}

func TestGetMaskedEnvironmentVariable_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Environment variable not found."}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	_, err := client.GetMaskedEnvironmentVariable(context.Background(), "gh/CircleCI-Public/test-project", "MISSING")

	assertError(t, err)
	if !errors.Is(err, circleci.ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}

func TestGetProjectSettings_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/test-project/settings")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
  "advanced": {
    "autocancel_builds": true,
    "build_fork_prs": false,
    "build_prs_only": false,
    "disable_ssh": true,
    "forks_receive_secret_env_vars": false,
    "oss": true,
    "set_github_status": true,
    "setup_workflows": false,
    "write_settings_requires_admin": true,
    "pr_only_branch_overrides": false
  }
}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	settings, err := client.GetProjectSettings(context.Background(), "gh", "CircleCI-Public", "test-project")

	assertNoError(t, err)
	assertEqual(t, true, settings.AutocancelBuilds)
	assertEqual(t, true, settings.DisableSSH)
	assertEqual(t, true, settings.OSS)
}

func TestGetProjectSettings_Failed(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/test-project/settings")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Internal Server Error"}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	_, err := client.GetProjectSettings(context.Background(), "gh", "CircleCI-Public", "test-project")

	assertError(t, err)
}

func TestUpdateProjectSettings_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/test-project/settings")
		assertMethod(t, r, http.MethodPatch)
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
  "advanced": {
    "autocancel_builds": true,
    "build_fork_prs": false,
    "build_prs_only": false,
    "disable_ssh": false,
    "forks_receive_secret_env_vars": false,
    "oss": false,
    "set_github_status": false,
    "setup_workflows": false,
    "write_settings_requires_admin": false,
    "pr_only_branch_overrides": false
  }
}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	settings, err := client.UpdateProjectSettings(context.Background(), "gh", "CircleCI-Public", "test-project", circleci.UpdateProjectSettingsBody{
		Advanced: circleci.ProjectSettings{AutocancelBuilds: true},
	})

	assertNoError(t, err)
	assertEqual(t, true, settings.AutocancelBuilds)
	assertEqual(t, false, settings.DisableSSH)
}

func TestUpdateProjectSettings_Failed(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/CircleCI-Public/test-project/settings")
		assertMethod(t, r, http.MethodPatch)
		assertAuth(t, r)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Internal Server Error"}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	_, err := client.UpdateProjectSettings(context.Background(), "gh", "CircleCI-Public", "test-project", circleci.UpdateProjectSettingsBody{
		Advanced: circleci.ProjectSettings{AutocancelBuilds: true},
	})

	assertError(t, err)
}
