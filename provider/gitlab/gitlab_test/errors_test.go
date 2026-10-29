package gitlab_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oktanthq/ciforge/provider/gitlab"
)

func TestAuthHeaders(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, "test-token", r.Header.Get("PRIVATE-TOKEN"))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]any{})
	}))
	defer srv.Close()

	client := gitlab.New("test-token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	_, err := client.ListPipelines(context.Background(), 42, nil)
	assertNoError(t, err)
}

func TestNotFoundError(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "404 Not Found"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	_, err := client.GetPipeline(context.Background(), 42, 999)
	assertError(t, err)
	if !errors.Is(err, gitlab.ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}

func TestUnauthorizedError(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"message": "401 Unauthorized"})
	}))
	defer srv.Close()

	client := gitlab.New("bad-token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	_, err := client.ListJobsForProject(context.Background(), 42, nil)
	assertError(t, err)
	if !errors.Is(err, gitlab.ErrUnauthorized) {
		t.Errorf("err = %v, want ErrUnauthorized", err)
	}
}

func TestRateLimitedError(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		json.NewEncoder(w).Encode(map[string]string{"message": "Retry later"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	_, err := client.ListRunners(context.Background(), nil)
	assertError(t, err)
	if !errors.Is(err, gitlab.ErrRateLimited) {
		t.Errorf("err = %v, want ErrRateLimited", err)
	}
}

func TestServerError(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "boom"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	err := client.DeleteTriggerToken(context.Background(), 42, 1)
	assertError(t, err)
}
