package pow

import (
	"context"
	"errors"
	"time"

	"github.com/volvofixthis/pow-server/internal/core/models"
	"github.com/volvofixthis/pow-server/internal/core/ports"
	"github.com/volvofixthis/pow-server/internal/core/utils"
	"github.com/volvofixthis/pow-server/internal/infra/config"
	"go.uber.org/zap"
)

var (
	ErrWrongResult  = errors.New("Result isn't right")
	ErrCreateFailed = errors.New("Pow task creation failed")

	HashInput = "You miss 100% of the shots you don’t take."

	cleanupDelay = 100 * time.Millisecond
)

func NewPowService(log *zap.Logger, config *config.AppConfig, powRepo ports.PowRepository) *PowService {
	return &PowService{
		powRepo:    powRepo,
		c:          config,
		log:        log,
		shutdownCh: make(chan struct{}),
		taskCh:     make(chan *models.PowTask, config.EmissionSize),
	}
}

type PowService struct {
	powRepo    ports.PowRepository
	log        *zap.Logger
	c          *config.AppConfig
	shutdownCh chan struct{} // Channel to signal service shutdown
	taskCh     chan *models.PowTask
}

func (ps *PowService) Create(ctx context.Context) (*models.PowTask, error) {
	select {
	case task := <-ps.taskCh:
		task.CreatedAt = time.Now().UTC()
		if err := ps.powRepo.Create(ctx, task); err != nil {
			return nil, err
		}
		return task, nil
	case <-time.After(ps.c.EmissionDelay):
		return nil, ErrCreateFailed
	}
}

func (ps *PowService) preparePowTask() (*models.PowTask, error) {
	salt, err := utils.GenerateSalt(utils.SaltLength)
	if err != nil {
		return nil, err
	}
	hash := utils.GenerateProofOfWork(HashInput, salt, uint32(ps.c.PowIteration), uint32(ps.c.PowMemory))
	return &models.PowTask{Text: HashInput, Salt: salt, Hash: hash, Iteration: uint32(ps.c.PowIteration), Memory: uint32(ps.c.PowMemory)}, nil
}

func (ps *PowService) Verify(ctx context.Context, result *models.PowResult) error {
	if _, err := ps.powRepo.Get(ctx, result.Hash); err != nil {
		return ErrWrongResult
	}
	return nil
}

func (ps *PowService) cleanup(ctx context.Context) error {
	for v := range ps.powRepo.Items(ctx) {
		if time.Since(v.CreatedAt) > ps.c.PowTimeout {
			if err := ps.powRepo.Delete(ctx, v); err != nil {
				return err
			}
		}
	}
	return nil
}

var _ ports.PowService = (*PowService)(nil)

func StartService(ps ports.PowService) error {
	psv := ps.(*PowService)
	go func() {
		for {
			select {
			case <-psv.shutdownCh:
				return
			default:
				time.Sleep(cleanupDelay)
				if err := psv.cleanup(context.Background()); err != nil {
					psv.log.Error("Received error when cleanup", zap.Error(err))
					return
				}
			}
		}
	}()
	go func() {
		for {
			task, err := psv.preparePowTask()
			if err != nil {
				psv.log.Warn("Error when preparing pow task", zap.Error(err))
				return
			}
			select {
			case <-psv.shutdownCh:
				return
			case psv.taskCh <- task:
				psv.log.Warn("Added pow task to queue")
			default:
				time.Sleep(psv.c.EmissionDelay)
			}
		}
	}()
	return nil
}

func StopService(ps ports.PowService) error {
	psv := ps.(*PowService)
	psv.shutdownCh <- struct{}{}
	psv.shutdownCh <- struct{}{}
	return nil
}
