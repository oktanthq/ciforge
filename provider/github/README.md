# GitHub Provider

GitHub Actions provider for CIForge.

## Usage

```go
import (
    "github.com/abdulahshoaib/ciforge/ci"
    "github.com/abdulahshoaib/ciforge/provider/github"
)

client := ci.New(
    ci.WithProvider(github.New("ghp_your_token")),
)
```

## Architecture

```
GitHub REST API
    ↓
provider DTOs (private)
    ↓
Mapper
    ↓
models.Pipeline / models.Run / models.Job / models.Artifact
```

## Endpoints

### Workflows

| Method | GitHub API | Model |
|--------|-----------|-------|
| ListWorkflows | GET /repos/{owner}/{repo}/actions/workflows | []models.Pipeline |
| GetWorkflow | GET /repos/{owner}/{repo}/actions/workflows/{id} | *models.Pipeline |
| DisableWorkflow | PUT /repos/{owner}/{repo}/actions/workflows/{id}/disable | error |
| EnableWorkflow | PUT /repos/{owner}/{repo}/actions/workflows/{id}/enable | error |
| CreateWorkflowDispatch | POST /repos/{owner}/{repo}/actions/workflows/{id}/dispatches | models.DispatchEvent |

### Workflow Runs

| Method | GitHub API | Model |
|--------|-----------|-------|
| ListRunsForRepo | GET /repos/{owner}/{repo}/actions/runs | []models.Run |
| GetWorkflowRun | GET /repos/{owner}/{repo}/actions/runs/{id} | models.Run |
| DeleteWorkflowRun | DELETE /repos/{owner}/{repo}/actions/runs/{id} | error |
| GetWorkflowRunReviewHistory | GET /repos/{owner}/{repo}/actions/runs/{id}/approvals | []workflowApprovalDTO |
| GetWorkflowRunAttempt | GET /repos/{owner}/{repo}/actions/runs/{id}/attempts/{attempt} | models.Run |
| CancelWorkflowRun | POST /repos/{owner}/{repo}/actions/runs/{id}/cancel | error |
| ForceCancelWorkflowRun | POST /repos/{owner}/{repo}/actions/runs/{id}/force-cancel | error |
| DownloadWorkflowRunLogs | GET /repos/{owner}/{repo}/actions/runs/{id}/logs | error |
| DeleteWorkflowRunLogs | DELETE /repos/{owner}/{repo}/actions/runs/{id}/logs | error |
| GetPendingDeployments | GET /repos/{owner}/{repo}/actions/runs/{id}/pending_deployments | error |
| ReRunFailedJobs | POST /repos/{owner}/{repo}/actions/runs/{id}/rerun-failed-jobs | error |
| ReRunWorkflow | POST /repos/{owner}/{repo}/actions/runs/{id}/rerun | error |
| ListRunsForWorkflow | GET /repos/{owner}/{repo}/actions/workflows/{id}/runs | []models.Run |

### Workflow Jobs

| Method | GitHub API | Model |
|--------|-----------|-------|
| GetWorkflowJob | GET /repos/{owner}/{repo}/actions/jobs/{id} | []models.Job |
| DownloadWorkflowJobLogs | GET /repos/{owner}/{repo}/actions/jobs/{id}/logs | error |
| ListJobsForWorkflowRun | GET /repos/{owner}/{repo}/actions/runs/{id}/jobs | []models.Job |
| ListJobsForWorkflowRunAttempt | GET /repos/{owner}/{repo}/actions/runs/{id}/attempts/{attempt}/jobs | []models.Job |

### Artifacts

| Method | GitHub API | Model |
|--------|-----------|-------|
| ListArtifacts | GET /repos/{owner}/{repo}/actions/artifacts | []models.Artifact |
| GetArtifact | GET /repos/{owner}/{repo}/actions/artifacts/{id} | models.Artifact |
| DownloadArtifact | GET {download_url} | []byte |
| DeleteArtifact | DELETE /repos/{owner}/{repo}/actions/artifacts/{id} | error |
| ListWorkflowRunArtifacts | GET /repos/{owner}/{repo}/actions/runs/{id}/artifacts | []models.Artifact |

### Logs

| Method | GitHub API | Model |
|--------|-----------|-------|
| GetWorkflowRunLogs | GET /repos/{owner}/{repo}/actions/runs/{id}/logs | *models.Log |

No GitHub-specific types escape this package.
