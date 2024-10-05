package ports

import (
	"context"
	"iter"

	"github.com/volvofixthis/pow-server/internal/core/models"
)

type PowRepository interface {
	Get(ctx context.Context, Hash []byte) (*models.PowTask, error)
	Create(ctx context.Context, task *models.PowTask) error
	Items(ctx context.Context) iter.Seq[*models.PowTask]
	Delete(ctx context.Context, task *models.PowTask) error
}
