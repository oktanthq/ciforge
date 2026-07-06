package models

type RunStatus string

const (
	RunStatusQueued    RunStatus = "queued"
	RunStatusPending   RunStatus = "pending"
	RunStatusRunning   RunStatus = "running"
	RunStatusCompleted RunStatus = "completed"
	RunStatusCancelled RunStatus = "cancelled"
	RunStatusWaiting   RunStatus = "waiting"
	RunStatusSkipped   RunStatus = "skipped"
	RunStatusUnknown   RunStatus = "unknown"
)

type RunResult string

const (
	RunResultSuccess   RunResult = "success"
	RunResultFailed    RunResult = "failed"
	RunResultCancelled RunResult = "cancelled"
	RunResultSkipped   RunResult = "skipped"
	RunResultTimedOut  RunResult = "timed_out"
	RunResultNeutral   RunResult = "neutral"
	RunResultActionReq RunResult = "action_required"
	RunResultUnknown   RunResult = "unknown"
)

type ArtifactType string

const (
	ArtifactArchive ArtifactType = "archive"
	ArtifactBinary  ArtifactType = "binary"
	ArtifactReport  ArtifactType = "report"
	ArtifactImage   ArtifactType = "image"
	ArtifactLog     ArtifactType = "log"
	ArtifactOther   ArtifactType = "other"
)

type TriggerType string

const (
	TriggerPush         TriggerType = "push"
	TriggerPullRequest  TriggerType = "pull_request"
	TriggerTag          TriggerType = "tag"
	TriggerManual       TriggerType = "manual"
	TriggerSchedule     TriggerType = "schedule"
	TriggerWebhook      TriggerType = "webhook"
	TriggerAPI          TriggerType = "api"
	TriggerRelease      TriggerType = "release"
	TriggerWorkflowCall TriggerType = "workflow_call"
	TriggerUnknown      TriggerType = "unknown"
)
