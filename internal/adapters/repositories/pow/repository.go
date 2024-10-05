package pow

import (
	"context"
	"errors"
	"iter"
	"sync"

	"github.com/volvofixthis/pow-server/internal/core/models"
	"github.com/volvofixthis/pow-server/internal/core/ports"
	"github.com/volvofixthis/pow-server/internal/core/utils"
)

var ErrNotFound = errors.New("Task isn't found by hash")

func NewPowRepository() *PowRepository {
	return &PowRepository{}
}

type PowRepository struct {
	storage sync.Map
}

type keyHash [utils.OutputLength]byte

func makeHashable(hash []byte) keyHash {
	result := keyHash{}
	copy(result[:], hash)
	return result
}

func (pr *PowRepository) Get(ctx context.Context, hash []byte) (*models.PowTask, error) {
	if v, ok := pr.storage.LoadAndDelete(makeHashable(hash)); ok {
		if task, ok := v.(*models.PowTask); ok {
			return task, nil
		}
	}
	return nil, ErrNotFound
}

func (pr *PowRepository) Create(ctx context.Context, task *models.PowTask) error {
	pr.storage.Store(makeHashable(task.Hash), task)
	return nil
}

func (pr *PowRepository) Delete(ctx context.Context, task *models.PowTask) error {
	pr.storage.Delete(makeHashable(task.Hash))
	return nil
}

func (pr *PowRepository) Items(ctx context.Context) iter.Seq[*models.PowTask] {
	return func(yield func(*models.PowTask) bool) {
		pr.storage.Range(func(key any, v any) bool {
			return yield(v.(*models.PowTask))
		})
	}
}

var _ ports.PowRepository = (*PowRepository)(nil)
