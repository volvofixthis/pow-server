package pow

import (
	"context"

	"github.com/volvofixthis/pow-server/internal/core/ports"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(NewPowService, fx.As(new(ports.PowService))),
	),
	fx.Invoke(registerLifecycleHooks),
)

func registerLifecycleHooks(lc fx.Lifecycle, ps ports.PowService) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return StartService(ps)
		},
		OnStop: func(ctx context.Context) error {
			return StopService(ps)
		},
	})
}
