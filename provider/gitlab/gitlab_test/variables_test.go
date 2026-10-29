package gitlab_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oktanthq/ciforge/provider/gitlab"
)

func TestListProjectVariables(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/variables")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)
		assertQuery(t, r, "page", "2")
		assertQuery(t, r, "per_page", "10")

		json.NewEncoder(w).Encode([]map[string]any{
			{"key": "TOKEN", "value": "secret", "variable_type": "env_var", "protected": true, "masked": true, "environment_scope": "*"},
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	vars, err := client.ListProjectVariables(context.Background(), 42, &gitlab.ListProjectVariablesOptions{
		Page:    2,
		PerPage: 10,
	})
	assertNoError(t, err)
	assertEqual(t, 1, len(vars))
	assertEqual(t, "TOKEN", vars[0].Key)
	assertEqual(t, true, vars[0].Masked)
}

func TestGetProjectVariable(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/variables/TOKEN")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"key": "TOKEN", "value": "secret", "variable_type": "env_var", "environment_scope": "*",
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	v, err := client.GetProjectVariable(context.Background(), 42, "TOKEN", nil)
	assertNoError(t, err)
	assertEqual(t, "TOKEN", v.Key)
	assertEqual(t, "secret", v.Value)
}

func TestCreateProjectVariable(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/variables")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{
			"key": "TOKEN", "value": "secret", "variable_type": "env_var", "masked": true,
		})
	}))
	defer srv.Close()

	masked := true
	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	v, err := client.CreateProjectVariable(context.Background(), 42, &gitlab.CreateProjectVariableBody{
		Key:    "TOKEN",
		Value:  "secret",
		Masked: &masked,
	})
	assertNoError(t, err)
	assertEqual(t, "TOKEN", v.Key)
	assertEqual(t, true, v.Masked)
}

func TestUpdateProjectVariable(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/variables/TOKEN")
		assertMethod(t, r, http.MethodPut)

		json.NewEncoder(w).Encode(map[string]any{"key": "TOKEN", "value": "rotated", "variable_type": "env_var"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	v, err := client.UpdateProjectVariable(context.Background(), 42, "TOKEN", &gitlab.UpdateProjectVariableBody{Value: "rotated"})
	assertNoError(t, err)
	assertEqual(t, "rotated", v.Value)
}

func TestDeleteProjectVariable(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/variables/TOKEN")
		assertMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	err := client.DeleteProjectVariable(context.Background(), 42, "TOKEN", nil)
	assertNoError(t, err)
}
