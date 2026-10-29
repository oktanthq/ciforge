package gitlab_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oktanthq/ciforge/provider/gitlab"
)

func TestListPipelineSchedules(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipeline_schedules")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)
		assertQuery(t, r, "scope", "active")

		json.NewEncoder(w).Encode([]map[string]any{
			{"id": 1, "description": "nightly", "ref": "main", "cron": "0 0 * * *", "active": true, "created_at": "2026-01-02T15:04:05Z", "updated_at": "2026-01-02T15:04:05Z", "next_run_at": "2026-01-03T00:00:00Z"},
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	schedules, err := client.ListPipelineSchedules(context.Background(), 42, &gitlab.ListPipelineSchedulesOptions{
		Scope: gitlab.PipelineScheduleScopeActive,
	})
	assertNoError(t, err)
	assertEqual(t, 1, len(schedules))
	assertEqual(t, "nightly", schedules[0].Description)
	assertEqual(t, true, schedules[0].Active)
}

func TestGetPipelineSchedule(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipeline_schedules/1")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"id": 1, "description": "nightly", "ref": "main", "cron": "0 0 * * *", "active": true,
			"created_at": "2026-01-02T15:04:05Z", "updated_at": "2026-01-02T15:04:05Z", "next_run_at": "2026-01-03T00:00:00Z",
			"variables": []map[string]any{{"key": "ENV", "value": "prod", "variable_type": "env_var"}},
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	schedule, err := client.GetPipelineSchedule(context.Background(), 42, 1)
	assertNoError(t, err)
	assertEqual(t, 1, schedule.ID)
	assertEqual(t, 1, len(schedule.Variables))
	assertEqual(t, "ENV", schedule.Variables[0].Key)
}

func TestListPipelinesBySchedule(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipeline_schedules/1/pipelines")
		assertMethod(t, r, http.MethodGet)
		assertQuery(t, r, "status", "success")

		json.NewEncoder(w).Encode([]map[string]any{
			{"id": 50, "status": "success", "ref": "main", "created_at": "2026-01-02T15:04:05Z", "updated_at": "2026-01-02T15:04:05Z"},
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	pipelines, err := client.ListPipelinesBySchedule(context.Background(), 42, 1, &gitlab.ListPipelinesByScheduleOptions{
		Status: gitlab.PipelineStatusSuccess,
	})
	assertNoError(t, err)
	assertEqual(t, 1, len(pipelines))
	assertEqual(t, 50, pipelines[0].ID)
}

func TestCreatePipelineSchedule(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipeline_schedules")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{
			"id": 2, "description": "hourly", "ref": "main", "cron": "0 * * * *", "active": true,
			"created_at": "2026-01-02T15:04:05Z", "updated_at": "2026-01-02T15:04:05Z", "next_run_at": "2026-01-02T16:00:00Z",
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	schedule, err := client.CreatePipelineSchedule(context.Background(), 42, &gitlab.CreatePipelineScheduleBody{
		Cron:        "0 * * * *",
		Description: "hourly",
		Ref:         "main",
	})
	assertNoError(t, err)
	assertEqual(t, 2, schedule.ID)
	assertEqual(t, "hourly", schedule.Description)
}

func TestUpdatePipelineSchedule(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipeline_schedules/2")
		assertMethod(t, r, http.MethodPut)

		json.NewEncoder(w).Encode(map[string]any{
			"id": 2, "description": "daily", "ref": "main", "cron": "0 0 * * *",
			"created_at": "2026-01-02T15:04:05Z", "updated_at": "2026-01-02T15:04:05Z", "next_run_at": "2026-01-03T00:00:00Z",
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	schedule, err := client.UpdatePipelineSchedule(context.Background(), 42, 2, &gitlab.UpdatePipelineScheduleBody{
		Description: "daily",
		Cron:        "0 0 * * *",
	})
	assertNoError(t, err)
	assertEqual(t, "daily", schedule.Description)
}

func TestTakeOwnershipPipelineSchedule(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipeline_schedules/2/take_ownership")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{
			"id": 2, "description": "daily",
			"created_at": "2026-01-02T15:04:05Z", "updated_at": "2026-01-02T15:04:05Z", "next_run_at": "2026-01-03T00:00:00Z",
			"owner": map[string]any{"id": 8, "username": "me"},
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	schedule, err := client.TakeOwnershipPipelineSchedule(context.Background(), 42, 2)
	assertNoError(t, err)
	assertEqual(t, "me", schedule.Owner.Username)
}

func TestDeletePipelineSchedule(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipeline_schedules/2")
		assertMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	err := client.DeletePipelineSchedule(context.Background(), 42, 2)
	assertNoError(t, err)
}

func TestPlayPipelineSchedule(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipeline_schedules/2/play")
		assertMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	err := client.PlayPipelineSchedule(context.Background(), 42, 2)
	assertNoError(t, err)
}

func TestCreateScheduleVariable(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipeline_schedules/2/variables")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{"key": "ENV", "value": "prod", "variable_type": "env_var"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	v, err := client.CreateScheduleVariable(context.Background(), 42, 2, &gitlab.PipelineScheduleVariableBody{
		Key:   "ENV",
		Value: "prod",
	})
	assertNoError(t, err)
	assertEqual(t, "ENV", v.Key)
	assertEqual(t, "prod", v.Value)
}

func TestGetScheduleVariable(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipeline_schedules/2/variables/ENV")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{"key": "ENV", "value": "prod", "variable_type": "env_var"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	v, err := client.GetScheduleVariable(context.Background(), 42, 2, "ENV")
	assertNoError(t, err)
	assertEqual(t, "prod", v.Value)
}

func TestUpdateScheduleVariable(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipeline_schedules/2/variables/ENV")
		assertMethod(t, r, http.MethodPut)

		json.NewEncoder(w).Encode(map[string]any{"key": "ENV", "value": "staging", "variable_type": "env_var"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	v, err := client.UpdateScheduleVariable(context.Background(), 42, 2, "ENV", &gitlab.PipelineScheduleVariableBody{Value: "staging"})
	assertNoError(t, err)
	assertEqual(t, "staging", v.Value)
}

func TestDeleteScheduleVariable(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/pipeline_schedules/2/variables/ENV")
		assertMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	err := client.DeleteScheduleVariable(context.Background(), 42, 2, "ENV")
	assertNoError(t, err)
}
