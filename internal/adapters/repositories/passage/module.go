package passage

import (
	"github.com/volvofixthis/pow-server/internal/core/ports"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(NewPassageRepository, fx.As(new(ports.PassageRepository))),
	),
)
