package gitlab

import (
	"context"
	"fmt"
	"net/http"
)

// ListPipelineSchedules retrieves all pipeline schedules for a project.
//
// By default, this request returns 20 results at a time because the API
// results are paginated.
//
// The caller must have at least reporter access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - opts: Optional filtering parameters.
//
// Returns a slice of pipeline schedules or an error if the request fails or
// the API returns a non-success status.
func (c *Client) ListPipelineSchedules(ctx context.Context, projectID int, opts *ListPipelineSchedulesOptions) ([]PipelineSchedule, error) {
	path := fmt.Sprintf("/projects/%d/pipeline_schedules", projectID)
	path = appendQueryParams(path, opts)

	var resp []PipelineSchedule
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list pipeline schedules: %w", err)
	}
	return resp, nil
}

// GetPipelineSchedule retrieves a pipeline schedule for a project.
//
// The caller must have at least reporter access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - scheduleID: The ID of the pipeline schedule.
//
// Returns the requested pipeline schedule or an error if the request fails,
// the schedule does not exist, or the API returns a non-success status.
func (c *Client) GetPipelineSchedule(ctx context.Context, projectID, scheduleID int) (PipelineSchedule, error) {
	path := fmt.Sprintf("/projects/%d/pipeline_schedules/%d", projectID, scheduleID)

	var resp PipelineSchedule
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return PipelineSchedule{}, fmt.Errorf("failed to get pipeline schedule: %w", err)
	}
	return resp, nil
}

// ListPipelinesBySchedule retrieves all pipelines triggered by a pipeline schedule.
//
// The caller must have at least reporter access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - scheduleID: The ID of the pipeline schedule.
//   - opts: Optional filtering parameters.
//
// Returns a slice of pipelines or an error if the request fails or the API
// returns a non-success status.
func (c *Client) ListPipelinesBySchedule(ctx context.Context, projectID, scheduleID int, opts *ListPipelinesByScheduleOptions) ([]Pipeline, error) {
	path := fmt.Sprintf("/projects/%d/pipeline_schedules/%d/pipelines", projectID, scheduleID)
	path = appendQueryParams(path, opts)

	var resp []Pipeline
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list pipelines by schedule: %w", err)
	}
	return resp, nil
}

// CreatePipelineSchedule creates a new pipeline schedule for a project.
//
// The caller must have at least maintainer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - opts: The schedule configuration including cron, description, and ref.
//
// Returns the newly created pipeline schedule or an error if the request fails
// or the API returns a non-success status.
func (c *Client) CreatePipelineSchedule(ctx context.Context, projectID int, opts *CreatePipelineScheduleRequest) (PipelineSchedule, error) {
	path := fmt.Sprintf("/projects/%d/pipeline_schedules", projectID)

	var resp PipelineSchedule
	if err := c.do(ctx, http.MethodPost, path, opts, &resp); err != nil {
		return PipelineSchedule{}, fmt.Errorf("failed to create pipeline schedule: %w", err)
	}
	return resp, nil
}

// UpdatePipelineSchedule updates a pipeline schedule for a project.
//
// After the update is done, the schedule is rescheduled automatically.
//
// The caller must have at least maintainer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - scheduleID: The ID of the pipeline schedule.
//   - opts: The fields to update.
//
// Returns the updated pipeline schedule or an error if the request fails or
// the API returns a non-success status.
func (c *Client) UpdatePipelineSchedule(ctx context.Context, projectID, scheduleID int, opts *UpdatePipelineScheduleRequest) (PipelineSchedule, error) {
	path := fmt.Sprintf("/projects/%d/pipeline_schedules/%d", projectID, scheduleID)

	var resp PipelineSchedule
	if err := c.do(ctx, http.MethodPut, path, opts, &resp); err != nil {
		return PipelineSchedule{}, fmt.Errorf("failed to update pipeline schedule: %w", err)
	}
	return resp, nil
}

// TakeOwnershipPipelineSchedule updates the owner of a pipeline schedule to
// the authenticated user.
//
// The caller must have at least maintainer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - scheduleID: The ID of the pipeline schedule.
//
// Returns the updated pipeline schedule or an error if the request fails or
// the API returns a non-success status.
func (c *Client) TakeOwnershipPipelineSchedule(ctx context.Context, projectID, scheduleID int) (PipelineSchedule, error) {
	path := fmt.Sprintf("/projects/%d/pipeline_schedules/%d/take_ownership", projectID, scheduleID)

	var resp PipelineSchedule
	if err := c.do(ctx, http.MethodPost, path, nil, &resp); err != nil {
		return PipelineSchedule{}, fmt.Errorf("failed to take ownership of pipeline schedule: %w", err)
	}
	return resp, nil
}

// DeletePipelineSchedule deletes a pipeline schedule for a project.
//
// The caller must have at least maintainer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - scheduleID: The ID of the pipeline schedule.
//
// Returns an error if the deletion fails or the API returns a non-success
// status.
func (c *Client) DeletePipelineSchedule(ctx context.Context, projectID, scheduleID int) error {
	path := fmt.Sprintf("/projects/%d/pipeline_schedules/%d", projectID, scheduleID)

	if err := c.do(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("failed to delete pipeline schedule: %w", err)
	}
	return nil
}

// PlayPipelineSchedule runs a pipeline schedule immediately.
//
// The next scheduled run of this pipeline is not affected.
//
// The caller must have at least developer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - scheduleID: The ID of the pipeline schedule.
//
// Returns an error if the request fails or the API returns a non-success
// status.
func (c *Client) PlayPipelineSchedule(ctx context.Context, projectID, scheduleID int) error {
	path := fmt.Sprintf("/projects/%d/pipeline_schedules/%d/play", projectID, scheduleID)

	if err := c.do(ctx, http.MethodPost, path, nil, nil); err != nil {
		return fmt.Errorf("failed to play pipeline schedule: %w", err)
	}
	return nil
}

// CreateScheduleVariable creates a new variable for a pipeline schedule.
//
// The caller must have at least maintainer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - scheduleID: The ID of the pipeline schedule.
//   - opts: The variable configuration including key, value, and optional type.
//
// Returns the newly created variable or an error if the request fails or the
// API returns a non-success status.
func (c *Client) CreateScheduleVariable(ctx context.Context, projectID, scheduleID int, opts *PipelineScheduleVariableRequest) (PipelineScheduleVariable, error) {
	path := fmt.Sprintf("/projects/%d/pipeline_schedules/%d/variables", projectID, scheduleID)

	var resp PipelineScheduleVariable
	if err := c.do(ctx, http.MethodPost, path, opts, &resp); err != nil {
		return PipelineScheduleVariable{}, fmt.Errorf("failed to create schedule variable: %w", err)
	}
	return resp, nil
}

// GetScheduleVariable retrieves a variable for a pipeline schedule.
//
// The caller must have at least reporter access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - scheduleID: The ID of the pipeline schedule.
//   - key: The key of the variable.
//
// Returns the requested variable or an error if the request fails, the
// variable does not exist, or the API returns a non-success status.
func (c *Client) GetScheduleVariable(ctx context.Context, projectID, scheduleID int, key string) (PipelineScheduleVariable, error) {
	path := fmt.Sprintf("/projects/%d/pipeline_schedules/%d/variables/%s", projectID, scheduleID, key)

	var resp PipelineScheduleVariable
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return PipelineScheduleVariable{}, fmt.Errorf("failed to get schedule variable: %w", err)
	}
	return resp, nil
}

// UpdateScheduleVariable updates a variable for a pipeline schedule.
//
// The caller must have at least maintainer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - scheduleID: The ID of the pipeline schedule.
//   - key: The key of the variable.
//   - opts: The fields to update including value and optional type.
//
// Returns the updated variable or an error if the request fails or the API
// returns a non-success status.
func (c *Client) UpdateScheduleVariable(ctx context.Context, projectID, scheduleID int, key string, opts *PipelineScheduleVariableRequest) (PipelineScheduleVariable, error) {
	path := fmt.Sprintf("/projects/%d/pipeline_schedules/%d/variables/%s", projectID, scheduleID, key)

	var resp PipelineScheduleVariable
	if err := c.do(ctx, http.MethodPut, path, opts, &resp); err != nil {
		return PipelineScheduleVariable{}, fmt.Errorf("failed to update schedule variable: %w", err)
	}
	return resp, nil
}

// DeleteScheduleVariable deletes a variable for a pipeline schedule.
//
// The caller must have at least maintainer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - scheduleID: The ID of the pipeline schedule.
//   - key: The key of the variable.
//
// Returns an error if the deletion fails or the API returns a non-success
// status.
func (c *Client) DeleteScheduleVariable(ctx context.Context, projectID, scheduleID int, key string) error {
	path := fmt.Sprintf("/projects/%d/pipeline_schedules/%d/variables/%s", projectID, scheduleID, key)

	if err := c.do(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("failed to delete schedule variable: %w", err)
	}
	return nil
}
