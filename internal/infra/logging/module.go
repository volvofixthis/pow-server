package logging

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Options(
	fx.Provide(NewZapLogger),
	fx.Invoke(registerLifecycleHooks),
)

func registerLifecycleHooks(lc fx.Lifecycle, log *zap.Logger) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return log.Sync()
		},
	})
}
