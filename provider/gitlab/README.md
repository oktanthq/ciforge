# GitLab REST API

## Jobs API

```http
GET  /projects/:id/jobs
GET  /projects/:id/pipelines/:pipeline_id/jobs
GET  /projects/:id/pipelines/:pipeline_id/trigger_jobs
GET  /job
GET  /job/allowed_agents
GET  /projects/:id/jobs/:job_id
GET  /projects/:id/jobs/:job_id/trace
POST /projects/:id/jobs/:job_id/cancel
POST /projects/:id/jobs/:job_id/retry
POST /projects/:id/jobs/:job_id/erase
POST /projects/:id/jobs/:job_id/play
```

### Functions

| Function                   | Method | Endpoint                                          |
| -------------------------- | ------ | ------------------------------------------------- |
| ListJobsForProject         | GET    | /projects/:id/jobs                                |
| ListJobsForPipeline        | GET    | /projects/:id/pipelines/:pipeline_id/jobs         |
| ListTriggerJobsForPipeline | GET    | /projects/:id/pipelines/:pipeline_id/trigger_jobs |
| GetJobByToken              | GET    | /job                                              |
| GetAllowedAgents           | GET    | /job/allowed_agents                               |
| GetJob                     | GET    | /projects/:id/jobs/:job_id                        |
| GetJobTrace                | GET    | /projects/:id/jobs/:job_id/trace                  |
| CancelJob                  | POST   | /projects/:id/jobs/:job_id/cancel                 |
| RetryJob                   | POST   | /projects/:id/jobs/:job_id/retry                  |
| EraseJob                   | POST   | /projects/:id/jobs/:job_id/erase                  |
| PlayJob                    | POST   | /projects/:id/jobs/:job_id/play                   |

## Pipelines API

```http
GET    /projects/:id/pipelines
GET    /projects/:id/pipelines/:pipeline_id
GET    /projects/:id/pipelines/latest
GET    /projects/:id/pipelines/:pipeline_id/variables
GET    /projects/:id/pipelines/:pipeline_id/test_report
GET    /projects/:id/pipelines/:pipeline_id/test_report_summary
POST   /projects/:id/pipeline?ref=...
POST   /projects/:id/pipelines/:pipeline_id/retry
POST   /projects/:id/pipelines/:pipeline_id/cancel
DELETE /projects/:id/pipelines/:pipeline_id
PUT    /projects/:id/pipelines/:pipeline_id/metadata
```

### Functions

| Function                     | Method | Endpoint                                                 |
| ---------------------------- | ------ | -------------------------------------------------------- |
| ListPipelines                | GET    | /projects/:id/pipelines                                  |
| GetPipeline                  | GET    | /projects/:id/pipelines/:pipeline_id                     |
| GetLatestPipeline            | GET    | /projects/:id/pipelines/latest                           |
| GetPipelineVariables         | GET    | /projects/:id/pipelines/:pipeline_id/variables           |
| GetPipelineTestReport        | GET    | /projects/:id/pipelines/:pipeline_id/test_report         |
| GetPipelineTestReportSummary | GET    | /projects/:id/pipelines/:pipeline_id/test_report_summary |
| CreatePipeline               | POST   | /projects/:id/pipeline?ref=...                           |
| RetryPipeline                | POST   | /projects/:id/pipelines/:pipeline_id/retry               |
| CancelPipeline               | POST   | /projects/:id/pipelines/:pipeline_id/cancel              |
| DeletePipeline               | DELETE | /projects/:id/pipelines/:pipeline_id                     |
| UpdatePipelineMetadata       | PUT    | /projects/:id/pipelines/:pipeline_id/metadata            |

## Trigger Tokens API

```http
GET    /projects/:id/triggers
GET    /projects/:id/triggers/:trigger_id
POST   /projects/:id/triggers
PUT    /projects/:id/triggers/:trigger_id
DELETE /projects/:id/triggers/:trigger_id
POST   /projects/:id/trigger/pipeline
```

### Functions

| Function          | Method | Endpoint                                  |
| ----------------- | ------ | ----------------------------------------- |
| ListTriggerTokens | GET    | /projects/:id/triggers                    |
| GetTriggerToken   | GET    | /projects/:id/triggers/:trigger_id        |
| CreateTriggerToken| POST   | /projects/:id/triggers                    |
| UpdateTriggerToken| PUT    | /projects/:id/triggers/:trigger_id        |
| DeleteTriggerToken| DELETE | /projects/:id/triggers/:trigger_id        |
| TriggerPipeline   | POST   | /projects/:id/trigger/pipeline            |

## Pipeline Schedules API

```http
GET    /projects/:id/pipeline_schedules
GET    /projects/:id/pipeline_schedules/:pipeline_schedule_id
GET    /projects/:id/pipeline_schedules/:pipeline_schedule_id/pipelines
POST   /projects/:id/pipeline_schedules
PUT    /projects/:id/pipeline_schedules/:pipeline_schedule_id
POST   /projects/:id/pipeline_schedules/:pipeline_schedule_id/take_ownership
DELETE /projects/:id/pipeline_schedules/:pipeline_schedule_id
POST   /projects/:id/pipeline_schedules/:pipeline_schedule_id/play
POST   /projects/:id/pipeline_schedules/:pipeline_schedule_id/variables
GET    /projects/:id/pipeline_schedules/:pipeline_schedule_id/variables/:key
PUT    /projects/:id/pipeline_schedules/:pipeline_schedule_id/variables/:key
DELETE /projects/:id/pipeline_schedules/:pipeline_schedule_id/variables/:key
```

### Functions

| Function                      | Method | Endpoint                                                               |
| ----------------------------- | ------ | ---------------------------------------------------------------------- |
| ListPipelineSchedules         | GET    | /projects/:id/pipeline_schedules                                       |
| GetPipelineSchedule           | GET    | /projects/:id/pipeline_schedules/:pipeline_schedule_id                 |
| ListPipelinesBySchedule       | GET    | /projects/:id/pipeline_schedules/:pipeline_schedule_id/pipelines       |
| CreatePipelineSchedule        | POST   | /projects/:id/pipeline_schedules                                       |
| UpdatePipelineSchedule        | PUT    | /projects/:id/pipeline_schedules/:pipeline_schedule_id                 |
| TakeOwnershipPipelineSchedule | POST   | /projects/:id/pipeline_schedules/:pipeline_schedule_id/take_ownership  |
| DeletePipelineSchedule        | DELETE | /projects/:id/pipeline_schedules/:pipeline_schedule_id                 |
| PlayPipelineSchedule          | POST   | /projects/:id/pipeline_schedules/:pipeline_schedule_id/play            |
| CreateScheduleVariable        | POST   | /projects/:id/pipeline_schedules/:pipeline_schedule_id/variables       |
| GetScheduleVariable           | GET    | /projects/:id/pipeline_schedules/:pipeline_schedule_id/variables/:key  |
| UpdateScheduleVariable        | PUT    | /projects/:id/pipeline_schedules/:pipeline_schedule_id/variables/:key  |
| DeleteScheduleVariable        | DELETE | /projects/:id/pipeline_schedules/:pipeline_schedule_id/variables/:key  |

## Job Artifacts API

```http
GET    /projects/:id/jobs/:job_id/artifacts
GET    /projects/:id/jobs/artifacts/:ref_name/download?job=...
GET    /projects/:id/jobs/:job_id/artifacts/*artifact_path
GET    /projects/:id/jobs/:job_id/artifacts/tree
GET    /projects/:id/jobs/artifacts/:ref_name/raw/*artifact_path?job=...
POST   /projects/:id/jobs/:job_id/artifacts/keep
DELETE /projects/:id/jobs/:job_id/artifacts
DELETE /projects/:id/artifacts
```

### Functions

| Function                       | Method   | Endpoint                                                    |
| ------------------------------ | -------- | ----------------------------------------------------------- |
| DownloadJobArtifacts           | GET      | /projects/:id/jobs/:job_id/artifacts                        |
| DownloadJobArtifactsByRef      | GET      | /projects/:id/jobs/artifacts/:ref_name/download             |
| DownloadSingleArtifactFile     | GET      | /projects/:id/jobs/:job_id/artifacts/*artifact_path         |
| ListArtifactTree               | GET      | /projects/:id/jobs/:job_id/artifacts/tree                   |
| DownloadSingleArtifactFileByRef| GET      | /projects/:id/jobs/artifacts/:ref_name/raw/*artifact_path   |
| KeepJobArtifacts               | POST     | /projects/:id/jobs/:job_id/artifacts/keep                   |
| DeleteJobArtifacts             | DELETE   | /projects/:id/jobs/:job_id/artifacts                        |
| DeleteAllJobArtifacts          | DELETE   | /projects/:id/artifacts                                     |

## Project CI/CD Variables API

```http
GET    /projects/:id/variables
GET    /projects/:id/variables/:key
POST   /projects/:id/variables
PUT    /projects/:id/variables/:key
DELETE /projects/:id/variables/:key
```

### Functions

| Function                 | Method | Endpoint                           |
| ------------------------ | ------ | ---------------------------------- |
| ListProjectVariables     | GET    | /projects/:id/variables            |
| GetProjectVariable       | GET    | /projects/:id/variables/:key       |
| CreateProjectVariable    | POST   | /projects/:id/variables            |
| UpdateProjectVariable    | PUT    | /projects/:id/variables/:key       |
| DeleteProjectVariable    | DELETE | /projects/:id/variables/:key       |

## Runners API

```http
GET    /runners
GET    /runners/all
GET    /runners/:id
PUT    /runners/:id
GET    /runners/:id/jobs
GET    /runners/:id/managers
GET    /projects/:id/runners
POST   /projects/:id/runners
DELETE /projects/:id/runners/:runner_id
GET    /groups/:id/runners
POST   /runners
DELETE /runners/:id
DELETE /runners
POST   /runners/verify
POST   /runners/reset_registration_token
POST   /projects/:id/runners/reset_registration_token
POST   /groups/:id/runners/reset_registration_token
POST   /runners/:id/reset_authentication_token
POST   /runners/reset_authentication_token
GET    /runners/router/discovery
```

### Functions

| Function                           | Method | Endpoint                                          |
| ---------------------------------- | ------ | ------------------------------------------------- |
| ListRunners                        | GET    | /runners                                          |
| ListAllRunners                     | GET    | /runners/all                                      |
| GetRunner                          | GET    | /runners/:id                                      |
| UpdateRunner                       | PUT    | /runners/:id                                      |
| ListRunnerJobs                     | GET    | /runners/:id/jobs                                 |
| ListRunnerManagers                 | GET    | /runners/:id/managers                             |
| ListProjectRunners                 | GET    | /projects/:id/runners                             |
| AssignRunnerToProject              | POST   | /projects/:id/runners                             |
| UnassignRunnerFromProject          | DELETE | /projects/:id/runners/:runner_id                  |
| ListGroupRunners                   | GET    | /groups/:id/runners                               |
| CreateRunner                       | POST   | /runners                                          |
| DeleteRunner                       | DELETE | /runners/:id                                      |
| DeleteRunnerByToken                | DELETE | /runners                                          |
| VerifyRunner                       | POST   | /runners/verify                                   |
| ResetInstanceRegistrationToken     | POST   | /runners/reset_registration_token                 |
| ResetProjectRegistrationToken      | POST   | /projects/:id/runners/reset_registration_token    |
| ResetGroupRegistrationToken        | POST   | /groups/:id/runners/reset_registration_token      |
| ResetRunnerAuthenticationToken     | POST   | /runners/:id/reset_authentication_token           |
| ResetRunnerAuthenticationTokenByToken | POST | /runners/reset_authentication_token              |
| GetJobRouterDiscovery              | GET    | /runners/router/discovery                         |
