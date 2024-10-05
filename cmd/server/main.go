package main

import (
	"github.com/volvofixthis/pow-server/internal/adapters"
	"github.com/volvofixthis/pow-server/internal/core"
	"github.com/volvofixthis/pow-server/internal/infra"
	"github.com/volvofixthis/pow-server/internal/infra/config"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	cfg := config.NewCfg()
	app := fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(func() *config.AppConfig {
			return cfg
		}),
		core.Module,
		adapters.Module,
		infra.Module,
	)
	app.Run()
}
