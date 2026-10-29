package circleci_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oktanthq/ciforge/provider/circleci"
)

func TestGetJob(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/jobs/j1")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		json.NewEncoder(w).Encode(map[string]any{
			"id": "j1", "name": "test", "status": "success", "job_number": 12, "type": "build",
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	job, err := client.GetJob(context.Background(), "j1")
	assertNoError(t, err)
	assertEqual(t, "j1", job.ID)
	assertEqual(t, "test", job.Name)
	assertEqual(t, "success", job.Status)
}

func TestCancelJob(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/jobs/j1/cancel")
		assertMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusAccepted)
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.CancelJob(context.Background(), "j1")
	assertNoError(t, err)
}

func TestCancelJobByJobNumber(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/project/gh/org/proj/job/12/cancel")
		assertMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusAccepted)
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	err := client.CancelJobByJobNumber(context.Background(), "gh/org/proj", 12)
	assertNoError(t, err)
}
