package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	gh "github.com/abdulahshoaib/ciforge/provider/github"
)

// Demonstrates GitHub Actions API operations using the ciforge GitHub provider.
func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("set GITHUB_TOKEN env var")
	}

	// ── Create client ───────────────────────────────────────────────
	// gh.New creates a new GitHub API client with the provided token.
	client := gh.New(token)

	// gh.New with options configures a custom HTTP client and base URL for GitHub Enterprise.
	_ = gh.New(token,
		gh.WithHTTPClient(&http.Client{Timeout: 30 * time.Second}),
		gh.WithBaseURL("https://github.example.com/api/v3"),
	)

	ctx := context.Background()
	owner := "owner"
	repo := "repo"

	// ── 1. List workflows ──────────────────────────────────────────
	// ListWorkflows returns all GitHub Actions workflows for a repository.
	workflows, err := client.ListWorkflows(ctx, repo, owner)
	if err != nil {
		log.Fatal(err)
	}
	for _, w := range workflows {
		fmt.Printf("workflow: %s (id=%d, disabled=%v)\n", w.Name, w.ID, w.State != "active")
	}

	// ── 2. Get a single workflow ───────────────────────────────────
	// GetWorkflow retrieves a specific workflow by its ID.
	workflow, err := client.GetWorkflow(ctx, repo, owner, 12345)
	if err != nil {
		if errors.Is(err, gh.ErrNotFound) {
			fmt.Println("workflow not found")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("workflow: %s, url: %s\n", workflow.Name, workflow.HTMLURL)
	}

	// ── 3. List runs for a workflow ─────────────────────────────────
	// ListRunsForWorkflow returns all runs for a specific workflow with optional filtering.
	runs, err := client.ListRunsForWorkflow(ctx, repo, owner, 12345, gh.ListRunsForWorkflowOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range runs {
		fmt.Printf("run #%d: status=%s result=%s branch=%s sha=%s\n",
			r.RunNumber, r.Status, r.Conclusion, r.HeadBranch, r.HeadSHA)
	}

	// ── 4. Get a single run ────────────────────────────────────────
	// GetWorkflowRun retrieves a specific workflow run by its ID.
	run, err := client.GetWorkflowRun(ctx, repo, owner, 99999)
	if err != nil {
		if errors.Is(err, gh.ErrUnauthorized) {
			fmt.Println("bad token")
		} else if errors.Is(err, gh.ErrRateLimited) {
			fmt.Println("rate limited — back off")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("run #%d: %s\n", run.RunNumber, run.Conclusion)
	}

	// ── 5. Cancel a run ────────────────────────────────────────────
	// CancelWorkflowRun cancels a running workflow by its ID.
	if err := client.CancelWorkflowRun(ctx, repo, owner, 99999); err != nil {
		log.Fatal(err)
	}
	fmt.Println("run cancelled")

	// ── 6. Re-run a workflow ───────────────────────────────────────
	// ReRunWorkflow re-runs all jobs in a workflow run.
	if err := client.ReRunWorkflow(ctx, repo, owner, 99999); err != nil {
		log.Fatal(err)
	}
	fmt.Println("run re-triggered")

	// ── 7. Re-run failed jobs ──────────────────────────────────────
	// ReRunFailedJobs re-runs only the failed jobs in a workflow run.
	if err := client.ReRunFailedJobs(ctx, repo, owner, 99999); err != nil {
		log.Fatal(err)
	}
	fmt.Println("failed jobs re-triggered")

	// ── 8. Force cancel a run ──────────────────────────────────────
	// ForceCancelWorkflowRun forcefully cancels a workflow run, terminating any running jobs.
	if err := client.ForceCancelWorkflowRun(ctx, repo, owner, 99999); err != nil {
		log.Fatal(err)
	}
	fmt.Println("run force cancelled")

	// ── 9. List jobs in a run ──────────────────────────────────────
	// ListJobsForWorkflowRun returns all jobs for a specific workflow run.
	jobs, err := client.ListJobsForWorkflowRun(ctx, repo, owner, 99999)
	if err != nil {
		log.Fatal(err)
	}
	for _, j := range jobs {
		fmt.Printf("job: %s status=%s result=%s runner=%s\n",
			j.Name, j.Status, j.Conclusion, j.RunnerName)
	}

	// ── 10. List jobs for run attempt ──────────────────────────────
	// ListJobsForWorkflowRunAttempt returns jobs for a specific attempt of a workflow run.
	attemptJobs, err := client.ListJobsForWorkflowRunAttempt(ctx, repo, owner, 99999, 1)
	if err != nil {
		log.Fatal(err)
	}
	for _, j := range attemptJobs {
		fmt.Printf("attempt job: %s status=%s\n", j.Name, j.Status)
	}

	// ── 11. List artifacts for a repo ──────────────────────────────
	// ListArtifacts returns all artifacts for a repository with optional pagination.
	artifacts, err := client.ListArtifacts(ctx, repo, owner, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range artifacts {
		fmt.Printf("artifact: %s size=%d expired=%v\n", a.Name, a.SizeInBytes, a.Expired)
	}

	// ── 12. List artifacts for a run ───────────────────────────────
	// ListWorkflowRunArtifacts returns artifacts for a specific workflow run.
	runArtifacts, err := client.ListWorkflowRunArtifacts(ctx, repo, owner, 99999, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range runArtifacts {
		fmt.Printf("run artifact: %s size=%d\n", a.Name, a.SizeInBytes)
	}

	// ── 13. Download an artifact ────────────────────────────────────
	// DownloadArtifact downloads an artifact archive from its download URL.
	if len(artifacts) > 0 {
		data, err := client.DownloadArtifact(ctx, artifacts[0].ArchiveDownloadURL)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("downloaded %d bytes\n", len(data))
		if err := os.WriteFile("artifact.zip", data, 0644); err != nil {
			log.Fatal(err)
		}
	}

	// ── 14. Get logs for a run ─────────────────────────────────────
	// GetWorkflowRunLogs retrieves the log output for a workflow run.
	logs, err := client.GetWorkflowRunLogs(ctx, repo, owner, 99999)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("logs:\n%s\n", logs)

	// ── 15. Download attempt logs ──────────────────────────────────
	// DownloadWorkflowRunAttemptLogs downloads logs for a specific run attempt.
	attemptLogs, err := client.DownloadWorkflowRunAttemptLogs(ctx, repo, owner, 99999, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("attempt logs: %d bytes\n", len(attemptLogs))

	// ── 16. Delete run logs ────────────────────────────────────────
	// DeleteWorkflowRunLogs removes stored logs for a workflow run.
	if err := client.DeleteWorkflowRunLogs(ctx, repo, owner, 99999); err != nil {
		log.Fatal(err)
	}
	fmt.Println("logs deleted")

	// ── 17. Delete a workflow run ──────────────────────────────────
	// DeleteWorkflowRun permanently deletes a workflow run and its associated data.
	if err := client.DeleteWorkflowRun(ctx, repo, owner, 99999); err != nil {
		log.Fatal(err)
	}
	fmt.Println("run deleted")

	// ── 18. Get workflow usage ─────────────────────────────────────
	// GetWorkflowUsage returns billable usage information for a workflow.
	usage, err := client.GetWorkflowUsage(ctx, repo, owner, 12345)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("workflow usage: ubuntu=%v macos=%v windows=%v\n",
		usage.Ubuntu, usage.MacOS, usage.Windows)

	// ── 19. Get run usage ──────────────────────────────────────────
	// GetWorkflowRunUsage returns billable minutes and storage for a run.
	runUsage, err := client.GetWorkflowRunUsage(ctx, repo, owner, 99999)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("run usage: ubuntu=%v macos=%v windows=%v\n",
		runUsage.Ubuntu, runUsage.MacOS, runUsage.Windows)

	// ── Error handling patterns ─────────────────────────────────────
	// errors.Is checks for specific API errors like unauthorized, not found, or rate limited.
	_, err = client.GetWorkflow(ctx, repo, owner, 0)
	switch {
	case errors.Is(err, gh.ErrUnauthorized):
		fmt.Println("check your token")
	case errors.Is(err, gh.ErrNotFound):
		fmt.Println("resource does not exist")
	case errors.Is(err, gh.ErrRateLimited):
		fmt.Println("slow down")
	case err != nil:
		fmt.Println("unexpected:", err)
	}
}
