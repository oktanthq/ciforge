package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/abdulahshoaib/ciforge/ci"
	"github.com/abdulahshoaib/ciforge/models"
	gh "github.com/abdulahshoaib/ciforge/provider/github"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("set GITHUB_TOKEN env var")
	}

	// ── Create client ───────────────────────────────────────────────
	client := gh.New(token)

	// Optional: custom HTTP client, custom base URL (GitHub Enterprise)
	_ = gh.New(token,
		gh.WithHTTPClient(&http.Client{Timeout: 30 * time.Second}),
		gh.WithBaseURL("https://github.example.com/api/v3"),
	)

	// Wrap with ci.Client for provider-agnostic usage
	ciClient := ci.New(ci.WithProvider(client))

	_ = ciClient

	ctx := context.Background()
	owner := "owner"
	repo := "repo"

	// ── 1. List workflows ──────────────────────────────────────────
	workflows, err := client.ListWorkflows(ctx, repo, owner)
	if err != nil {
		log.Fatal(err)
	}
	for _, w := range workflows {
		fmt.Printf("workflow: %s (id=%s, disabled=%v)\n", w.Name, w.ID, w.Disabled)
	}

	// ── 2. Get a single workflow ───────────────────────────────────
	workflow, err := client.GetWorkflow(ctx, repo, owner, "12345")
	if err != nil {
		if errors.Is(err, ci.ErrNotFound) {
			fmt.Println("workflow not found")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("workflow: %s, url: %s\n", workflow.Name, workflow.URL)
	}

	// ── 3. List runs for a workflow ─────────────────────────────────
	runs, err := client.ListRunsForWorkflow(ctx, repo, owner, "12345")
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range runs {
		fmt.Printf("run #%d: status=%s result=%s branch=%s sha=%s\n",
			r.Number, r.Status, r.Result, r.Branch, r.CommitSHA)
	}

	// ── 4. Get a single run ────────────────────────────────────────
	run, err := client.GetWorkflowRun(ctx, repo, owner, "99999")
	if err != nil {
		if errors.Is(err, ci.ErrUnauthorized) {
			fmt.Println("bad token")
		} else if errors.Is(err, ci.ErrRateLimited) {
			fmt.Println("rate limited — back off")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("run #%d: %s (took %s)\n", run.Number, run.Result, run.Duration)
	}

	// ── 5. Cancel a run ────────────────────────────────────────────
	if err := client.CancelWorkflowRun(ctx, repo, owner, "99999"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("run cancelled")

	// ── 6. Re-run a workflow ───────────────────────────────────────
	if err := client.ReRunWorkflow(ctx, repo, owner, "99999"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("run re-triggered")

	// ── 7. List jobs in a run ──────────────────────────────────────
	jobs, err := client.ListJobsForWorkflowRun(ctx, repo, owner, "99999")
	if err != nil {
		log.Fatal(err)
	}
	for _, j := range jobs {
		fmt.Printf("job: %s status=%s result=%s runner=%s\n",
			j.Name, j.Status, j.Result, j.Runner)
	}

	// ── 8. List artifacts for a run ────────────────────────────────
	artifacts, err := client.ListArtifacts(ctx, repo, owner, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range artifacts {
		fmt.Printf("artifact: %s size=%d expired=%v\n", a.Name, a.Size, a.Expired)
	}

	// ── 9. Download an artifact ────────────────────────────────────
	if len(artifacts) > 0 {
		data, err := client.DownloadArtifact(ctx, artifacts[0].DownloadURL)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("downloaded %d bytes\n", len(data))
		// data is a raw zip — write to file, extract, etc.
		if err := os.WriteFile("artifact.zip", data, 0644); err != nil {
			log.Fatal(err)
		}
	}

	// ── 10. Get logs for a run ─────────────────────────────────────
	logs, err := client.GetWorkflowRunLogs(ctx, repo, owner, "99999")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("logs for run %s:\n%s\n", logs.RunID, logs.Content)

	// ── Status / Result enums ──────────────────────────────────────
	_ = models.RunStatusRunning
	_ = models.RunResultSuccess
	_ = models.TriggerPush
	_ = models.ArtifactArchive

	// ── Error handling patterns ─────────────────────────────────────
	_, err = client.GetWorkflow(ctx, repo, owner, "bad-id")
	switch {
	case errors.Is(err, ci.ErrUnauthorized):
		fmt.Println("check your token")
	case errors.Is(err, ci.ErrNotFound):
		fmt.Println("resource does not exist")
	case errors.Is(err, ci.ErrRateLimited):
		fmt.Println("slow down")
	case err != nil:
		fmt.Println("unexpected:", err)
	}
}
