package github

import (
	"fmt"
	"time"

	"github.com/abdulahshoaib/ciforge/models"
)

func mapWorkflow(w workflowDTO) models.Pipeline {
	return models.Pipeline{
		ID:        fmt.Sprintf("%d", w.ID),
		Name:      w.Name,
		URL:       w.HTMLURL,
		CreatedAt: w.CreatedAt,
		UpdatedAt: w.UpdatedAt,
		Disabled:  w.State != "active",
	}
}

func mapWorkflows(dtos []workflowDTO) []models.Pipeline {
	pipelines := make([]models.Pipeline, len(dtos))
	for i, dto := range dtos {
		pipelines[i] = mapWorkflow(dto)
	}
	return pipelines
}

func mapWorkflowRun(r workflowRunDTO) models.Run {
	run := models.Run{
		ID:         fmt.Sprintf("%d", r.ID),
		PipelineID: fmt.Sprintf("%d", r.WorkflowID),
		Number:     r.RunNumber,
		Status:     MapRunStatus(r.Status),
		Result:     MapRunResult(r.Conclusion),
		Trigger:    MapTriggerType(r.Event),
		Branch:     r.HeadBranch,
		CommitSHA:  r.HeadSHA,
		Message:    r.DisplayTitle,
		URL:        r.HTMLURL,
		StartedAt:  &r.CreatedAt,
		FinishedAt: &r.UpdatedAt,
	}

	if run.StartedAt != nil && run.FinishedAt != nil {
		run.Duration = run.FinishedAt.Sub(*run.StartedAt)
	}

	return run
}

func mapWorkflowRuns(dtos []workflowRunDTO) []models.Run {
	runs := make([]models.Run, len(dtos))
	for i, dto := range dtos {
		runs[i] = mapWorkflowRun(dto)
	}
	return runs
}

func mapJob(j jobDTO) models.Job {
	job := models.Job{
		ID:         fmt.Sprintf("%d", j.ID),
		RunID:      fmt.Sprintf("%d", j.RunID),
		Name:       j.Name,
		Status:     MapRunStatus(j.Status),
		Result:     MapRunResult(j.Conclusion),
		URL:        j.HTMLURL,
		StartedAt:  j.StartedAt,
		FinishedAt: j.CompletedAt,
	}

	if job.StartedAt != nil && job.FinishedAt != nil {
		job.Duration = job.FinishedAt.Sub(*job.StartedAt)
	}

	return job
}

func mapJobs(dtos []jobDTO) []models.Job {
	jobs := make([]models.Job, len(dtos))
	for i, dto := range dtos {
		jobs[i] = mapJob(dto)
	}
	return jobs
}

func mapArtifact(dto artifactDTO) models.Artifact {
	var runID int64
	if dto.WorkflowRun != nil {
		runID = dto.WorkflowRun.ID
	}

	return models.Artifact{
		ID:          dto.ID,
		RunID:       runID,
		Name:        dto.Name,
		Type:        models.ArtifactArchive,
		Size:        dto.SizeInBytes,
		Expired:     dto.Expired,
		CreatedAt:   dto.CreatedAt,
		ExpiresAt:   dto.ExpiresAt,
		DownloadURL: dto.ArchiveDownloadURL,
	}
}

func mapArtifacts(dtos []artifactDTO) []models.Artifact {
	artifacts := make([]models.Artifact, len(dtos))
	for i, dto := range dtos {
		artifacts[i] = mapArtifact(dto)
	}
	return artifacts
}

func MapRunStatus(status string) models.RunStatus {
	switch status {
	case "queued":
		return models.RunStatusQueued
	case "in_progress":
		return models.RunStatusRunning
	case "completed":
		return models.RunStatusCompleted
	case "waiting":
		return models.RunStatusWaiting
	case "pending":
		return models.RunStatusPending
	case "cancelled":
		return models.RunStatusCancelled
	default:
		return models.RunStatusUnknown
	}
}

func MapRunResult(conclusion string) models.RunResult {
	switch conclusion {
	case "success":
		return models.RunResultSuccess
	case "failure":
		return models.RunResultFailed
	case "cancelled":
		return models.RunResultCancelled
	case "skipped":
		return models.RunResultSkipped
	case "timed_out":
		return models.RunResultTimedOut
	case "neutral":
		return models.RunResultNeutral
	case "action_required":
		return models.RunResultActionReq
	default:
		return models.RunResultUnknown
	}
}

func MapTriggerType(event string) models.TriggerType {
	switch event {
	case "push":
		return models.TriggerPush
	case "pull_request":
		return models.TriggerPullRequest
	case "release":
		return models.TriggerRelease
	case "schedule":
		return models.TriggerSchedule
	case "workflow_dispatch":
		return models.TriggerManual
	case "repository_dispatch":
		return models.TriggerWebhook
	case "workflow_call":
		return models.TriggerWorkflowCall
	default:
		return models.TriggerUnknown
	}
}

func mapDuration(startedAt, finishedAt *time.Time) time.Duration {
	if startedAt == nil || finishedAt == nil {
		return 0
	}
	return finishedAt.Sub(*startedAt)
}

func mapDispatchEvent(dto workflowDispatchResponse) models.DispatchEvent {
	return models.DispatchEvent{
		RunID:  dto.WorkflowRunID,
		RunURL: dto.RunURL,
		URL:    dto.HTMLURL,
	}
}
