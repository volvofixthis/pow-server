package pow

import (
	"github.com/volvofixthis/pow-server/internal/core/ports"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(NewPowRepository, fx.As(new(ports.PowRepository))),
	),
)
