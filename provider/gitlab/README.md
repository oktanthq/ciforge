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

| Function | Method | Endpoint |
|---|---|---|
| `ListJobsForProject` | `GET` | `/projects/:id/jobs` |
| `ListJobsForPipeline` | `GET` | `/projects/:id/pipelines/:pipeline_id/jobs` |
| `ListTriggerJobsForPipeline` | `GET` | `/projects/:id/pipelines/:pipeline_id/trigger_jobs` |
| `GetJobByToken` | `GET` | `/job` |
| `GetAllowedAgents` | `GET` | `/job/allowed_agents` |
| `GetJob` | `GET` | `/projects/:id/jobs/:job_id` |
| `GetJobTrace` | `GET` | `/projects/:id/jobs/:job_id/trace` |
| `CancelJob` | `POST` | `/projects/:id/jobs/:job_id/cancel` |
| `RetryJob` | `POST` | `/projects/:id/jobs/:job_id/retry` |
| `EraseJob` | `POST` | `/projects/:id/jobs/:job_id/erase` |
| `PlayJob` | `POST` | `/projects/:id/jobs/:job_id/play` |

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

| Function | Method | Endpoint |
|---|---|---|
| `ListPipelines` | `GET` | `/projects/:id/pipelines` |
| `GetPipeline` | `GET` | `/projects/:id/pipelines/:pipeline_id` |
| `GetLatestPipeline` | `GET` | `/projects/:id/pipelines/latest` |
| `GetPipelineVariables` | `GET` | `/projects/:id/pipelines/:pipeline_id/variables` |
| `GetPipelineTestReport` | `GET` | `/projects/:id/pipelines/:pipeline_id/test_report` |
| `GetPipelineTestReportSummary` | `GET` | `/projects/:id/pipelines/:pipeline_id/test_report_summary` |
| `CreatePipeline` | `POST` | `/projects/:id/pipeline?ref=...` |
| `RetryPipeline` | `POST` | `/projects/:id/pipelines/:pipeline_id/retry` |
| `CancelPipeline` | `POST` | `/projects/:id/pipelines/:pipeline_id/cancel` |
| `DeletePipeline` | `DELETE` | `/projects/:id/pipelines/:pipeline_id` |
| `UpdatePipelineMetadata` | `PUT` | `/projects/:id/pipelines/:pipeline_id/metadata` |
