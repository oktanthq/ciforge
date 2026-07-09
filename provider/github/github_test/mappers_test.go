package github_test

import (
	"testing"

	"github.com/abdulahshoaib/ciforge/models"
	"github.com/abdulahshoaib/ciforge/provider/github"
)

func TestMapperRunStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  models.RunStatus
	}{
		{"queued", models.RunStatusQueued},
		{"in_progress", models.RunStatusRunning},
		{"completed", models.RunStatusCompleted},
		{"waiting", models.RunStatusWaiting},
		{"pending", models.RunStatusPending},
		{"cancelled", models.RunStatusCancelled},
		{"unknown_status", models.RunStatusUnknown},
	}

	for _, tc := range tests {
		got := github.MapRunStatus(tc.input)
		if got != tc.want {
			t.Errorf("MapRunStatus(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

func TestMapperRunResult(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  models.RunResult
	}{
		{"success", models.RunResultSuccess},
		{"failure", models.RunResultFailed},
		{"cancelled", models.RunResultCancelled},
		{"skipped", models.RunResultSkipped},
		{"timed_out", models.RunResultTimedOut},
		{"neutral", models.RunResultNeutral},
		{"action_required", models.RunResultActionReq},
		{"unknown_result", models.RunResultUnknown},
	}

	for _, tc := range tests {
		got := github.MapRunResult(tc.input)
		if got != tc.want {
			t.Errorf("MapRunResult(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

func TestMapperTriggerType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  models.TriggerType
	}{
		{"push", models.TriggerPush},
		{"pull_request", models.TriggerPullRequest},
		{"release", models.TriggerRelease},
		{"schedule", models.TriggerSchedule},
		{"workflow_dispatch", models.TriggerManual},
		{"repository_dispatch", models.TriggerWebhook},
		{"workflow_call", models.TriggerWorkflowCall},
		{"unknown_event", models.TriggerUnknown},
	}

	for _, tc := range tests {
		got := github.MapTriggerType(tc.input)
		if got != tc.want {
			t.Errorf("MapTriggerType(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}
