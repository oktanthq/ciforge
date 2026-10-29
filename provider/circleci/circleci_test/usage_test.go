package circleci_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/oktanthq/ciforge/provider/circleci"
)

func TestCreateUsageExport(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/organizations/org-1/usage_export_job")
		assertMethod(t, r, http.MethodPost)
		assertAuth(t, r)

		json.NewEncoder(w).Encode(map[string]any{
			"usage_export_job_id": "exp-1", "state": "created",
			"start": "2026-01-01T00:00:00Z", "end": "2026-01-31T00:00:00Z",
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	exp, err := client.CreateUsageExport(context.Background(), "org-1", circleci.UsageExportBody{
		Start: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2026, 1, 31, 0, 0, 0, 0, time.UTC),
	})
	assertNoError(t, err)
	assertEqual(t, "exp-1", exp.UsageExportJobID)
	assertEqual(t, "created", exp.State)
}

func TestGetUsageExport(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/organizations/org-1/usage_export_job/exp-1")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{
			"usage_export_job_id": "exp-1", "state": "completed",
			"download_urls": []string{"https://example.com/export.csv"},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	exp, err := client.GetUsageExport(context.Background(), "org-1", "exp-1")
	assertNoError(t, err)
	assertEqual(t, "completed", exp.State)
	assertEqual(t, 1, len(exp.DownloadURLs))
}
