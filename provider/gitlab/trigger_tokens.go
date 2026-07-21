package gitlab

import (
	"context"
	"fmt"
	"net/http"
)

// ListTriggerTokens retrieves all pipeline trigger tokens for a project.
//
// The trigger token is displayed in full if the trigger token was created by
// the authenticated user. Trigger tokens created by other users are shortened
// to four characters.
//
// The caller must have at least maintainer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//
// Returns a slice of trigger tokens or an error if the request fails or the
// API returns a non-success status.
func (c *Client) ListTriggerTokens(ctx context.Context, projectID int) ([]TriggerToken, error) {
	path := fmt.Sprintf("/projects/%d/triggers", projectID)

	var resp []TriggerToken
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list trigger tokens: %w", err)
	}
	return resp, nil
}

// GetTriggerToken retrieves details of a project's pipeline trigger token.
//
// The caller must have at least maintainer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - triggerID: The ID of the trigger token.
//
// Returns the requested trigger token or an error if the request fails, the
// trigger token does not exist, or the API returns a non-success status.
func (c *Client) GetTriggerToken(ctx context.Context, projectID, triggerID int) (TriggerToken, error) {
	path := fmt.Sprintf("/projects/%d/triggers/%d", projectID, triggerID)

	var resp TriggerToken
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return TriggerToken{}, fmt.Errorf("failed to get trigger token: %w", err)
	}
	return resp, nil
}

// CreateTriggerToken creates a pipeline trigger token for a project.
//
// The caller must have at least maintainer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - description: The trigger token description/name.
//
// Returns the newly created trigger token or an error if the request fails
// or the API returns a non-success status.
func (c *Client) CreateTriggerToken(ctx context.Context, projectID int, description string) (TriggerToken, error) {
	path := fmt.Sprintf("/projects/%d/triggers", projectID)

	body := map[string]string{"description": description}

	var resp TriggerToken
	if err := c.do(ctx, http.MethodPost, path, body, &resp); err != nil {
		return TriggerToken{}, fmt.Errorf("failed to create trigger token: %w", err)
	}
	return resp, nil
}

// UpdateTriggerToken updates a project's pipeline trigger token.
//
// The caller must have at least maintainer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - triggerID: The ID of the trigger token.
//   - description: The new trigger token description/name.
//
// Returns the updated trigger token or an error if the request fails or the
// API returns a non-success status.
func (c *Client) UpdateTriggerToken(ctx context.Context, projectID, triggerID int, description string) (TriggerToken, error) {
	path := fmt.Sprintf("/projects/%d/triggers/%d", projectID, triggerID)

	body := map[string]string{"description": description}

	var resp TriggerToken
	if err := c.do(ctx, http.MethodPut, path, body, &resp); err != nil {
		return TriggerToken{}, fmt.Errorf("failed to update trigger token: %w", err)
	}
	return resp, nil
}

// DeleteTriggerToken deletes a project's pipeline trigger token.
//
// The caller must have at least maintainer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - triggerID: The ID of the trigger token.
//
// Returns an error if the deletion fails or the API returns a non-success
// status.
func (c *Client) DeleteTriggerToken(ctx context.Context, projectID, triggerID int) error {
	path := fmt.Sprintf("/projects/%d/triggers/%d", projectID, triggerID)

	if err := c.do(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("failed to delete trigger token: %w", err)
	}
	return nil
}

// TriggerPipeline triggers a pipeline using a pipeline trigger token or
// CI/CD job token for authentication.
//
// With a CI/CD job token, the triggered pipeline is a multi-project pipeline.
// The job that authenticates the request becomes associated with the upstream
// pipeline, which is visible on the pipeline graph.
//
// If you use a trigger token in a job, the job is not associated with the
// upstream pipeline.
//
// The caller must have at least developer access to the project.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The ID or URL-encoded path of the project.
//   - opts: The trigger pipeline request containing ref, token, and optional
//     variables and inputs.
//
// Returns the newly created pipeline or an error if the request fails or the
// API returns a non-success status.
func (c *Client) TriggerPipeline(ctx context.Context, projectID int, opts *TriggerPipelineRequest) (Pipeline, error) {
	path := fmt.Sprintf("/projects/%d/trigger/pipeline?ref=%s&token=%s", projectID, opts.Ref, opts.Token)

	var resp Pipeline
	if err := c.do(ctx, http.MethodPost, path, opts, &resp); err != nil {
		return Pipeline{}, fmt.Errorf("failed to trigger pipeline: %w", err)
	}
	return resp, nil
}
