package gitlab_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oktanthq/ciforge/provider/gitlab"
)

func TestDownloadJobArtifacts(t *testing.T) {
	t.Parallel()

	expected := []byte("zip-bytes")
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/jobs/7/artifacts")
		assertMethod(t, r, http.MethodGet)
		assertQuery(t, r, "job_token", "jt")
		w.Write(expected)
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	data, err := client.DownloadJobArtifacts(context.Background(), 42, 7, "jt")
	assertNoError(t, err)
	assertEqual(t, string(expected), string(data))
}

func TestDownloadJobArtifactsByRef(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/jobs/artifacts/main/download")
		assertMethod(t, r, http.MethodGet)
		assertQuery(t, r, "job", "build")
		w.Write([]byte("archive"))
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	data, err := client.DownloadJobArtifactsByRef(context.Background(), 42, "main", "build", nil)
	assertNoError(t, err)
	assertEqual(t, "archive", string(data))
}

func TestDownloadSingleArtifactFile(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/jobs/7/artifacts/report.xml")
		assertMethod(t, r, http.MethodGet)
		w.Write([]byte("<report/>"))
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	data, err := client.DownloadSingleArtifactFile(context.Background(), 42, 7, "report.xml", "")
	assertNoError(t, err)
	assertEqual(t, "<report/>", string(data))
}

func TestDownloadSingleArtifactFileByRef(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/jobs/artifacts/main/raw/report.xml")
		assertMethod(t, r, http.MethodGet)
		assertQuery(t, r, "job", "build")
		w.Write([]byte("file"))
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	data, err := client.DownloadSingleArtifactFileByRef(context.Background(), 42, "main", "report.xml", "build", nil)
	assertNoError(t, err)
	assertEqual(t, "file", string(data))
}

func TestListArtifactTree(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/jobs/7/artifacts/tree")
		assertMethod(t, r, http.MethodGet)
		assertQuery(t, r, "path", "dist")
		assertQuery(t, r, "recursive", "true")

		json.NewEncoder(w).Encode([]map[string]any{
			{"name": "app", "path": "dist/app", "type": "file", "mode": "0644"},
		})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	entries, err := client.ListArtifactTree(context.Background(), 42, 7, &gitlab.ListArtifactTreeOptions{
		Path:      "dist",
		Recursive: true,
	})
	assertNoError(t, err)
	assertEqual(t, 1, len(entries))
	assertEqual(t, "app", entries[0].Name)
	assertEqual(t, "file", entries[0].Type)
}

func TestKeepJobArtifacts(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/jobs/7/artifacts/keep")
		assertMethod(t, r, http.MethodPost)

		json.NewEncoder(w).Encode(map[string]any{"id": 7, "name": "build", "status": "success"})
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	job, err := client.KeepJobArtifacts(context.Background(), 42, 7)
	assertNoError(t, err)
	assertEqual(t, 7, job.ID)
}

func TestDeleteJobArtifacts(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/jobs/7/artifacts")
		assertMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	err := client.DeleteJobArtifacts(context.Background(), 42, 7)
	assertNoError(t, err)
}

func TestDeleteAllJobArtifacts(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/projects/42/artifacts")
		assertMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusAccepted)
	}))
	defer srv.Close()

	client := gitlab.New("token", gitlab.WithBaseURL(srv.URL), gitlab.WithHTTPClient(srv.Client()))
	err := client.DeleteAllJobArtifacts(context.Background(), 42)
	assertNoError(t, err)
}
