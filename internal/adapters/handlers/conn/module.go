package conn

import (
	"github.com/volvofixthis/pow-server/internal/core/ports"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(
			NewConnAdapter,
			fx.As(new(ports.ConnAdapter)),
		),
	),
)
