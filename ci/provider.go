package ci

import (
	"context"

	"github.com/abdulahshoaib/ciforge/models"
)

type Provider interface {
	ListWorkflows(ctx context.Context, repo, owner string) ([]models.Pipeline, error)
	GetWorkflow(ctx context.Context, repo, owner, id string) (*models.Pipeline, error)
	ListRunsForRepo(ctx context.Context, repo, owner string) ([]models.Run, error)
	GetWorkflowRun(ctx context.Context, repo, owner, runID string) (models.Run, error)
	CancelWorkflowRun(ctx context.Context, repo, owner, runID string) error
	ReRunWorkflow(ctx context.Context, repo, owner, runID string) error
}
