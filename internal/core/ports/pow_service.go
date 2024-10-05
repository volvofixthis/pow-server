package ports

import (
	"context"

	"github.com/volvofixthis/pow-server/internal/core/models"
)

type PowService interface {
	Create(ctx context.Context) (*models.PowTask, error)
	Verify(ctx context.Context, result *models.PowResult) error
}
