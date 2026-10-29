package gitlab_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oktanthq/ciforge/provider/gitlab"
)

func TestListRunners(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/runners")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)
		assertQuery(t, r, "type", "project_type")
		assertQuery(t, r, "status", "online")

		json.NewEncoder(w).Encode([]map[string]any{
			{"id": 1, "description": "linux", "active": true, "runner_type": "project_type", "status": "online"},
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	runners, err := client.ListRunners(context.Background(), &gitlab.ListRunnersOptions{
		Type:   gitlab.RunnerTypeProject,
		Status: gitlab.RunnerStatusOnline,
	})
	assertNoError(t, err)
	assertEqual(t, 1, len(runners))
	assertEqual(t, 1, runners[0].ID)
	assertEqual(t, "linux", runners[0].Description)
}

func TestListAllRunners(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/runners/all")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode([]map[string]any{
			{"id": 2, "description": "shared", "runner_type": "instance_type", "status": "offline"},
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	runners, err := client.ListAllRunners(context.Background(), nil)
	assertNoError(t, err)
	assertEqual(t, 1, len(runners))
	assertEqual(t, gitlab.RunnerTypeInstance, runners[0].RunnerType)
}

func TestGetRunner(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/runners/1")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{"id": 1, "description": "linux", "status": "online"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	runner, err := client.GetRunner(context.Background(), 1)
	assertNoError(t, err)
	assertEqual(t, 1, runner.ID)
	assertEqual(t, "linux", runner.Description)
}

func TestUpdateRunner(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/runners/1")
		assertMethod(t, r, http.MethodPut)

		json.NewEncoder(w).Encode(map[string]any{"id": 1, "description": "updated", "status": "online"})
	}))
	defer srv.Close()

	desc := "updated"
	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	runner, err := client.UpdateRunner(context.Background(), 1, &gitlab.UpdateRunnerBody{Description: &desc})
	assertNoError(t, err)
	assertEqual(t, "updated", runner.Description)
}

func TestListRunnerJobs(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/runners/1/jobs")
		assertMethod(t, r, http.MethodGet)
		assertQuery(t, r, "status", "running")

		json.NewEncoder(w).Encode([]map[string]any{
			{"id": 9, "name": "test", "status": "running", "stage": "test"},
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	jobs, err := client.ListRunnerJobs(context.Background(), 1, &gitlab.ListRunnerJobsOptions{Status: "running"})
	assertNoError(t, err)
	assertEqual(t, 1, len(jobs))
	assertEqual(t, 9, jobs[0].ID)
}

func TestListRunnerManagers(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/runners/1/managers")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode([]map[string]any{
			{"id": 3, "system_id": "s-1", "version": "17.0", "status": "online", "created_at": "2026-01-02T15:04:05Z", "contacted_at": "2026-01-02T15:04:05Z"},
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	managers, err := client.ListRunnerManagers(context.Background(), 1)
	assertNoError(t, err)
	assertEqual(t, 1, len(managers))
	assertEqual(t, "s-1", managers[0].SystemID)
}

func TestListProjectRunners(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/runners")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode([]map[string]any{{"id": 1, "description": "proj", "status": "online"}})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	runners, err := client.ListProjectRunners(context.Background(), 42, nil)
	assertNoError(t, err)
	assertEqual(t, 1, len(runners))
}

func TestAssignRunnerToProject(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/runners")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{"id": 1, "description": "assigned"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	runner, err := client.AssignRunnerToProject(context.Background(), 42, 1)
	assertNoError(t, err)
	assertEqual(t, 1, runner.ID)
}

func TestUnassignRunnerFromProject(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/runners/1")
		assertMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	err := client.UnassignRunnerFromProject(context.Background(), 42, 1)
	assertNoError(t, err)
}

func TestListGroupRunners(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/groups/5/runners")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode([]map[string]any{{"id": 4, "description": "group"}})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	runners, err := client.ListGroupRunners(context.Background(), 5, nil)
	assertNoError(t, err)
	assertEqual(t, 1, len(runners))
}

func TestCreateRunner(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/runners")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{"id": 9, "token": "glrt-xxx"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	resp, err := client.CreateRunner(context.Background(), &gitlab.CreateRunnerBody{Token: "reg-token", Description: "new"})
	assertNoError(t, err)
	assertEqual(t, 9, resp.ID)
	assertEqual(t, "glrt-xxx", resp.Token)
}

func TestDeleteRunner(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/runners/9")
		assertMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	err := client.DeleteRunner(context.Background(), 9)
	assertNoError(t, err)
}

func TestDeleteRunnerByToken(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/runners")
		assertMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	err := client.DeleteRunnerByToken(context.Background(), "glrt-xxx")
	assertNoError(t, err)
}

func TestVerifyRunner(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/runners/verify")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{"id": 9, "token": "glrt-xxx"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	resp, err := client.VerifyRunner(context.Background(), "glrt-xxx", "s-1")
	assertNoError(t, err)
	assertEqual(t, 9, resp.ID)
}

func TestResetInstanceRegistrationToken(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/runners/reset_registration_token")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{"token": "new-reg"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	token, err := client.ResetInstanceRegistrationToken(context.Background())
	assertNoError(t, err)
	assertEqual(t, "new-reg", token)
}

func TestResetProjectRegistrationToken(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/runners/reset_registration_token")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{"token": "proj-reg"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	token, err := client.ResetProjectRegistrationToken(context.Background(), 42)
	assertNoError(t, err)
	assertEqual(t, "proj-reg", token)
}

func TestResetGroupRegistrationToken(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/groups/5/runners/reset_registration_token")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{"token": "group-reg"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	token, err := client.ResetGroupRegistrationToken(context.Background(), 5)
	assertNoError(t, err)
	assertEqual(t, "group-reg", token)
}

func TestResetRunnerAuthenticationToken(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/runners/9/reset_authentication_token")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{"id": 9, "token": "new-auth"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	resp, err := client.ResetRunnerAuthenticationToken(context.Background(), 9)
	assertNoError(t, err)
	assertEqual(t, "new-auth", resp.Token)
}

func TestResetRunnerAuthenticationTokenByToken(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/runners/reset_authentication_token")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{"id": 9, "token": "newer-auth"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	resp, err := client.ResetRunnerAuthenticationTokenByToken(context.Background(), "old")
	assertNoError(t, err)
	assertEqual(t, "newer-auth", resp.Token)
}

func TestGetJobRouterDiscovery(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/runners/router/discovery")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{"server_url": "https://jobs.example.com"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	resp, err := client.GetJobRouterDiscovery(context.Background())
	assertNoError(t, err)
	assertEqual(t, "https://jobs.example.com", resp.ServerURL)
}
