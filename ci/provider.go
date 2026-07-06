package ci

import (
	"context"

	"github.com/abdulahshoaib/ciforge/models"
)

type Provider interface {
	ListPipelines(ctx context.Context, repo string) ([]models.Pipeline, error)
	GetPipeline(ctx context.Context, repo string, id string) (*models.Pipeline, error)
	ListRuns(ctx context.Context, repo string, pipeline_id string) ([]models.Run, error)
	GetRuns(ctx context.Context, repo string, run_id string) (*models.Run, error)
	CancelRun(ctx context.Context, repo string, run_id string) error
	RetryRun(ctx context.Context, repo string, run_id string) error
}
