package passage

import (
	"context"

	"github.com/volvofixthis/pow-server/internal/core/models"
	"github.com/volvofixthis/pow-server/internal/core/ports"
)

func NewPassageService(passageRepo ports.PassageRepository) *PassageService {
	return &PassageService{passageRepo: passageRepo}
}

type PassageService struct {
	passageRepo ports.PassageRepository
}

func (ps *PassageService) Get(ctx context.Context) (*models.Passage, error) {
	passage, err := ps.passageRepo.Get(ctx)
	if err != nil {
		return nil, err
	}
	return passage, nil
}
