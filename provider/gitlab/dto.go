package gitlab

import "time"

type JobCommit struct {
	ID             string   `json:"id"`
	ShortID        string   `json:"short_id"`
	Title          string   `json:"title"`
	CreatedAt      string   `json:"created_at"`
	ParentIDs      []string `json:"parent_ids"`
	Message        string   `json:"message"`
	AuthorName     string   `json:"author_name"`
	AuthorEmail    string   `json:"author_email"`
	AuthoredDate   string   `json:"authored_date"`
	CommitterName  string   `json:"committer_name"`
	CommitterEmail string   `json:"committer_email"`
	CommittedDate  string   `json:"committed_date"`
}

type JobPipeline struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"project_id,omitempty"`
	Ref       string `json:"ref"`
	SHA       string `json:"sha"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	WebURL    string `json:"web_url,omitempty"`
}

type JobRunner struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	IPAddress   any    `json:"ip_address"`
	Active      bool   `json:"active"`
	Paused      bool   `json:"paused"`
	IsShared    bool   `json:"is_shared"`
	RunnerType  string `json:"runner_type"`
	Name        any    `json:"name"`
	Online      bool   `json:"online"`
	Status      string `json:"status"`
}

type JobRunnerManager struct {
	ID           int       `json:"id"`
	SystemID     string    `json:"system_id"`
	Version      string    `json:"version"`
	Revision     string    `json:"revision"`
	Platform     string    `json:"platform"`
	Architecture string    `json:"architecture"`
	CreatedAt    time.Time `json:"created_at"`
	ContactedAt  time.Time `json:"contacted_at"`
	IPAddress    string    `json:"ip_address"`
	Status       string    `json:"status"`
}

type JobUser struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	State     string    `json:"state"`
	AvatarURL string    `json:"avatar_url"`
	WebURL    string    `json:"web_url"`
	CreatedAt time.Time `json:"created_at"`
	Bio       any       `json:"bio"`
	Location  any       `json:"location"`
	PublicEmail string  `json:"public_email"`
	LinkedIn  string    `json:"linkedin"`
	Twitter   string    `json:"twitter"`
	WebsiteURL string   `json:"website_url"`
	Organization any   `json:"organization"`
}

type JobArtifact struct {
	FileType   string `json:"file_type"`
	Size       int    `json:"size"`
	Filename   string `json:"filename"`
	FileFormat string `json:"file_format"`
}

type JobProject struct {
	CiJobTokenScopeEnabled bool `json:"ci_job_token_scope_enabled"`
}

type Job struct {
	ID                 int              `json:"id"`
	Name               string           `json:"name"`
	Stage              string           `json:"stage"`
	Status             string           `json:"status"`
	FailureReason      string           `json:"failure_reason,omitempty"`
	CreatedAt          string           `json:"created_at"`
	StartedAt          *string          `json:"started_at"`
	FinishedAt         *string          `json:"finished_at"`
	ErasedAt           *string          `json:"erased_at"`
	Duration           *float64         `json:"duration"`
	QueuedDuration     *float64         `json:"queued_duration"`
	Ref                string           `json:"ref"`
	Tag                bool             `json:"tag"`
	TagList            []string         `json:"tag_list"`
	AllowFailure       bool             `json:"allow_failure"`
	Archived           bool             `json:"archived,omitempty"`
	Coverage           any              `json:"coverage"`
	Source             string           `json:"source,omitempty"`
	WebURL             string           `json:"web_url"`
	ArtifactsExpireAt  *string          `json:"artifacts_expire_at"`
	ArtifactsFile      *JobArtifactsFile `json:"artifacts_file,omitempty"`
	Artifacts          []JobArtifact    `json:"artifacts"`
	Commit             JobCommit        `json:"commit"`
	Pipeline           JobPipeline      `json:"pipeline"`
	Runner             *JobRunner       `json:"runner"`
	RunnerManager      *JobRunnerManager `json:"runner_manager"`
	Project            *JobProject      `json:"project"`
	User               *JobUser         `json:"user"`
}

type JobArtifactsFile struct {
	Filename string `json:"filename"`
	Size     int    `json:"size"`
}

type TriggerJobDownstreamPipeline struct {
	ID        int    `json:"id"`
	SHA       string `json:"sha"`
	Ref       string `json:"ref"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	WebURL    string `json:"web_url"`
}

type TriggerJob struct {
	Job
	DownstreamPipeline *TriggerJobDownstreamPipeline `json:"downstream_pipeline"`
}

type AllowedAgentConfigProject struct {
	ID                int    `json:"id"`
	Description       any    `json:"description"`
	Name              string `json:"name"`
	NameWithNamespace string `json:"name_with_namespace"`
	Path              string `json:"path"`
	PathWithNamespace string `json:"path_with_namespace"`
	CreatedAt         string `json:"created_at"`
}

type AllowedAgent struct {
	ID             int                      `json:"id"`
	ConfigProject  AllowedAgentConfigProject `json:"config_project"`
}

type AllowedAgentsResponse struct {
	AllowedAgents []AllowedAgent `json:"allowed_agents"`
	Job           struct {
		ID int `json:"id"`
	} `json:"job"`
	Pipeline struct {
		ID int `json:"id"`
	} `json:"pipeline"`
	Project struct {
		ID     int `json:"id"`
		Groups []struct {
			ID int `json:"id"`
		} `json:"groups"`
	} `json:"project"`
	User *JobUser `json:"user"`
}

type Variable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PlayJobRequest struct {
	JobInputs              map[string]any `json:"job_inputs,omitempty"`
	JobVariablesAttributes []Variable     `json:"job_variables_attributes,omitempty"`
}

type RetryJobRequest struct {
	JobInputs map[string]any `json:"job_inputs,omitempty"`
}

// Pipeline represents a GitLab CI/CD pipeline.
//
// A pipeline is a group of jobs that are organized into stages and run in
// sequence or in parallel. Pipelines are the fundamental unit of work in
// GitLab CI/CD.
type Pipeline struct {
	ID        int    `json:"id"`
	IID       int    `json:"iid"`
	ProjectID int    `json:"project_id"`
	Status    string `json:"status"`
	Source    string `json:"source"`
	Ref       string `json:"ref"`
	SHA       string `json:"sha"`
	Name      string `json:"name"`
	WebURL    string `json:"web_url"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	BeforeSHA string  `json:"before_sha"`
	Tag       bool    `json:"tag"`
	YAMLErrors *string `json:"yaml_errors"`

	User *PipelineUser `json:"user,omitempty"`

	StartedAt      *time.Time `json:"started_at,omitempty"`
	FinishedAt     *time.Time `json:"finished_at,omitempty"`
	CommittedAt    *time.Time `json:"committed_at,omitempty"`
	Duration       *int       `json:"duration,omitempty"`
	QueuedDuration *float64   `json:"queued_duration,omitempty"`
	Coverage       *string    `json:"coverage,omitempty"`
	Archived       bool       `json:"archived"`
}

// PipelineUser represents the user who triggered a pipeline.
type PipelineUser struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	State     string `json:"state"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

// PipelineDetailedStatus contains detailed status information about a pipeline.
type PipelineDetailedStatus struct {
	Icon        string `json:"icon"`
	Text        string `json:"text"`
	Label       string `json:"label"`
	Group       string `json:"group"`
	Tooltip     string `json:"tooltip"`
	HasDetails  bool   `json:"has_details"`
	DetailsPath string `json:"details_path"`
	Illustration any   `json:"illustration"`
	Favicon     string `json:"favicon"`
}

// PipelineDetailed extends Pipeline with additional detail fields returned
// by the single pipeline endpoint.
type PipelineDetailed struct {
	Pipeline
	DetailedStatus PipelineDetailedStatus `json:"detailed_status"`
}

// PipelineVariable represents a variable defined in a pipeline.
type PipelineVariable struct {
	Key          string `json:"key"`
	VariableType string `json:"variable_type"`
	Value        string `json:"value"`
}

// PipelineTestReport represents the test report for a pipeline.
type PipelineTestReport struct {
	TotalTime    int                 `json:"total_time"`
	TotalCount   int                 `json:"total_count"`
	SuccessCount int                 `json:"success_count"`
	FailedCount  int                 `json:"failed_count"`
	SkippedCount int                 `json:"skipped_count"`
	ErrorCount   int                 `json:"error_count"`
	TestSuites   []PipelineTestSuite `json:"test_suites"`
}

// PipelineTestSuite represents a single test suite within a pipeline test report.
type PipelineTestSuite struct {
	Name         string             `json:"name"`
	TotalTime    int                `json:"total_time"`
	TotalCount   int                `json:"total_count"`
	SuccessCount int                `json:"success_count"`
	FailedCount  int                `json:"failed_count"`
	SkippedCount int                `json:"skipped_count"`
	ErrorCount   int                `json:"error_count"`
	TestCases    []PipelineTestCase `json:"test_cases"`
}

// PipelineTestCase represents a single test case within a test suite.
type PipelineTestCase struct {
	Status        string  `json:"status"`
	Name          string  `json:"name"`
	Classname     string  `json:"classname"`
	ExecutionTime int     `json:"execution_time"`
	SystemOutput  *string `json:"system_output"`
	StackTrace    *string `json:"stack_trace"`
}

// PipelineTestReportSummary represents a summary of the test report for a pipeline.
type PipelineTestReportSummary struct {
	Total      PipelineTestReportTotal          `json:"total"`
	TestSuites []PipelineTestReportSummarySuite `json:"test_suites"`
}

// PipelineTestReportTotal contains the total counts for a pipeline test report.
type PipelineTestReportTotal struct {
	Time       int     `json:"time"`
	Count      int     `json:"count"`
	Success    int     `json:"success"`
	Failed     int     `json:"failed"`
	Skipped    int     `json:"skipped"`
	Error      int     `json:"error"`
	SuiteError *string `json:"suite_error"`
}

// PipelineTestReportSummarySuite represents a single test suite in the summary.
type PipelineTestReportSummarySuite struct {
	Name         string  `json:"name"`
	TotalTime    int     `json:"total_time"`
	TotalCount   int     `json:"total_count"`
	SuccessCount int     `json:"success_count"`
	FailedCount  int     `json:"failed_count"`
	SkippedCount int     `json:"skipped_count"`
	ErrorCount   int     `json:"error_count"`
	BuildIDs     []int   `json:"build_ids"`
	SuiteError   *string `json:"suite_error"`
}

// PipelineMetadata represents the metadata of a pipeline that can be updated.
type PipelineMetadata struct {
	Name string `json:"name"`
}
