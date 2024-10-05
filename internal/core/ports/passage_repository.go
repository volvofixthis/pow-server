package ports

import (
	"context"

	"github.com/volvofixthis/pow-server/internal/core/models"
)

type PassageRepository interface {
	Get(ctx context.Context) (*models.Passage, error)
}
