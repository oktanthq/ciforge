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

| Method | GitHub API | Model |
|--------|-----------|-------|
| ListPipelines | GET /repos/{owner}/{repo}/actions/workflows | []models.Pipeline |
| GetPipeline | GET /repos/{owner}/{repo}/actions/workflows/{id} | *models.Pipeline |
| ListRuns | GET /repos/{owner}/{repo}/actions/workflows/{id}/runs | []models.Run |
| GetRuns | GET /repos/{owner}/{repo}/actions/runs/{id} | *models.Run |
| CancelRun | POST /repos/{owner}/{repo}/actions/runs/{id}/cancel | error |
| RetryRun | POST /repos/{owner}/{repo}/actions/runs/{id}/rerun | error |
| ListJobs | GET /repos/{owner}/{repo}/actions/runs/{id}/jobs | []models.Job |
| ListArtifacts | GET /repos/{owner}/{repo}/actions/runs/{id}/artifacts | []models.Artifact |
| DownloadArtifact | GET /repos/{owner}/{repo}/actions/artifacts/{id}/zip | []byte |
| GetLogs | GET /repos/{owner}/{repo}/actions/runs/{id}/logs | *models.Log |

No GitHub-specific types escape this package.
