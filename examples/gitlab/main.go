package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	gl "github.com/abdulahshoaib/ciforge/provider/gitlab"
)

// Demonstrates GitLab CI/CD API operations using the ciforge GitLab provider.
func main() {
	token := os.Getenv("GITLAB_TOKEN")
	if token == "" {
		log.Fatal("set GITLAB_TOKEN env var")
	}

	// ── Create client ───────────────────────────────────────────────
	// gl.New creates a new GitLab API client with the provided token.
	client := gl.New(token)

	// gl.New with options configures a custom HTTP client and base URL for self-hosted GitLab instances.
	_ = gl.New(token,
		gl.WithHTTPClient(&http.Client{Timeout: 30 * time.Second}),
		gl.WithBaseURL("https://gitlab.example.com/api/v4"),
	)

	ctx := context.Background()
	projectID := 12345

	// ── 1. List pipelines ──────────────────────────────────────────
	// ListPipelines returns all pipelines for a project with optional filtering.
	pipelines, err := client.ListPipelines(ctx, projectID, &gl.ListPipelinesOptions{Ref: "main"})
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range pipelines {
		fmt.Printf("pipeline #%d: status=%s ref=%s\n", p.ID, p.Status, p.Ref)
	}

	// ── 2. Get a single pipeline ───────────────────────────────────
	// GetPipeline retrieves a specific pipeline by its ID with full details.
	pipeline, err := client.GetPipeline(ctx, projectID, 1)
	if err != nil {
		if errors.Is(err, gl.ErrNotFound) {
			fmt.Println("pipeline not found")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("pipeline #%d: status=%s\n", pipeline.ID, pipeline.Status)
	}

	// ── 3. Create a pipeline ───────────────────────────────────────
	// CreatePipeline triggers a new pipeline for a project.
	created, err := client.CreatePipeline(ctx, projectID, "main")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created pipeline #%d\n", created.ID)

	// ── 4. Cancel a pipeline ───────────────────────────────────────
	// CancelPipeline cancels all jobs in a pipeline.
	cancelled, err := client.CancelPipeline(ctx, projectID, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("pipeline #%d cancelled\n", cancelled.ID)

	// ── 5. List pipeline jobs ──────────────────────────────────────
	// ListJobsForPipeline returns all jobs for a specific pipeline.
	jobs, err := client.ListJobsForPipeline(ctx, projectID, 1, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, j := range jobs {
		fmt.Printf("job: %s status=%s stage=%s\n", j.Name, j.Status, j.Stage)
	}

	// ── 6. Get a single job ────────────────────────────────────────
	// GetJob retrieves a specific job by its ID.
	job, err := client.GetJob(ctx, projectID, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("job: %s runner=%d\n", job.Name, job.Runner.ID)

	// ── 7. Retry a job ─────────────────────────────────────────────
	// RetryJob retries a failed job.
	retried, err := client.RetryJob(ctx, projectID, 1, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retried job: %s status=%s\n", retried.Name, retried.Status)

	// ── 8. Download job artifacts ───────────────────────────────────
	// DownloadJobArtifacts downloads the artifact archive for a completed job.
	data, err := client.DownloadJobArtifacts(ctx, projectID, 1, "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("downloaded %d bytes of artifacts\n", len(data))

	// ── 9. List pipeline schedules ─────────────────────────────────
	// ListPipelineSchedules returns all schedules for a project.
	schedules, err := client.ListPipelineSchedules(ctx, projectID, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range schedules {
		fmt.Printf("schedule: %s ref=%s next_run=%s\n", s.Description, s.Ref, s.NextRunAt)
	}

	// ── 10. Create a pipeline schedule ─────────────────────────────
	// CreatePipelineSchedule creates a new schedule for recurring pipeline runs.
	schedule, err := client.CreatePipelineSchedule(ctx, projectID, &gl.CreatePipelineScheduleBody{
		Description: "nightly build",
		Ref:         "main",
		Cron:        "0 2 * * *",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created schedule: %s (id=%d)\n", schedule.Description, schedule.ID)

	// ── 11. Get runner details ─────────────────────────────────────
	// GetRunner retrieves details about a specific runner.
	runner, err := client.GetRunner(ctx, 1)
	if err != nil {
		if errors.Is(err, gl.ErrUnauthorized) {
			fmt.Println("check your token")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("runner: %s online=%v\n", runner.Description, runner.Online)
	}

	// ── 12. List runner jobs ───────────────────────────────────────
	// ListRunnerJobs returns jobs running on a specific runner.
	runnerJobs, err := client.ListRunnerJobs(ctx, 1, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, j := range runnerJobs {
		fmt.Printf("runner job: %s status=%s\n", j.Name, j.Status)
	}

	// ── Error handling patterns ─────────────────────────────────────
	// errors.Is checks for specific API errors like unauthorized, not found, or rate limited.
	_, err = client.GetPipeline(ctx, projectID, 0)
	switch {
	case errors.Is(err, gl.ErrUnauthorized):
		fmt.Println("check your token")
	case errors.Is(err, gl.ErrNotFound):
		fmt.Println("resource does not exist")
	case errors.Is(err, gl.ErrRateLimited):
		fmt.Println("slow down")
	case err != nil:
		fmt.Println("unexpected:", err)
	}
}
