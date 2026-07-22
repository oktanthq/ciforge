# GitHub Actions REST API

## Workflows API

```http
GET    /repos/{owner}/{repo}/actions/workflows
GET    /repos/{owner}/{repo}/actions/workflows/{workflow_id}
PUT    /repos/{owner}/{repo}/actions/workflows/{workflow_id}/disable
POST   /repos/{owner}/{repo}/actions/workflows/{workflow_id}/dispatches
PUT    /repos/{owner}/{repo}/actions/workflows/{workflow_id}/enable
GET    /repos/{owner}/{repo}/actions/workflows/{workflow_id}/timing
```

### Functions

| Function                | Method | Endpoint                                                                 |
| ----------------------- | ------ | ------------------------------------------------------------------------ |
| ListWorkflows           | GET    | /repos/{owner}/{repo}/actions/workflows                                  |
| GetWorkflow             | GET    | /repos/{owner}/{repo}/actions/workflows/{workflow_id}                    |
| DisableWorkflow         | PUT    | /repos/{owner}/{repo}/actions/workflows/{workflow_id}/disable            |
| CreateWorkflowDispatch  | POST   | /repos/{owner}/{repo}/actions/workflows/{workflow_id}/dispatches          |
| EnableWorkflow          | PUT    | /repos/{owner}/{repo}/actions/workflows/{workflow_id}/enable             |
| GetWorkflowUsage        | GET    | /repos/{owner}/{repo}/actions/workflows/{workflow_id}/timing             |

## Runs API

```http
GET    /repos/{owner}/{repo}/actions/runs
GET    /repos/{owner}/{repo}/actions/runs/{run_id}
DELETE /repos/{owner}/{repo}/actions/runs/{run_id}
GET    /repos/{owner}/{repo}/actions/runs/{run_id}/approvals
POST   /repos/{owner}/{repo}/actions/runs/{run_id}/approve
GET    /repos/{owner}/{repo}/actions/runs/{run_id}/attempts/{attempt_number}
GET    /repos/{owner}/{repo}/actions/runs/{run_id}/attempts/{attempt_number}/logs
POST   /repos/{owner}/{repo}/actions/runs/{run_id}/cancel
POST   /repos/{owner}/{repo}/actions/runs/{run_id}/force-cancel
GET    /repos/{owner}/{repo}/actions/runs/{run_id}/logs
DELETE /repos/{owner}/{repo}/actions/runs/{run_id}/logs
GET    /repos/{owner}/{repo}/actions/runs/{run_id}/pending_deployments
POST   /repos/{owner}/{repo}/actions/runs/{run_id}/pending_deployments
POST   /repos/{owner}/{repo}/actions/runs/{run_id}/rerun
POST   /repos/{owner}/{repo}/actions/runs/{run_id}/rerun-failed-jobs
GET    /repos/{owner}/{repo}/actions/runs/{run_id}/timing
GET    /repos/{owner}/{repo}/actions/workflows/{workflow_id}/runs
```

### Functions

| Function                        | Method | Endpoint                                                                   |
| ------------------------------- | ------ | -------------------------------------------------------------------------- |
| ListRunsForRepo                 | GET    | /repos/{owner}/{repo}/actions/runs                                        |
| GetWorkflowRun                  | GET    | /repos/{owner}/{repo}/actions/runs/{run_id}                                |
| DeleteWorkflowRun               | DELETE | /repos/{owner}/{repo}/actions/runs/{run_id}                                |
| GetWorkflowRunReviewHistory     | GET    | /repos/{owner}/{repo}/actions/runs/{run_id}/approvals                      |
| ApproveWorflowRunForForkPR      | POST   | /repos/{owner}/{repo}/actions/runs/{run_id}/approve                        |
| GetWorkflowRunAttempt           | GET    | /repos/{owner}/{repo}/actions/runs/{run_id}/attempts/{attempt_number}       |
| DownloadWorkflowRunAttemptLogs  | GET    | /repos/{owner}/{repo}/actions/runs/{run_id}/attempts/{attempt_number}/logs  |
| CancelWorkflowRun               | POST   | /repos/{owner}/{repo}/actions/runs/{run_id}/cancel                         |
| ForceCancelWorkflowRun          | POST   | /repos/{owner}/{repo}/actions/runs/{run_id}/force-cancel                   |
| DownloadWorkflowRunLogs         | GET    | /repos/{owner}/{repo}/actions/runs/{run_id}/logs                           |
| DeleteWorkflowRunLogs           | DELETE | /repos/{owner}/{repo}/actions/runs/{run_id}/logs                           |
| GetPendingDeployments           | GET    | /repos/{owner}/{repo}/actions/runs/{run_id}/pending_deployments             |
| ReviewPendingDeployments        | POST   | /repos/{owner}/{repo}/actions/runs/{run_id}/pending_deployments             |
| ReRunWorkflow                   | POST   | /repos/{owner}/{repo}/actions/runs/{run_id}/rerun                          |
| ReRunFailedJobs                 | POST   | /repos/{owner}/{repo}/actions/runs/{run_id}/rerun-failed-jobs              |
| GetWorkflowRunUsage             | GET    | /repos/{owner}/{repo}/actions/runs/{run_id}/timing                         |
| ListRunsForWorkflow             | GET    | /repos/{owner}/{repo}/actions/workflows/{workflow_id}/runs                 |

## Jobs API

```http
GET    /repos/{owner}/{repo}/actions/jobs/{job_id}
GET    /repos/{owner}/{repo}/actions/jobs/{job_id}/logs
GET    /repos/{owner}/{repo}/actions/runs/{run_id}/jobs
GET    /repos/{owner}/{repo}/actions/runs/{run_id}/attempts/{attempt_number}/jobs
```

### Functions

| Function                        | Method | Endpoint                                                                   |
| ------------------------------- | ------ | -------------------------------------------------------------------------- |
| GetWorkflowJob                  | GET    | /repos/{owner}/{repo}/actions/jobs/{job_id}                                |
| DownloadWorkflowJobLogs         | GET    | /repos/{owner}/{repo}/actions/jobs/{job_id}/logs                           |
| ListJobsForWorkflowRun          | GET    | /repos/{owner}/{repo}/actions/runs/{run_id}/jobs                           |
| ListJobsForWorkflowRunAttempt   | GET    | /repos/{owner}/{repo}/actions/runs/{run_id}/attempts/{attempt_number}/jobs  |

## Artifacts API

```http
GET    /repos/{owner}/{repo}/actions/artifacts
GET    /repos/{owner}/{repo}/actions/artifacts/{artifact_id}
DELETE /repos/{owner}/{repo}/actions/artifacts/{artifact_id}
GET    /repos/{owner}/{repo}/actions/runs/{run_id}/artifacts
GET    {download_url}
```

### Functions

| Function                   | Method | Endpoint                                                        |
| -------------------------- | ------ | --------------------------------------------------------------- |
| ListArtifacts              | GET    | /repos/{owner}/{repo}/actions/artifacts                         |
| GetArtifact                | GET    | /repos/{owner}/{repo}/actions/artifacts/{artifact_id}           |
| DeleteArtifact             | DELETE | /repos/{owner}/{repo}/actions/artifacts/{artifact_id}           |
| ListWorkflowRunArtifacts   | GET    | /repos/{owner}/{repo}/actions/runs/{run_id}/artifacts            |
| DownloadArtifact            | GET    | {download_url}                                                  |

## Logs API

```http
GET    /repos/{owner}/{repo}/actions/runs/{run_id}/logs
```

### Functions

| Function              | Method | Endpoint                                               |
| --------------------- | ------ | ------------------------------------------------------ |
| GetWorkflowRunLogs    | GET    | /repos/{owner}/{repo}/actions/runs/{run_id}/logs       |

## Concurrency Groups API

```http
GET    /repos/{owner}/{repo}/actions/concurrency_groups
GET    /repos/{owner}/{repo}/actions/concurrency_groups/{concurrency_group_name}
GET    /repos/{owner}/{repo}/actions/runs/{run_id}/concurrency_groups
```

### Functions

| Function                          | Method | Endpoint                                                                 |
| --------------------------------- | ------ | ------------------------------------------------------------------------ |
| ListConcurrencyGroupsForRepo      | GET    | /repos/{owner}/{repo}/actions/concurrency_groups                         |
| GetConcurrencyGroupsForRepo        | GET    | /repos/{owner}/{repo}/actions/concurrency_groups/{concurrency_group_name} |
| ListConcurrencyGroupsForWorflow   | GET    | /repos/{owner}/{repo}/actions/runs/{run_id}/concurrency_groups           |
