package circleci

import (
	"context"
	"fmt"
	"net/http"
)

// ListPipelineDefinitionTriggers returns all triggers for a pipeline definition.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The unique identifier of the project.
//   - pipelineDefID: The unique identifier of the pipeline definition.
//
// Returns the list of triggers or an error if the request fails.
func (c *Client) ListPipelineDefinitionTriggers(ctx context.Context, projectID, pipelineDefID string) ([]PipelineDefinitionTrigger, error) {
	path := fmt.Sprintf("/projects/%s/pipeline-definitions/%s/triggers", projectID, pipelineDefID)

	var resp []PipelineDefinitionTrigger
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return []PipelineDefinitionTrigger{}, fmt.Errorf("failed to get pipeline definition triggers: %w", err)
	}
	return resp, nil
}

// CreateTrigger creates a new trigger for a pipeline definition.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The unique identifier of the project.
//   - pipelineDefID: The unique identifier of the pipeline definition.
//   - body: The trigger configuration including event source and settings.
//
// Returns the created trigger or an error if the request fails.
func (c *Client) CreateTrigger(ctx context.Context, projectID, pipelineDefID string, body CreateTriggerBody) (PipelineDefinitionTrigger, error) {
	path := fmt.Sprintf("/projects/%s/pipeline-definitions/%s/triggers", projectID, pipelineDefID)

	var resp PipelineDefinitionTrigger
	if err := c.do(ctx, http.MethodPost, path, body, &resp); err != nil {
		return PipelineDefinitionTrigger{}, fmt.Errorf("failed to create pipeline definition trigger: %w", err)
	}
	return resp, nil
}

// GetTrigger returns a single trigger by its ID.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The unique identifier of the project.
//   - triggerID: The unique identifier of the trigger.
//
// Returns the trigger or an error if the request fails.
func (c *Client) GetTrigger(ctx context.Context, projectID, triggerID string) (PipelineDefinitionTrigger, error) {
	path := fmt.Sprintf("/projects/%s/triggers/%s", projectID, triggerID)

	var resp PipelineDefinitionTrigger
	if err := c.do(ctx, http.MethodGet, path, nil, &resp); err != nil {
		return PipelineDefinitionTrigger{}, fmt.Errorf("failed to get trigger{%s}: %w", triggerID, err)
	}
	return resp, nil
}

// PatchTrigger updates an existing trigger with the provided fields.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The unique identifier of the project.
//   - pipelineDefID: The unique identifier of the pipeline definition.
//   - body: The updated trigger configuration.
//
// Returns the updated trigger or an error if the request fails.
func (c *Client) PatchTrigger(ctx context.Context, projectID, pipelineDefID string, body CreateTriggerBody) (PipelineDefinitionTrigger, error) {
	path := fmt.Sprintf("/projects/%s/triggers/%s", projectID, pipelineDefID)

	var resp PipelineDefinitionTrigger
	if err := c.do(ctx, http.MethodPatch, path, body, &resp); err != nil {
		return PipelineDefinitionTrigger{}, fmt.Errorf("failed to create pipeline definition trigger: %w", err)
	}
	return resp, nil
}

// DeleteTrigger deletes a trigger by its ID.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts.
//   - projectID: The unique identifier of the project.
//   - triggerID: The unique identifier of the trigger to delete.
//
// Returns an error if the request fails.
func (c *Client) DeleteTrigger(ctx context.Context, projectID, triggerID string) error {
	path := fmt.Sprintf("/projects/%s/triggers/%s", projectID, triggerID)

	if err := c.do(ctx, http.MethodGet, path, nil, nil); err != nil {
		return fmt.Errorf("failed to get trigger{%s}: %w", triggerID, err)
	}
	return nil
}
