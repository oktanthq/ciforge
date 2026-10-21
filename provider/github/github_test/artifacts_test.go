package github_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/oktanthq/ciforge/provider/github"
)

func TestListArtifacts(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/artifacts")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"total_count": 2,
			"artifacts": []map[string]interface{}{
				{
					"id": 1, "name": "build.zip", "size_in_bytes": 1024,
					"created_at": time.Now(), "expired": false,
					"archive_download_url": "https://api.github.com/repos/owner/repo/actions/artifacts/1/zip",
					"workflow_run":        map[string]interface{}{"id": 100},
				},
				{
					"id": 2, "name": "test-report.xml", "size_in_bytes": 512,
					"created_at": time.Now(), "expired": true,
					"archive_download_url": "https://api.github.com/repos/owner/repo/actions/artifacts/2/zip",
					"workflow_run":        map[string]interface{}{"id": 100},
				},
			},
		})
	}))
	defer srv.Close()

	client := github.New("token", github.WithBaseURL(srv.URL), github.WithHTTPClient(srv.Client()))
	artifacts, err := client.ListArtifacts(context.Background(), "repo", "owner", nil)
	assertNoError(t, err)
	assertEqual(t, 2, len(artifacts))
	assertEqual(t, int64(1), artifacts[0].ID)
	assertEqual(t, "build.zip", artifacts[0].Name)
	assertEqual(t, false, artifacts[0].Expired)
	assertEqual(t, int64(2), artifacts[1].ID)
	assertEqual(t, "test-report.xml", artifacts[1].Name)
	assertEqual(t, true, artifacts[1].Expired)
}

func TestGetArtifact(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/artifacts/1")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"id": 1, "name": "build.zip", "size_in_bytes": 1024,
			"created_at": time.Now(), "expired": false,
			"archive_download_url": "https://api.github.com/repos/owner/repo/actions/artifacts/1/zip",
			"workflow_run":         map[string]interface{}{"id": 100},
		})
	}))
	defer srv.Close()

	client := github.New("token", github.WithBaseURL(srv.URL), github.WithHTTPClient(srv.Client()))
	artifact, err := client.GetArtifact(context.Background(), "repo", "owner", 1)
	assertNoError(t, err)
	assertEqual(t, int64(1), artifact.ID)
	assertEqual(t, "build.zip", artifact.Name)
	assertEqual(t, false, artifact.Expired)
}

func TestDeleteArtifact(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/artifacts/1")
		assertMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := github.New("token", github.WithBaseURL(srv.URL), github.WithHTTPClient(srv.Client()))
	err := client.DeleteArtifact(context.Background(), "repo", "owner", 1)
	assertNoError(t, err)
}

func TestListWorkflowRunArtifacts(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/runs/100/artifacts")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"total_count": 1,
			"artifacts": []map[string]interface{}{
				{
					"id": 1, "name": "build.zip", "size_in_bytes": 1024,
					"created_at": time.Now(), "expired": false,
					"archive_download_url": "https://api.github.com/repos/owner/repo/actions/artifacts/1/zip",
					"workflow_run":         map[string]interface{}{"id": 100},
				},
			},
		})
	}))
	defer srv.Close()

	client := github.New("token", github.WithBaseURL(srv.URL), github.WithHTTPClient(srv.Client()))
	artifacts, err := client.ListWorkflowRunArtifacts(context.Background(), "repo", "owner", 100, nil)
	assertNoError(t, err)
	assertEqual(t, 1, len(artifacts))
	assertEqual(t, int64(1), artifacts[0].ID)
	assertEqual(t, "build.zip", artifacts[0].Name)
}

func TestDownloadArtifact(t *testing.T) {
	t.Parallel()

	expected := []byte("zip-content")
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/repos/owner/repo/actions/artifacts/1/zip")
		assertMethod(t, r, http.MethodGet)
		w.Header().Set("Content-Type", "application/zip")
		w.Write(expected)
	}))
	defer srv.Close()

	client := github.New("token", github.WithBaseURL(srv.URL), github.WithHTTPClient(srv.Client()))
	data, err := client.DownloadArtifact(context.Background(), srv.URL+"/repos/owner/repo/actions/artifacts/1/zip")
	assertNoError(t, err)
	assertEqual(t, string(expected), string(data))
}
